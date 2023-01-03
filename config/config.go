package config

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var configAll *Configuration

type Configuration struct {
	Database *DatabaseConfig
	Server   *ServerConfig
	Chain    *ChainConfig
}

func GetConfig() *Configuration {
	if configAll == nil {
		return Config()
	}
	return configAll
}

// env - load the configurations from .env
func env() {
	// Load environment variables
	projectName := regexp.MustCompile(`^(.*` + PROJECT_DIR_NAME + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.WithError(err).Panic("failed to load .env file")
	}
}

// Config - load all the configurations
func Config() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.Server = server()
	configuration.Chain = chain()

	configAll = &configuration

	return configAll
}
