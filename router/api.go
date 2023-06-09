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
			block.GET("/latest", gcontroller.GetBlockLatest)
			block.GET("/height/:height", gcontroller.GetBlockByHeight)
			block.GET("/time/:time", gcontroller.GetBlockByTime)
			block.GET("/hash/:hash", gcontroller.GetBlockByHash)

			block.GET("/avgtime/:height", gcontroller.GetBlockAvgTime)
		}

		blocks := v1.Group("/blocks")
		{
			blocks.GET("/:from/:to", gcontroller.GetBlocksFromTo)
		}

		tx := v1.Group("/tx")
		{
			tx.GET("/:hash", gcontroller.GetTxByHash)
			// tx.GET("/mempool", gcontroller.GetMempoolTx)
		}

		txs := v1.Group("/txs")
		{
			txs.GET("/height/:height", gcontroller.GetTxsByHeight)
			txs.GET("/account/:account", gcontroller.GetTxsByAccount)
			txs.GET("/offset/:offset", gcontroller.GetTxsByOffset)

			txs.GET("/mempool", gcontroller.GetMempoolTxs)
			txs.GET("/gasprice", gcontroller.GetGasPrice)
		}

		analysis := v1.Group("/analysis")
		{
			analysis.GET("/tx-volume/:start/:end", gcontroller.GetTxVolume)
			analysis.GET("/tx-count/:start/:end", gcontroller.GetTxCount)
		}
	}
}
