package handler

import (
	"strconv"

	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/gin-gonic/gin"
)

// 善举上架申请
//func KindApprove() (error, int) {
//
//}

// 善举入库
func AddKind(c *gin.Context) {
	logger.Debug("Add kind count ...")
	id := c.Param("id")
	count := c.PostForm("count")
	amount, err := strconv.Atoi(count)
	if err != nil {
		logger.Errorf("Param err %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	err = db.AddKind(id, uint64(amount))
	if err != nil {
		logger.Errorf("Add kind amount failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

//善行上架申请
func ApplyKind(c *gin.Context) {
	logger.Debug("Entering kind applying...")

	req := &db.Kind{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.ReadRequestBodyErr, nil)
		return
	}
	logger.Infof("Request info %+v", req)
	err = db.InsertRow(req)
	if err != nil {
		logger.Errorf("Insert kind apply failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

//审核善行上架信息
func ApproveKind(c *gin.Context) {
	logger.Debug("Auditing kind ...")
	id := c.PostForm("id")
	status := c.PostForm("status")
	founder, err := db.GetKindByID(id)
	if err != nil {
		logger.Errorf("Get kind info failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	err = db.UpdateKindStatus(founder, status)
	if err != nil {
		logger.Errorf("Update kind status failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}

func DeleteKind(c *gin.Context) {
	logger.Debug("Deleting kind ...")
	id := c.PostForm("id")
	err := db.DeleteKind(id)
	if err != nil {
		logger.Errorf("Delete kind info failed %s", err.Error())
		utils.Response(c, err, define.DeleteDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, nil)
	return
}
