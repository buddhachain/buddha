package factory

import (
	"encoding/json"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuperchain/core/pb"
)

var stream pb.EventService_SubscribeClient
var events chan *pb.Event

func HandleStram() {
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
		}
	}
}
