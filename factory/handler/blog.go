package handler

import (
	"strconv"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/gin-gonic/gin"
)

const CONTENT = "content"

//发布博客信息
func PubBlog(c *gin.Context) {
	logger.Debug("Blogging ...")
	user := c.GetHeader("user")
	content := c.PostForm("content")
	if content == "" {
		logger.Error("Content is nil")
		utils.Response(c, define.ErrParam, define.ParamErr, nil)
		return
	}
	cid, err := AddContent([]byte(content))
	if err != nil {
		logger.Error("Ipfs add failed %s", err.Error())
		utils.Response(c, err, define.IpfsAddErr, nil)
		return
	}
	err = db.InsertRow(&db.Blog{CID: cid, Sender: user})
	if err != nil {
		logger.Error("Insert blog %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	utils.Response(c, nil, 0, &cid)
	return
}

//blog 查询自己发的blog
func GetBlogs(c *gin.Context) {
	logger.Debug("Getting blogs...")
	user := c.GetHeader("user")
	page := c.Query("page")
	p, err := strconv.Atoi(page)
	if err != nil {
		logger.Error("Param page %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	if p < 1 {
		p = 1
	}
	count := c.Query("count")
	ucount, err := strconv.Atoi(count)
	if err != nil {
		logger.Error("Param count %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	if ucount < 2 {
		ucount = 2
	}
	blogs, err := db.GetBlogs(user, p, ucount)
	if err != nil {
		logger.Error("Get blogs %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, blogs)
	return
}

// delete blog
func DeleteBlog(c *gin.Context) {
	logger.Debug("Deleting blog ...")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	user := c.GetHeader("user")
	err = db.DeleteBlog(uid, user)
	if err != nil {
		logger.Errorf("delete blog failed %s", err.Error())
		utils.Response(c, err, define.DeleteDBErr, nil)
		return
	}
	logger.Infof("Delete blog %d success", uid)
	utils.Response(c, nil, define.Success, nil)
	return
}
