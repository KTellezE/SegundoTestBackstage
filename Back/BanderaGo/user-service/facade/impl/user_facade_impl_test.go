package impl

import (
	"application/dtos/input"
	"application/dtos/output"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de UserService para pruebas
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error) {
	args := m.Called(userIn)
	return args.Get(0).(output.CreateUserOut), args.Error(1)
}

func (m *MockUserService) GetUserByID(id uint) (output.GetUserOut, error) {
	args := m.Called(id)
	return args.Get(0).(output.GetUserOut), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]output.GetUsersOut, error) {
	args := m.Called()
	return args.Get(0).([]output.GetUsersOut), args.Error(1)
}

func (m *MockUserService) UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error) {
	args := m.Called(id, userIn)
	return args.Get(0).(output.UpdateUserOut), args.Error(1)
}

func (m *MockUserService) DeleteUser(id uint) (output.DeleteUserOut, error) {
	args := m.Called(id)
	return args.Get(0).(output.DeleteUserOut), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockUserService := new(MockUserService)
	userFacade := NewUserFacade(mockUserService)

	// Configurar expectativas en el mock
	mockUserService.On("CreateUser", mock.Anything).Return(output.CreateUserOut{}, nil)

	// Ejecutar la función a probar
	result, err := userFacade.CreateUser(input.CreateUserIn{})

	// Verificar resultado y expectativas en el mock
	assert.NotNil(t, result)
	assert.NoError(t, err)
	mockUserService.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockUserService := new(MockUserService)
	userFacade := NewUserFacade(mockUserService)

	// Configurar expectativas en el mock
	mockUserService.On("GetUserByID", uint(1)).Return(output.GetUserOut{}, nil)

	// Ejecutar la función a probar
	result, err := userFacade.GetUserByID(1)

	// Verificar resultado y expectativas en el mock
	assert.NotNil(t, result)
	assert.NoError(t, err)
	mockUserService.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	mockUserService := new(MockUserService)
	userFacade := NewUserFacade(mockUserService)

	// Datos simulados de usuarios
	mockUsers := []output.GetUsersOut{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	// Configurar expectativas en el mock
	mockUserService.On("GetAllUsers").Return(mockUsers, nil)

	// Ejecutar la función a probar
	usersOut, err := userFacade.GetAllUsers()

	// Verificar resultado y expectativas en el mock
	assert.NotNil(t, usersOut)
	assert.NoError(t, err)
	assert.Equal(t, len(mockUsers), len(usersOut))
	for i, user := range usersOut {
		assert.Equal(t, mockUsers[i].ID, user.ID)
		assert.Equal(t, mockUsers[i].Name, user.Name)
	}
	mockUserService.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockUserService := new(MockUserService)
	userFacade := NewUserFacade(mockUserService)

	// Configurar expectativas en el mock
	mockUserID := uint(1)
	mockUserIn := input.UpdateUserIn{Name: "Updated Name", LastName: "Updated Lastname"}
	mockUserService.On("UpdateUser", mockUserID, mockUserIn).Return(output.UpdateUserOut{}, nil)

	// Ejecutar la función a probar
	result, err := userFacade.UpdateUser(mockUserID, mockUserIn)

	// Verificar resultado y expectativas en el mock
	assert.NotNil(t, result)
	assert.NoError(t, err)
	mockUserService.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockUserService := new(MockUserService)
	userFacade := NewUserFacade(mockUserService)

	// Configurar expectativas en el mock
	mockUserService.On("DeleteUser", uint(1)).Return(output.DeleteUserOut{Success: true}, nil)

	// Ejecutar la función a probar
	result, err := userFacade.DeleteUser(1)

	// Verificar resultado y expectativas en el mock
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.True(t, result.Success)
	mockUserService.AssertExpectations(t)
}
