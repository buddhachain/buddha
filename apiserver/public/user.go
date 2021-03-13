package public

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/db/mongo"
	"github.com/gin-gonic/gin"
)

var logger = utils.NewLogger("debug", "factory/public")

//新用户信息上传
func PostNewUser(c *gin.Context) {
	logger.Debug("Post new user ...")
	user := &mongo.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	logger.Infof("Request body info %v", user)
	err = mongo.InsertUser(user)
	if err != nil {
		logger.Errorf("Db insert user info error: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, err.Error())
		return
	}
	logger.Infof("Insert user %s success", user.Account)
	token := GenerateToken(user.Account)
	utils.Response(c, nil, define.Success, token)
	return
}

//根据钱包account获取用户信息
func GetUserInfo(c *gin.Context) {
	user := c.GetHeader("user")
	logger.Infof("Getting %s user info", user)
	info, err := mongo.GetUserByAccount(user)
	if err != nil {
		logger.Errorf("Get user info error: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, err.Error())
		return
	}
	logger.Infof("User info %v", info)
	utils.Response(c, nil, define.Success, info)
	return
}

//更新用户头像
func UpdateUserImage(c *gin.Context) {
	account := c.GetHeader("user")
	logger.Infof("Entering updating %s image", account)
	image := c.PostForm("image")
	err := mongo.UpdateUserImage(account, image)
	if err != nil {
		logger.Errorf("Update user image error: %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, err.Error())
		return
	}
	logger.Infof("Update user %s image to %s", account, image)
	utils.Response(c, nil, define.Success, nil)
	return
}

//更新用户昵称
func UpdateUserNickname(c *gin.Context) {
	account := c.GetHeader("user")
	logger.Infof("Entering updating %s image", account)
	name := c.PostForm("nickname")
	err := mongo.UpdateUserNickname(account, name)
	if err != nil {
		logger.Errorf("Update user nickname error: %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, err.Error())
		return
	}
	logger.Infof("Update user %s nickname to %s", account, name)
	utils.Response(c, nil, define.Success, nil)
	return
}
