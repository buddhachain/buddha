package handler

import (
	"strconv"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/gin-gonic/gin"
)

//新增评论
func CreateComment(c *gin.Context) {
	logger.Debug("Entering comment ...")
	var comment db.Comment
	err := readBody(c, comment)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = db.InsertRow(&comment)
	if err != nil {
		logger.Errorf("Insert comment failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, nil)
	return
}

//格局 blog id获取comments
func GetComments(c *gin.Context) {
	logger.Debug("Getting blog comments...")
	id := c.Query("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	comments, err := db.GetCommentsByBlogID(uid)
	if err != nil {
		logger.Errorf("Get comments by blog id failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, comments)
	return
}

//删除评论
func DeleteComment(c *gin.Context) {
	logger.Debug("Deleting comment ...")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	user := c.GetHeader("user")
	err = db.DeleteComment(uid, user)
	if err != nil {
		logger.Errorf("delete comment failed %s", err.Error())
		utils.Response(c, err, define.DeleteDBErr, nil)
		return
	}
	logger.Infof("Delete comment %d success", uid)
	utils.Response(c, nil, define.Success, nil)
	return
}
