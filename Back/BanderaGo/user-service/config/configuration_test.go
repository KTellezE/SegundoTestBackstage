package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Caso de prueba: archivo .env se carga correctamente
func TestLoadEnvVariablesOk(t *testing.T) {
	err := os.WriteFile(".env", []byte("KEY=value\n"), 0644) //Creando archivo temporal en el paquete para test
	require.NoError(t, err, "Failed to create .env file")
	defer os.Remove(".env")

	err = LoadEnvVariables()
	assert.Nil(t, err, "Expected no error when .env file is loaded")
}

// Caso de prueba: archivo .env no se carga
func TestLoadEnvVariablesNonEnvFile(t *testing.T) {
	os.Rename(".env", ".env.bak")       // Renombrar temporalmente el archivo .env para simular su ausencia
	defer os.Rename(".env.bak", ".env") // Restaurar el archivo .env al final de la prueba

	err := LoadEnvVariables()
	assert.NotNil(t, err, "Expected error when .env file is not found")
	assert.Contains(t, err.Error(), "error al cargar archivo .env", "Expected error message to mention failure to load .env file")
}

// Caso de prueba: variable de entorno existe
func TestGetEnvVariableOk(t *testing.T) {
	// Configuración de una variable de entorno de prueba
	testKey := "TEST_KEY"
	testValue := "test_value"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey) // Limpiar la variable de entorno al final de la prueba

	value, err := GetEnvVariable(testKey)
	assert.Nil(t, err, "Expected no error when environment variable exists")
	assert.Equal(t, testValue, value, "Expected value to match the set environment variable")
}

// Caso de prueba: variable de entorno no existe
func TestGetEnvVariableMissingKey(t *testing.T) {
	value, err := GetEnvVariable("NON_EXISTENT_KEY")
	assert.NotNil(t, err, "Expected error when environment variable does not exist")
	assert.Contains(t, err.Error(), "la variable de entorno 'NON_EXISTENT_KEY' no está configurada", "Expected error message to mention missing environment variable")
	assert.Empty(t, value, "Expected value to be empty when environment variable does not exist")
}
