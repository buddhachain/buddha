package handler

import (
	"encoding/json"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/factory/db"
)

func applyFounder(amount, initiator string, args []byte) (error, int) {
	//txInfo := &db.Product{Initiator: tx.Initiator}
	info := db.Founder{}
	err := json.Unmarshal(args, &info)
	if err != nil {
		return err, define.UnmarshalErr
	}
	info.Name = initiator
	info.Amount = amount
	info.Status = 1
	err = db.InsertRow(&info)
	if err != nil {
		return err, define.InsertDBErr
	}
	return nil, 0
}

func commentFounder(args map[string][]byte) (error, int) {
	name, ok := args["name"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	status, ok := args["status"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	founder, err := db.GetFounderByName(string(name))
	if err != nil {
		return err, define.QueryDBErr
	}
	err = db.UpdateFounderStatus(founder, string(status))
	if err != nil {
		return err, define.UpdateDBErr
	}
	return nil, 0
}
