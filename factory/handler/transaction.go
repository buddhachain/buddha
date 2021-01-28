package handler

import (
	"io/ioutil"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func unmarshalProto(c *gin.Context, pb proto.Message) (error, int) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err, define.ReadRequestBodyErr
	}
	err = proto.Unmarshal(body, pb)
	if err != nil {
		return err, define.UnmarshalErr
	}
	return nil, define.Success
}

func NewcomerCharge(c *gin.Context) {
	addr := c.PostForm("addr")
	if addr == "" {
		logger.Error("Request account err")
		utils.Response(c, define.ErrRequest, define.RequestErr, nil)
		return
	}
	isNewcomer, err := db.IsNewcomer(addr)
	if err != nil {
		logger.Errorf("Get newcomer db info failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	if !isNewcomer {
		logger.Errorf("Account %s is not newcomer", addr)
		utils.Response(c, define.ErrRight, define.RightErr, nil)
		return
	}
	txid, err := xuper.Recharge(addr, NEWCOMERGIFTBAG)
	if err != nil {
		logger.Errorf("Newcomer git bag recharge failed %s", err.Error())
		utils.Response(c, err, define.PostTxErr, nil)
		return
	}
	err = db.InsertRow(&db.NewBag{Addr: addr, TxId: txid})
	if err != nil {
		logger.Errorf("Insert db failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	err = db.InsertTxInfo(&db.Transaction{
		From:   xuper.RootAddr,
		To:     addr,
		Amount: NEWCOMERGIFTBAG,
		TxId:   txid,
	})
	if err != nil {
		logger.Errorf("Insert db failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	logger.Infof("Account %s get newcomer git success; txid: %s", addr, txid)
	utils.Response(c, nil, 0, txid)
	return
}
