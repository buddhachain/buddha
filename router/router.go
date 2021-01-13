package router

import (
	"sync"

	"github.com/buddhachain/buddha/factory/handler"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var onceCreateRouter sync.Once

func GetRouter() *gin.Engine {
	onceCreateRouter.Do(func() {
		createRouter()
	})
	return router
}

func createRouter() {
	gin.ForceConsoleColor()
	router = gin.Default()

	vGroup := router.Group("/v1")
	{
		vGroup.GET("/balance/:id", handler.GetBalance)
		vGroup.GET("/tx/:id", handler.GetTx)

		vGroup.POST("/pretx", handler.PreExec)
		vGroup.POST("/tx", handler.PostRealTx)
	}
}
