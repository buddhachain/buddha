package factory

import (
	"context"
	"encoding/json"
	"io"

	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/eventserver/config"
	"github.com/buddhachain/buddha/eventserver/db"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuperchain/core/pb"
	"google.golang.org/grpc"
)

var stream pb.EventService_SubscribeClient
var events chan *pb.Event
var client pb.EventServiceClient

var logger = utils.NewLogger("debug", "config")

func HandleStream() {
	events = make(chan *pb.Event, 10)
	go func() {
		for {
			event, err := stream.Recv()
			if err != nil && err != io.EOF {
				panic(err)
			}
			events <- event
		}
	}()
	go HandlerEvents()
}

func InitXchainClient() error {
	conf := config.ChainConf()
	logger.Infof("Using chain config %+v", conf)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, conf.Node, grpc.WithInsecure())
	if err != nil {
		return errors.WithMessage(err, "dial xchain server failed")
	}
	client = pb.NewEventServiceClient(conn)
	filter := &pb.BlockFilter{
		Bcname: conf.BcName,
	}

	buf, _ := proto.Marshal(filter)
	request := &pb.SubscribeRequest{
		Type:   pb.SubscribeType_BLOCK,
		Filter: buf,
	}

	stream, err = client.Subscribe(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func HandlerEvents() error {
	for {
		event := <-events
		var block pb.FilteredBlock
		err := proto.Unmarshal(event.Payload, &block)
		if err != nil {
			return errors.WithMessage(err, "unmarshal block")
		}

		fBlock := &FilteredBlock{
			Bcname:      block.Bcname,
			BlockID:     block.Blockid,
			BlockHeight: block.BlockHeight,
			TxCount:     len(block.Txs),
		}
		logger.Infof("Parsing block %+v", fBlock)
		_, err = db.MDB.Collection("block").InsertOne(context.TODO(), &block)
		if err != nil {
			logger.Errorf("Insert proto block failed %s", err.Error())
		}
		err = InsertFilteredBlock(fBlock)
		if err != nil {
			logger.Errorf("Insert block failed %s", err.Error())
			return err
		}
		for _, fTx := range block.Txs {
			events, err := json.Marshal(fTx.Events)
			if err != nil {
				logger.Errorf("Marshal tx events failed %s", err.Error())
				return err
			}
			tx := &FilteredTransaction{
				Txid:    fTx.Txid,
				BlockID: block.Blockid,
				Events:  string(events),
			}
			err = InsertFilteredTx(tx)
			if err != nil {
				logger.Errorf("Insert tx failed %s", err.Error())
				return err
			}
			for _, event := range fTx.Events {
				cEvent := &ContractEvent{
					Txid:     fTx.Txid,
					BlockID:  block.Blockid,
					Contract: event.Contract,
					Name:     event.Name,
					Body:     string(event.Body),
				}
				err = InsertContractEvent(cEvent)
				if err != nil {
					logger.Errorf("Insert contract event failed %s", err.Error())
					return err
				}
			}
		}
	}
}
