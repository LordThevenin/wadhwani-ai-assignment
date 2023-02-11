package config

type Configuration struct {
	Port int
}

// Singleton variable for configuration
var config *Configuration

func GetConfig() *Configuration {
	return config
}
