package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host      string
	Port      string
	LogLevel  string
	StaticDir string
	ENV       string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "3008")
	viper.SetDefault("LOG_LEVEL", "INFO")
	viper.SetDefault("STATIC_DIR", "static")
	viper.SetDefault("ENV", "DEV")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	return Config{
		Host:      viper.GetString("HOST"),
		Port:      viper.GetString("PORT"),
		LogLevel:  viper.GetString("LOG_LEVEL"),
		StaticDir: viper.GetString("STATIC_DIR"),
		ENV:       viper.GetString("ENV"),
	}
}
