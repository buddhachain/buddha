package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
)

var logger = utils.NewLogger("debug", "factory/handler")

//获取账户余额
func GetBalance(c *gin.Context) {
	addr := c.Param("id")
	balance, err := xuper.GetBalance(addr)
	if err != nil {
		logger.Errorf("Get %s balance failed %s", addr, err.Error())
		utils.Response(c, err, define.EQueryFailed, nil)
		return
	}
	logger.Infof("Get account %s balance %s", addr, balance)
	utils.Response(c, nil, define.ESuccess, &balance)
	return
}

//根据txid获取交易详情
func GetTx(c *gin.Context) {
	tx := c.Param("id")
	res, err := xuper.GetTx(tx)
	if err != nil {
		logger.Errorf("Get %s transaction info failed %s", tx, err.Error())
		utils.Response(c, err, define.EQueryFailed, nil)
		return
	}
	logger.Infof("Get %s transaction info %+v", tx, res)
	utils.Response(c, nil, define.ESuccess, res)
	return
}
