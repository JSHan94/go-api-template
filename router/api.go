package router

import (
	"github.com/gin-gonic/gin"
	gcontroller "github.com/initia-labs/initia-apis/controller"
	"github.com/initia-labs/initia-apis/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func activateSwaggerAPI(engine *gin.Engine) {
	// Activate swagger API documentation
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func activateAPI(engine *gin.Engine) {
	// API:v1.0
	v1 := engine.Group("/v1/")
	{
		block := v1.Group("/block")
		{
			block.GET("", gcontroller.GetBlock)
			block.GET("/avgtime", gcontroller.GetBlockAvgTime)
		}

		blocks := v1.Group("/blocks")
		{
			blocks.GET("", gcontroller.GetBlocks)
		}

		tx := v1.Group("/tx")
		{
			tx.GET("", gcontroller.GetBlock)
			tx.GET("/mempool", gcontroller.GetMempoolTx)
		}

		txs := v1.Group("/txs")
		{
			txs.GET("", gcontroller.GetTxs)
			txs.GET("/mempool", gcontroller.GetMempoolTxs)
			txs.GET("/gasprice", gcontroller.GetGasPrice)
		}

		dashboard := v1.Group("/dashboard")
		{
			dashboard.GET("/tx/volume", gcontroller.GetTxVolume)
		}
	}
}
