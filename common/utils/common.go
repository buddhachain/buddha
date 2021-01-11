package utils

import (
	"encoding/json"
	"os"

	"github.com/buddhachain/buddha/common/define"
	"github.com/gin-gonic/gin"
)

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func Response(c *gin.Context, err error, errCode int, data interface{}){
	res := &define.ResponseInfo{
		ErrCode: errCode,
		Data: data,
	}
	if err != nil {
		res.ErrMsg = err.Error()
	}
	ret, _ := json.Marshal(res)
	c.Writer.Write(ret)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
}
