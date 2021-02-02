package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/factory/db"
)

func applyMaster(initiator string, args map[string][]byte) (error, int) {
	desc, ok := args["desc"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	err := db.InsertRow(&db.Master{Name: initiator, Desc: string(desc)})
	if err != nil {
		return err, define.InsertDBErr
	}
	return err, 0
}

func commentMaster(args map[string][]byte) (error, int) {
	name, ok := args["name"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	status, ok := args["status"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	master, err := db.GetMasterByName(string(name))
	if err != nil {
		return err, define.QueryDBErr
	}
	err = db.UpdateMasterStatus(master, string(status))
	if err != nil {
		return err, define.UpdateDBErr
	}
	return nil, 0
}
