package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host       string
	Port       string
	LogLevel   string
	StaticDir  string
	ENV        string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
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
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5420")
	viper.SetDefault("DB_USER", "admin")
	viper.SetDefault("DB_PASSWORD", "foobar")
	viper.SetDefault("DB_NAME", "hgmx")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	return Config{
		Host:       viper.GetString("HOST"),
		Port:       viper.GetString("PORT"),
		LogLevel:   viper.GetString("LOG_LEVEL"),
		StaticDir:  viper.GetString("STATIC_DIR"),
		ENV:        viper.GetString("ENV"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),
	}
}
