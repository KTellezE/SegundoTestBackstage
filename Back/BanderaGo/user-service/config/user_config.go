package config

import (
	"os"

	"github.com/joho/godotenv"
)

type UserConfig struct {
	DBUser          string
	DBPassword      string
	DBName          string
	DBPort          string
	DBTable         string
	ApplicationPort string
}

func NewUserConfig() (*UserConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	userConfig := &UserConfig{
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		DBPort:          os.Getenv("DB_PORT"),
		DBTable:         os.Getenv("DB_TABLE"),
		ApplicationPort: os.Getenv("APP_PORT"),
	}

	return userConfig, nil
}
