package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var configAll *Configuration

type Configuration struct {
	Database *DatabaseConfig
	Server   *ServerConfig
}

func GetConfig() *Configuration {
	return configAll
}

// env - load the configurations from .env
func env() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Panic("failed to load .env file")
	}
}

// Config - load all the configurations
func Config() *Configuration {
	var configuration Configuration

	configuration.Database = database()
	configuration.Server = server()

	configAll = &configuration

	return configAll
}
