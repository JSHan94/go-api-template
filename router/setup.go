package router

import (
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	gconfig "github.com/initia-labs/initia-apis/config"
)

func runRouter(configure *gconfig.Configuration, engine *gin.Engine) error {
	// Attaches the router to a http.Server and starts listening and serving HTTP requests
	var port string
	if gin.Mode() == gin.TestMode {
		port = configure.Server.TestPort
		if !isOpened(port) {
			return nil
		}
	} else {
		port = configure.Server.ServerPort
	}

	return engine.Run(":" + port)
}

// SetupRouter sets up all the routes
func SetupRouter(configure *gconfig.Configuration) (*gin.Engine, error) {
	r := gin.New()
	activateSwaggerAPI(r)
	activateAPI(r)
	if err := runRouter(configure, r); err != nil {
		panic(err)
	}
	return r, nil

}

func isOpened(port string) bool {
	timeout := 5 * time.Second
	host := "http://localhost"
	target := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}
