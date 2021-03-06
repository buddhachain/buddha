package handler

import (
	"encoding/json"

	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/gin-gonic/gin"
)

//未考虑链，暂未考虑资金冻结
func ApplyFounder(c *gin.Context) {
	logger.Debug("Entering founder applying...")

	req := &db.Founder{}
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

func ApproveFounder(c *gin.Context) {
	logger.Debug("Auditing founder ...")
	id := c.PostForm("id")
	status := c.PostForm("status")
	founder, err := db.GetFounderByID(id)
	if err != nil {
		logger.Errorf("Get founder info failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	err = db.UpdateFounderStatus(founder, status)
	if err != nil {
		logger.Errorf("Update founder status failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

func applyFounder(amount, initiator string, args []byte) (error, int) {
	//txInfo := &db.Product{Initiator: tx.Initiator}
	info := db.Founder{}
	err := json.Unmarshal(args, &info)
	if err != nil {
		return err, define.UnmarshalErr
	}
	//info.Name = initiator
	info.ID = initiator
	info.Guaranty = amount
	info.Status = 1
	err = db.InsertRow(&info)
	if err != nil {
		return err, define.InsertDBErr
	}
	return nil, 0
}

func commentFounder(args map[string][]byte) (error, int) {
	id, ok := args["id"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	status, ok := args["status"]
	if !ok {
		return define.ErrParam, define.ParamErr
	}
	founder, err := db.GetFounderByID(string(id))
	if err != nil {
		return err, define.QueryDBErr
	}
	err = db.UpdateFounderStatus(founder, string(status))
	if err != nil {
		return err, define.UpdateDBErr
	}
	return nil, 0
}
