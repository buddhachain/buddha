package handler

import (
	"encoding/json"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	FIND     = "find"
	ADD      = "add"
	EXCHANGE = "exchange"
)

func GetProductByID(c *gin.Context) {
	logger.Debug("Entering get production by id...")
	id := c.Param("id")
	addr := c.Query("addr")
	res, err := xuper.QueryWasmContract(addr, FIND, map[string]string{"id": id})
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
