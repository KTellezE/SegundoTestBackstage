package controllers

import (
	"application/dtos/input"
	"application/dtos/output"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockUserFacade es una implementación simulada de la interfaz UserFacade que devuelve objetos correctos
type MockUserFacade struct{}

func (m *MockUserFacade) CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error) {
	userOut := output.CreateUserOut{
		ID:       1,
		Name:     userIn.Name,
		LastName: userIn.LastName,
		// CreatedAt: time.Now(),
	}
	return userOut, nil
}
func (m *MockUserFacade) GetAllUsers() ([]output.GetUsersOut, error) {
	usersOut := []output.GetUsersOut{
		{ID: 1, Name: "John", LastName: "Doe"},
		{ID: 2, Name: "Jane", LastName: "Smith"},
	}
	return usersOut, nil
}
func (m *MockUserFacade) GetUserByID(id uint) (output.GetUserOut, error) {
	return output.GetUserOut{ID: 1, Name: "John", LastName: "Doe"}, nil
}
func (m *MockUserFacade) UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error) {
	userOut := output.UpdateUserOut{
		ID:       id,
		Name:     userIn.Name,
		LastName: userIn.LastName,
		// UpdatedAt: time.Now(),
	}
	return userOut, nil
}
func (m *MockUserFacade) DeleteUser(id uint) (output.DeleteUserOut, error) {
	return output.DeleteUserOut{Success: true}, nil
}

// MockUserFacadeError es una implementación simulada de la interfaz UserFacade que devuelve errores
type MockUserFacadeError struct{}

func (m *MockUserFacadeError) CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error) {
	return output.CreateUserOut{}, errors.New("create error")
}
func (m *MockUserFacadeError) GetAllUsers() ([]output.GetUsersOut, error) {
	return []output.GetUsersOut{}, errors.New("get list error")
}
func (m *MockUserFacadeError) GetUserByID(id uint) (output.GetUserOut, error) {
	return output.GetUserOut{}, errors.New("get first error")
}
func (m *MockUserFacadeError) UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error) {
	return output.UpdateUserOut{}, errors.New("update error")
}
func (m *MockUserFacadeError) DeleteUser(id uint) (output.DeleteUserOut, error) {
	return output.DeleteUserOut{}, errors.New("delete error")
}

// ---------------------Tests para CreateUser ---------------------
func TestCreateUser(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{} // Implementa tu propio mock de UserFacade
	userController := NewUserController(facadeMock)

	// Crear un usuario de prueba
	userIn := input.CreateUserIn{Name: "John", LastName: "Doe"}
	userOut := output.CreateUserOut{ID: 1, Name: "John", LastName: "Doe"}

	// Convertir el usuario de entrada a JSON
	userInJSON, _ := json.Marshal(userIn)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Ejecutar la función de controlador CreateUser
	userController.CreateUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedBody, _ := json.Marshal(userOut)
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestCreateUserErrorJson(t *testing.T) {
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear un usuario de prueba erroneo
	userInIncorrect := output.DeleteUserOut{Success: true}

	// Convertir el usuario de entrada a JSON
	userInJSON, _ := json.Marshal(userInIncorrect)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Ejecutar la función de controlador CreateUser
	userController.CreateUser(c)

	// Verificar el código de estado y el1 cuerpo de la respuesta
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorJson + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestCreateUserErrorCreation(t *testing.T) {
	facadeMock := &MockUserFacadeError{}
	userController := NewUserController(facadeMock)

	// Convertir el usuario de entrada a JSON
	userIn := input.CreateUserIn{Name: "John", LastName: "Doe"}

	// Convertir el usuario de entrada a JSON
	userInJSON, _ := json.Marshal(userIn)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Ejecutar la función de controlador CreateUser
	userController.CreateUser(c)

	// Verificar el código de estado y el1 cuerpo de la respuesta
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorCreation + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

// ---------------------Tests para GetAllUsers ---------------------
func TestGetAllUsers(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Ejecutar la función de controlador GetAllUsers
	userController.GetAllUsers(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `[{"id":1,"name":"John","last_name":"Doe"},{"id":2,"name":"Jane","last_name":"Smith"}]`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetAllUsersErrorGetList(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacadeError{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Ejecutar la función de controlador GetAllUsers
	userController.GetAllUsers(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorGetUsers + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

// ---------------------Tests para GetSingleUser ---------------------
func TestGetUserById(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.GetSingleUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"id":1,"name":"John","last_name":"Doe"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetUserByIdErrorInvalidID(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/api/users/notANumber", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "notANumber"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.GetSingleUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorID + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestGetUserByIdErrorGetFirstError(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacadeError{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.GetSingleUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusNotFound, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorUserNotFount + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

// ---------------------Tests para UpdateUser ---------------------
func TestUpdateUser(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	userIn := input.UpdateUserIn{
		Name:     "UpdatedName",
		LastName: "UpdatedLastName",
	}
	userInJSON, err := json.Marshal(userIn)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.UpdateUser(c)

	userOut := output.UpdateUserOut{ID: 1, Name: "UpdatedName", LastName: "UpdatedLastName"}
	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody, _ := json.Marshal(userOut)
	assert.JSONEq(t, string(expectedBody), w.Body.String())
}

func TestUpdateUserInvalidId(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	userIn := input.UpdateUserIn{
		Name:     "UpdatedName",
		LastName: "UpdatedLastName",
	}
	userInJSON, err := json.Marshal(userIn)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/notANumber", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "notANumber"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.UpdateUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorID + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestUpdateUserErrorJson(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	userInIncorrect := output.DeleteUserOut{Success: false}

	userInJSON, err := json.Marshal(userInIncorrect)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.UpdateUser(c)

	// Verificar el código de estado y el1 cuerpo de la respuesta
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorJson + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestUpdateUserUpdateError(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacadeError{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	userIn := input.UpdateUserIn{
		Name:     "UpdatedName",
		LastName: "UpdatedLastName",
	}
	userInJSON, err := json.Marshal(userIn)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(userInJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.UpdateUser(c)

	// Verificar el código de estado y el1 cuerpo de la respuesta
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorUpdateUser + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

// ---------------------Tests para DeleteUser ---------------------
func TestDeleteUser(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("DELETE", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.DeleteUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"success":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestDeleteUserErrorId(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacade{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("DELETE", "/api/users/notANumber", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "notANumber"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.DeleteUser(c)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorID + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}

func TestDeleteUserDeleteError(t *testing.T) {
	// Configurar el controlador y la fachada para las pruebas
	facadeMock := &MockUserFacadeError{}
	userController := NewUserController(facadeMock)

	// Crear una solicitud HTTP simulada
	req, err := http.NewRequest("DELETE", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un contexto de Gin con un grabador de respuesta simulado
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	// Ejecutar la función de controlador GetSingleUser
	userController.DeleteUser(c)

	// Verificar el código de estado y el1 cuerpo de la respuesta
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := "{\"error\":\"" + userController.constants.MessageErrorDeleteUser + "\"}"
	assert.Equal(t, string(expectedBody), w.Body.String())
}
