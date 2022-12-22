package main

import (
	"os"

	gconfig "github.com/initia-labs/initia-apis/config"
	"github.com/initia-labs/initia-apis/router"
	"github.com/sirupsen/logrus"
)

func main() {
	configure := gconfig.Config()
	_, err := router.SetupRouter(configure)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
