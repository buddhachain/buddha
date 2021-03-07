package router

import (
	"sync"

	"github.com/buddhachain/buddha/apiserver/factory/handler"
	"github.com/buddhachain/buddha/apiserver/public"
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
		user := vGroup.Group("/user")
		{
			user.POST("", public.PostNewUser)
			//user.Use(public.TokenAuthMiddleware())
			user.GET("", public.GetUserInfo)
			user.PUT("/image", public.UpdateUserImage)
			user.PUT("/nickname", public.UpdateUserNickname)
		}
		//vGroup.Use(public.TokenAuthMiddleware())
		vGroup.GET("/balance/:id", handler.GetBalance)
		vGroup.GET("/detail/balance/:id", handler.GetBalanceDetail)
		vGroup.GET("/tx/:id", handler.GetTx)
		vGroup.GET("/txs/:id", handler.GetTxsInfo)

		vGroup.POST("/pretx", handler.PreExec)
		vGroup.POST("/tx", handler.PostRealTx)

		contract := vGroup.Group("/contract")
		{
			contract.POST("/invoke", handler.PreInvoke)   //合约通用invoke与执行接口
			contract.GET("/query", handler.ContractQuery) //合约通用查询接口
			//contract.POST("/post", handler.PostContractRealTx)    //合约通用调用接口
			contract.POST("", handler.PostContractRealTxAndParse) //合约通用查询接口
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
			ipfs.POST("/image", handler.SaveImages)
		}

		charge := vGroup.Group("/charge")
		{
			charge.POST("/gift/newcomer", handler.NewcomerCharge)
		}

		{
			vGroup.POST("/founder", handler.ApplyFounder)  //基金成员申请
			vGroup.PUT("/founder", handler.ApproveFounder) //审核申请

			vGroup.POST("/master", handler.ApplyMaster)  //申请成为法师
			vGroup.PUT("/master", handler.ApproveMaster) //审核申请

			vGroup.PUT("/kind/:id", handler.AddKind)   //善举入库
			vGroup.POST("/kind", handler.ApplyKind)    //善举上架申请
			vGroup.PUT("/kind", handler.ApproveKind)   //审核申请
			vGroup.DELETE("/kind", handler.DeleteKind) //删除善行

			vGroup.POST("/order", handler.CreatOrder) //创建新订单
		}

		blog := vGroup.Group("/blog")
		{
			blog.POST("", handler.PubBlog)
			blog.GET("", handler.GetBlogs)
			blog.DELETE("/:id", handler.DeleteBlog)
			blog.POST("/vote/:id", handler.VoteBlog)
			blog.PUT("/vote/:id", handler.CancelVoteBlog)
		}
		comment := vGroup.Group("comment")
		{
			comment.POST("", handler.CreateComment)
			comment.GET("", handler.GetComments) //根据动态id获取评论
			comment.DELETE("/:id", handler.DeleteComment)
		}

		//Q&A
		qa := vGroup.Group("/qa")
		{
			qa.GET("", handler.GetQAInfo)
			qa.POST("/issue", handler.NewQuestion)
			qa.DELETE("/issue/:id", handler.DeleteQuestion)
			qa.PATCH("/issue/vote/:id", handler.VoteIssue)
			qa.PUT("/issue/vote/:id", handler.CancelVoteIssue)

			qa.POST("/answer", handler.NewAnswer)
		}
		practise := vGroup.Group("/practise")
		{
			practise.POST("/sutra", handler.UploadSutra)
			practise.GET("/sutra", handler.GetSutraInfo)
			practise.GET("/category/:pid/sutra", handler.GetCategorySutrasInfo)
			practise.POST("/sutra/category", handler.UploadSutraCategory)
			practise.GET("/sutra/category", handler.GetSutraCategoryInfo)
			practise.PUT("/sutra/read", handler.NewSutraRead)
			practise.GET("/sutra/history", handler.GetSutraReadingHistory)
		}
		rc := vGroup.Group("/rc") //融云相关接口
		{
			rc.POST("/login", handler.NewRCToken)
		}
	}

}
