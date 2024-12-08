package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Database string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return Config{
		Port:     viper.GetString("PORT"),
		Database: viper.GetString("DATABASE_URL"),
	}
}
