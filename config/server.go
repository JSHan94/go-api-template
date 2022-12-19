package config

import "os"

type ServerConfig struct {
	ServerPort string
	ServerEnv  string
}

// server - port and env
func server() *ServerConfig {
	config := &ServerConfig{}

	env()

	config.ServerPort = os.Getenv("APP_PORT")
	config.ServerEnv = os.Getenv("APP_ENV")

	return config
}
