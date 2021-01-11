package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/gin-gonic/gin"
)

var logger = utils.NewLogger("debug", "factory/handler")

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
}
