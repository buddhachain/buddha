package factory

import (
	"context"
	"encoding/json"

	"github.com/buddhachain/buddha/eventserver/db"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuperchain/core/pb"
)

func parseEvent(event *pb.Event) error {
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
	return nil
}
