package router

import (
	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
	gcontroller "github.com/initia-labs/initia-apis/controller"
	"github.com/initia-labs/initia-apis/docs"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter sets up all the routes
func SetupRouter(configure *gconfig.Configuration) (*gin.Engine, error) {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	// API:v1.0
	v1 := r.Group("/v1/")
	{
		block := v1.Group("/block")
		{
			block.GET("", gcontroller.GetBlock)
		}

		txs := v1.Group("/txs")
		{
			txs.GET("", gcontroller.GetTxs)
			txs.GET("/gasprice", gcontroller.GetGasPrice)
		}

		mempool := v1.Group("/mempool")
		{
			mempool.GET("/tx", gcontroller.GetMempoolTx)
			mempool.GET("/txs", gcontroller.GetMempoolTxs)
		}

		dashboard := v1.Group("/dashboard")
		{
			dashboard.GET("/tx/volume", gcontroller.GetTxVolume)
		}
	}

	// Activate swagger API documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Attaches the router to a http.Server and starts listening and serving HTTP requests
	err := r.Run(":" + configure.Server.ServerPort)
	if err != nil {
		logrus.Error(err)
	}

	return r, nil
}
