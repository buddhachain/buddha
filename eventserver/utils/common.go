package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

type ResponseInfo struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, err error, errCode int, data interface{}) {
	res := &ResponseInfo{
		ErrCode: errCode,
		Data:    data,
	}
	if err != nil {
		res.ErrMsg = err.Error()
	}
	ret, _ := json.Marshal(res)
	c.Writer.Write(ret)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	return
}
