package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error al cargar archivo .env: %v", err)
	}
	return nil
}

func GetEnvVariable(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("la variable de entorno '%s' no está configurada", key)
	}
	return value, nil
}
