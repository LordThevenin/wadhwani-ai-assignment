package config

import (
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port int
}

// Singleton variable for configuration
var config *Configuration

func GetConfig() *Configuration {
	return config
}

func Init() {
	// initialize Config
	_ = godotenv.Load()
}
