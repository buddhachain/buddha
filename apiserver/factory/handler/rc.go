package handler

import (
	"github.com/buddhachain/buddha/apiserver/factory/rc"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/db/mongo"
	"github.com/gin-gonic/gin"
)

func NewRCToken(c *gin.Context) {
	user := c.GetHeader("user")
	logger.Infof("user id %s", user)
	info, err := mongo.GetUserByAccount(user)
	if err != nil {
		logger.Errorf("Get user %s failed %s", user, err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	token, err := rc.RCUserRegister(user, info.Nickname, info.Image)
	if err != nil {
		logger.Errorf("Register user %s token failed %s", user, err.Error())
		utils.Response(c, err, define.RegisterTokenErr, nil)
		return
	}
	utils.Response(c, nil, 0, token)
	return
}
