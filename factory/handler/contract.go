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
	"github.com/xuperchain/xuper-sdk-go/pb"
)

func PreInvoke(c *gin.Context) {
	logger.Debug("Entering pre invoke contract...")

	req := &define.InvokeInfo{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	logger.Infof("Request info %+v", req)

	res, err := xuper.PreInvokeWasmContract(req.From, req.Amount, req.ContractName, req.Method, req.Args)
	if err != nil {
		logger.Errorf("Pre invoke wasm contract transaction failed %s", err.Error())
		utils.Response(c, err, define.PreInvokeWasmErr, nil)
		return
	}
	logger.Infof("Pre invoke wasm contract transaction result %+v", res)
	utils.Response(c, nil, define.Success, res)
	return
}

//合约通用查询接口
func ContractQuery(c *gin.Context) {
	logger.Debug("Entering pre invoke contract...")

	req := &define.InvokeInfo{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	logger.Infof("Request info %+v", req)

	res, err := xuper.QueryWasmContract(req.From, req.ContractName, req.Method, req.Args)
	if err != nil {
		logger.Errorf("Query wasm contract transaction failed %s", err.Error())
		utils.Response(c, err, define.PreInvokeWasmErr, nil)
		return
	}
	logger.Infof("Query wasm contract transaction result %+v", res)
	utils.Response(c, nil, define.Success, res)
	return
}

//合约通用上传合约接口
func PostContractRealTx(c *gin.Context) {
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
	txInfo, err := convertToContractTx(transaction)
	if err != nil {
		logger.Errorf("Convert to contract tx info failed: %s", err.Error())
		utils.Response(c, err, define.ConvertErr, nil)
		return
	}
	if err := db.InsertContractTx(txInfo); err != nil {
		logger.Errorf("Insert contract tx info failed: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, &txid)
	return
}

func convertToContractTx(tx *pb.Transaction) (*db.ContractTx, error) {
	txInfo := &db.ContractTx{
		TxBase: db.TxBase{TxId: hex.EncodeToString(tx.Txid), Initiator: tx.Initiator},
	}
	if len(tx.ContractRequests) == 0 {
		return txInfo, nil
	}
	txInfo.Amount = tx.ContractRequests[0].Amount
	txInfo.ContractName = tx.ContractRequests[0].ContractName
	txInfo.MethodName = tx.ContractRequests[0].MethodName
	args, err := json.Marshal(tx.ContractRequests[0].Args)
	if err != nil {
		return nil, err
	}
	txInfo.Args = string(args)
	return txInfo, nil
}
