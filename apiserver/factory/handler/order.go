package handler

import (
	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

//创建新订单
func CreatOrder(c *gin.Context) {
	logger.Debug("Creating order...")
	req := &db.Order{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	req.ID = xid.New().String()
	logger.Infof("Request info %+v", req)
	err = db.InsertRow(req)
	if err != nil {
		logger.Errorf("Insert order failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

//取消订单
func CancelOrder(c *gin.Context) {
	logger.Debug("Canceling order...")
	id := c.PostForm("id")
	order, err := db.GetOrderByID(id)
	if err != nil {
		logger.Errorf("Get order failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	err = db.CancelOrder(order)
	if err != nil {
		logger.Errorf("Cancel order failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}
