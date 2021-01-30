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

		contract := vGroup.Group("contract")
		{
			contract.POST("/invoke", handler.PreInvoke)        //合约通用invoke与执行接口
			contract.GET("/query", handler.ContractQuery)      //合约通用查询接口
			contract.POST("/post", handler.PostContractRealTx) //合约通用查询接口
		}

		exchange := vGroup.Group("/exchange")
		{
			exchange.GET("/product/:id", handler.GetProductByID)
			contract.POST("/product/post", handler.PostProductRealTx) //合约通用查询接口
			exchange.POST("/product", handler.PreAddProduct)
			exchange.POST("/product/delete", handler.PostDelProductRealTx)
		}

		ipfs := vGroup.Group("/ipfs")
		{
			ipfs.GET("/cat/:id", handler.CatIPFS)
		}

		charge := vGroup.Group("/charge")
		{
			charge.POST("/gift/newcomer", handler.NewcomerCharge)
		}
	}

}
