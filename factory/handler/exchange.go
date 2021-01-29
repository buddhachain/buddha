package handler

import (
	"encoding/hex"
	"encoding/json"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

const (
	FIND            = "find"
	ADD             = "add"
	EXCHANGE        = "exchange"
	NEWCOMERGIFTBAG = "1000000" //0.0001BUD
)

func GetProductByID(c *gin.Context) {
	logger.Debug("Entering get production by id...")
	id := c.Param("id")
	addr := c.Query("addr")
	res, err := xuper.QueryWasmContract(addr, EXCHANGE, FIND, map[string]string{"id": id})
	if err != nil {
		logger.Errorf("Get production by %s failed %s", id, err.Error())
		utils.Response(c, err, define.QueryContractErr, nil)
		return
	}
	logger.Infof("Get production by %s result %+v", addr, res)
	//pro := &define.Product{}
	//err = json.Unmarshal(res, pro)
	//if err != nil {
	//	logger.Errorf("Unmarshal res body failed %s", err.Error())
	//	utils.Response(c, err, define.UnmarshalErr, nil)
	//	return
	//}
	utils.Response(c, nil, define.Success, string(res))
	return
}

type addProductReq struct {
	*define.Product
	From string `json:"from"`
}

func PreAddProduct(c *gin.Context) {
	logger.Debug("Entering pre add production...")

	req := &addProductReq{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	//产品id由工具包生成uuid
	req.ID = xid.New().String()
	logger.Infof("Request info %+v", req)
	conReq, err := convertToMapString(req.Product)
	if err != nil {
		logger.Errorf("Convert request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	res, err := xuper.PreInvokeWasmContract(req.From, "0", EXCHANGE, ADD, conReq)
	if err != nil {
		logger.Errorf("Pre invoke wasm contract transaction failed %s", err.Error())
		utils.Response(c, err, define.PreInvokeWasmErr, nil)
		return
	}
	logger.Infof("Pre invoke wasm contract transaction result %+v", res)
	utils.Response(c, nil, define.Success, res)
	return
}

func PostProductRealTx(c *gin.Context) {
	logger.Debug("Entering post product real tx...")

	transaction := &pb.Transaction{}
	err, errCode := unmarshalProto(c, transaction)
	if err != nil {
		logger.Errorf("Parse request body to pb.Transaction failed: %s", err.Error())
		utils.Response(c, err, errCode, nil)
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
	txInfo, err := convertToProduction(transaction)
	if err != nil {
		logger.Errorf("Convert to production info failed: %s", err.Error())
		utils.Response(c, err, define.ConvertErr, nil)
		return
	}
	if err := db.InsertRow(txInfo); err != nil {
		logger.Errorf("Insert production info failed: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, &txid)
	return
}

func convertToMapString(info interface{}) (map[string]string, error) {
	byte, err := json.Marshal(info)
	if err != nil {
		return nil, errors.WithMessage(err, "marshal failed")
	}
	if byte == nil {
		return nil, nil
	}
	res := make(map[string]string)
	err = json.Unmarshal(byte, &res)
	return res, err
}

func convertToProduction(tx *pb.Transaction) (*db.Product, error) {
	txInfo := &db.Product{
		Initiator: tx.Initiator,
		TxId:      hex.EncodeToString(tx.Txid),
	}
	if len(tx.ContractRequests) == 0 {
		return txInfo, nil
	}
	args, err := json.Marshal(tx.ContractRequests[0].Args)
	if err != nil {
		return nil, err
	}
	pro := db.ProBase{}
	err = json.Unmarshal(args, &pro)
	if err != nil {
		return nil, err
	}
	txInfo.ProBase = pro
	return txInfo, nil
}
