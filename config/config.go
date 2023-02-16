package config

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)

type Configuration struct {
	Port                     int    `mapstructure:"PORT"`
	DbUser                   string `mapstructure:"DB_USER"`
	DbPassword               string `mapstructure:"DB_PASSWORD"`
	DbHost                   string `mapstructure:"DB_HOST"`
	DbName                   string `mapstructure:"DB_NAME"`
	TranslationConfigFile    string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
	RedisHost                string `mapstructure:"REDIS_HOST"`
	RedisPort                string `mapstructure:"REDIS_PORT"`
	RedisDB                  int    `mapstructure:"REDIS_DB"`
	TranslationConfiguration []byte
}

// Singleton variable for configuration
var config *Configuration
var once sync.Once

func GetConfig() *Configuration {
	return config
}

func Init() {
	// initialize Config
	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			// Log error in reading config values
		}
		config = new(Configuration)
		err = viper.Unmarshal(config)
		if err != nil {
			// Log error in setting config
		}
		config.TranslationConfiguration = getTranslationConfigurationFromFile(config.TranslationConfigFile)
	})
}

func getTranslationConfigurationFromFile(file string) []byte {
	jsonFile, err := os.ReadFile(file)
	if err != nil {
		panic("Failed to obtain translation credentials from file with error: " + err.Error())
	}

	return jsonFile
}
