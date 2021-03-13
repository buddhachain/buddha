package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/db/mongo"
	"github.com/gin-gonic/gin"
)

func UploadTemple(c *gin.Context) {
	logger.Debug("Uploading temple...")
	req := &mongo.Temple{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = mongo.InsertTemple(req)
	if err != nil {
		logger.Errorf("Db insert temple info error: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, err.Error())
		return
	}
	logger.Infof("Upload temple %s success.", req.UID)
	utils.Response(c, nil, define.Success, nil)
	return
}
