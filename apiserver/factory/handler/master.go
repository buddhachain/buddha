package handler

import (
	"strconv"

	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/gin-gonic/gin"
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

func ApplyMaster(c *gin.Context) {
	logger.Debug("Entering master applying...")

	req := &db.Master{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	logger.Infof("Request info %+v", req)
	err = db.InsertRow(req)
	if err != nil {
		logger.Errorf("Insert founder apply failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

func ApproveMaster(c *gin.Context) {
	logger.Debug("Auditing founder ...")
	id := c.PostForm("id")
	status := c.PostForm("status")
	uid, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf("ID format failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	founder, err := db.GetMasterByID(uint64(uid))
	if err != nil {
		logger.Errorf("Get founder info failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	err = db.UpdateMasterStatus(founder, status)
	if err != nil {
		logger.Errorf("Update founder status failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}
