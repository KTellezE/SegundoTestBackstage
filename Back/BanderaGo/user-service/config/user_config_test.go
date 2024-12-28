package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewUserConfig testea la función NewUserConfig
func TestNewUserConfig(t *testing.T) {
	// Crea un archivo .env temporal en el paquete testing
	content := `
				DB_USER=test_user
				DB_PASSWORD=test_password
				DB_NAME=test_db
				DB_PORT=3306
				DB_TABLE=test_table
				APP_PORT=8080
				`

	err := os.WriteFile(".env", []byte(content), 0644) //Creando archivo temporal en el paquete para test
	require.NoError(t, err, "Failed to create .env file")
	defer os.Remove(".env") //Eliminar el archivo .env al final de la prueba

	// carga las variables de ambiente del archivo .env
	config, err := NewUserConfig()
	require.NoError(t, err, "Failed to load user config")

	// Verificar que la configuración se haya cargado correctamente
	assert.Equal(t, "test_user", config.DBUser, "DBUser should be set to 'test_user'")
	assert.Equal(t, "test_password", config.DBPassword, "DBPassword should be set to 'test_password'")
	assert.Equal(t, "test_db", config.DBName, "DBName should be set to 'test_db'")
	assert.Equal(t, "3306", config.DBPort, "DBPort should be set to '3306'")
	assert.Equal(t, "test_table", config.DBTable, "DBTable should be set to 'test_table'")
	assert.Equal(t, "8080", config.ApplicationPort, "ApplicationPort should be set to '8080'")
}

// Probar cuando no existe el archivo .env
func TestNewUserConfigMissingEnvFile(t *testing.T) {
	os.Remove(".env") // Se elimina el archivo .env del paquete

	_, err := NewUserConfig()
	assert.Error(t, err, "Expected error when .env file is missing")
}
