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
		vGroup.GET("/detail/balance/:id", handler.GetBalanceDetail)
		vGroup.GET("/tx/:id", handler.GetTx)
		vGroup.GET("/txs/:id", handler.GetTxsInfo)

		vGroup.POST("/pretx", handler.PreExec)
		vGroup.POST("/tx", handler.PostRealTx)

		exchange := vGroup.Group("/exchange")
		{
			exchange.GET("/product/:id", handler.GetProductByID)
			exchange.POST("/product", handler.PreAddProduct)
		}

	}

}
