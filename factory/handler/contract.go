package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
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
