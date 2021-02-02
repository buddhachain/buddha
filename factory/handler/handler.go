package handler

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

var logger = utils.NewLogger("debug", "factory/handler")

//获取账户余额
func GetBalance(c *gin.Context) {
	addr := c.Param("id")
	balance, err := xuper.GetBalance(addr)
	if err != nil {
		logger.Errorf("Get %s balance failed %s", addr, err.Error())
		utils.Response(c, err, define.QueryErr, nil)
		return
	}
	logger.Infof("Get account %s balance %s", addr, balance)
	utils.Response(c, nil, define.Success, &balance)
	return
}

func GetBalanceDetail(c *gin.Context) {
	addr := c.Param("id")
	balance, err := xuper.GetBalanceDetail(addr)
	if err != nil {
		logger.Errorf("Get %s balance failed %s", addr, err.Error())
		utils.Response(c, err, define.QueryErr, nil)
		return
	}
	logger.Infof("Get account %s balance detail %+v", addr, balance)
	utils.Response(c, nil, define.Success, &balance)
	return
}

//根据txid获取交易详情
func GetTx(c *gin.Context) {
	tx := c.Param("id")
	res, err := xuper.GetTx(tx)
	if err != nil {
		logger.Errorf("Get %s transaction info failed %s", tx, err.Error())
		utils.Response(c, err, define.QueryErr, nil)
		return
	}
	logger.Infof("Get %s transaction info %+v", tx, res)
	utils.Response(c, nil, define.Success, res)
	return
}

//交易预处理
func PreExec(c *gin.Context) {
	logger.Debug("Entering pre exec...")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	req := &define.PreReqInfo{}
	err = json.Unmarshal(body, req)
	if err != nil {
		logger.Errorf("Unmarshal request body failed: %s", err.Error())
		utils.Response(c, err, define.UnmarshalErr, nil)
		return
	}
	logger.Infof("Request info %+v", req)
	res, err := xuper.PreExec(req.Desc, req.Amount, "0", req.Account, "")
	if err != nil {
		logger.Errorf("Pre exec transaction failed %s", err.Error())
		utils.Response(c, err, define.PreExecErr, nil)
		return
	}
	logger.Infof("Pre exec transaction result %+v", res)
	utils.Response(c, nil, define.Success, res)
	return
}

//将已签名的tx上传至链上，返回txid
func PostRealTx(c *gin.Context) {
	logger.Debug("Entering post real tx...")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	transaction := &pb.Transaction{}
	err = proto.Unmarshal(body, transaction)
	if err != nil {
		logger.Errorf("Unmarshal request body to pb.Transaction failed: %s", err.Error())
		utils.Response(c, err, define.UnmarshalErr, nil)
		return
	}
	logger.Infof("Request info %+v", transaction)
	txid, err := xuper.PostRealTx(transaction)
	if err != nil {
		logger.Errorf("Post real tx failed: %s", err.Error())
		utils.Response(c, err, define.PostTxErr, nil)
		return
	}
	logger.Info("Post tx: %s success", txid)
	txInfo := GetTxInfo(transaction)
	if err := db.InsertTxInfo(txInfo); err != nil {
		logger.Errorf("Insert tx info failed: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, &txid)
	return
}

//获取账户交易信息
func GetTxsInfo(c *gin.Context) {
	logger.Debug("Entering getting historical transaction information...")
	addr := c.Param("id")
	logger.Infof("Get %s txs", addr)
	txs, err := db.GetTxsByAddr(addr, 0)
	if err != nil {
		logger.Errorf("Get txs info failed: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	logger.Infof("Get %s txs: %+v", addr, txs)
	utils.Response(c, nil, define.Success, &txs)
	return
}

func readBody(c *gin.Context, req interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return errors.WithMessage(err, "read request body failed")
	}
	return json.Unmarshal(body, req)
}

func GetTxInfo(tx *pb.Transaction) *db.Transaction {
	txInfo := &db.Transaction{
		TxBase: db.TxBase{TxId: hex.EncodeToString(tx.Txid), Initiator: tx.Initiator,
			Timestamp: tx.Timestamp},
	}
	for _, txOutPut := range tx.TxOutputs {
		addr := txOutPut.ToAddr
		if string(addr) != tx.Initiator {
			txInfo.To = string(addr)
			amountBigInt := xuper.FromAmountBytes(txOutPut.Amount)
			txInfo.Amount = amountBigInt.String()
			break
		}
	}
	return txInfo
}
