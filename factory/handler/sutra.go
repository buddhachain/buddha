package handler

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/db/mongo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//上传佛经分类
func UploadSutraCategory(c *gin.Context) {
	logger.Debug("Uploading sutra category...")
	category := &mongo.Category{}
	err := readBody(c, category)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = mongo.InsertCategory(category)
	if err != nil {
		logger.Errorf("Db insert sutra category info error: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, err.Error())
		return
	}
	logger.Infof("Upload sutra category %s success.", category.Name)
	utils.Response(c, nil, define.Success, &category.ID)
	return
}

//获取佛经分类信息
func GetSutraCategoryInfo(c *gin.Context) {
	logger.Debug("Getting sutra info...")
	categories, err := mongo.GetSutraCategories()
	if err != nil {
		logger.Errorf("Get sutra categories info failed: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	logger.Infof("Sutra categories info %v", categories)
	utils.Response(c, nil, define.Success, categories)
	return
}

//上传佛经
func UploadSutra(c *gin.Context) {
	logger.Debug("Uploading sutra...")
	req := &mongo.Sutra{}
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = mongo.InsertSutra(req)
	if err != nil {
		logger.Errorf("Db insert sutra info error: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, err.Error())
		return
	}
	logger.Infof("Upload sutra %s success.", req.Name)
	utils.Response(c, nil, define.Success, nil)
	return
}

func GetSutraInfo(c *gin.Context) {
	logger.Debug("Getting sutra info...")
	id := c.Query("id")
	name := c.Query("name")
	if id == "" && name == "" {
		err := errors.New("id and name cannot be both empty ")
		logger.Errorf("Get sutra info failed: %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	var sutra *mongo.Sutra
	var err error
	if id != "" {
		sutra, err = mongo.GetSutraByID(id)
	} else {
		sutra, err = mongo.GetSutraByName(name)
	}
	if err != nil {
		logger.Errorf("Get sutra info failed: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	logger.Infof("Get sutra %s info success.", sutra.ID.Hex())
	utils.Response(c, nil, define.Success, sutra)
	return
}

//根据佛经分类获取sutras信息
func GetCategorySutrasInfo(c *gin.Context) {
	logger.Debug("Getting sutras info...")
	pid := c.Param("pid")
	sutras, err := mongo.GetSutrasByPid(pid)
	if err != nil {
		logger.Errorf("Get sutras info failed: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	utils.Response(c, nil, define.Success, sutras)
	return
}
