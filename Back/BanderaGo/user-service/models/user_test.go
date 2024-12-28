package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTableName(t *testing.T) {
	// Configuración de una variable de entorno de prueba
	testKey := "DB_TABLE"
	testValue := "test_value"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey) // Limpiar la variable de entorno al final de la prueba

	expectedTableName := "test_value"
	tableName := User{}.TableName()
	assert.Equal(t, expectedTableName, tableName)
}
