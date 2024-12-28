package impl

import (
	"application/dtos/input"
	"application/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock de UserRepository
type MockUserRepository struct {
	mock.Mock
}

// Implementación de CreateUser para el mock
func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Implementación de GetUserByID para el mock
func (m *MockUserRepository) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

// Implementación de GetAllUsers para el mock
func (m *MockUserRepository) GetAllUsers() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

// Implementación de UpdateUser para el mock
func (m *MockUserRepository) UpdateUser(id uint, user *models.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

// Implementación de DeleteUser para el mock
func (m *MockUserRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test para CreateUser en UserServiceImpl
func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userIn := input.CreateUserIn{
		Name:     "John",
		LastName: "Doe",
	}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("CreateUser", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*models.User)
		user.ID = 1 // Establecer un ID válido para el usuario creado
	})

	// Ejecutar el método CreateUser del servicio
	userOut, err := service.CreateUser(userIn)

	// Verificar que no se produzca un error
	assert.NoError(t, err)
	// Verificar que el ID del usuario de salida no sea cero
	assert.NotEqual(t, uint(0), userOut.ID)

	mockRepo.AssertExpectations(t)
}

// Test para GetUserByID en UserServiceImpl
func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userID := uint(1)
	user := &models.User{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("GetUserByID", userID).Return(user, nil)

	// Ejecutar el método GetUserByID del servicio
	userOut, err := service.GetUserByID(userID)

	// Verificar que no se produzca un error
	assert.NoError(t, err)
	// Verificar que el usuario de salida tenga los mismos datos que el usuario simulado
	assert.Equal(t, user.ID, userOut.ID)
	assert.Equal(t, user.Name, userOut.Name)
	assert.Equal(t, user.LastName, userOut.LastName)

	mockRepo.AssertExpectations(t)
}

// Test para GetAllUsers en UserServiceImpl
func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	users := []*models.User{
		{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"},
		{Model: gorm.Model{ID: 2}, Name: "Jane", LastName: "Smith"},
	}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("GetAllUsers").Return(users, nil)

	// Ejecutar el método GetAllUsers del servicio
	usersOut, err := service.GetAllUsers()

	// Verificar que no se produzca un error
	assert.NoError(t, err)
	// Verificar que la cantidad de usuarios de salida sea la misma que la cantidad de usuarios simulados
	assert.Len(t, usersOut, len(users))

	mockRepo.AssertExpectations(t)
}

// Test para UpdateUser en UserServiceImpl
func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userID := uint(1)
	userIn := input.UpdateUserIn{Name: "John Updated", LastName: "Doe Updated"}
	user := &models.User{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"}

	// Configurar el comportamiento esperado del mock
	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", userID, mock.AnythingOfType("*models.User")).Return(nil)

	// Ejecutar el método UpdateUser del servicio
	userOut, err := service.UpdateUser(userID, userIn)

	// Verificar que no se produzca un error
	assert.NoError(t, err)
	// Verificar que los datos del usuario de salida sean los esperados
	assert.Equal(t, userIn.Name, userOut.Name)
	assert.Equal(t, userIn.LastName, userOut.LastName)

	mockRepo.AssertExpectations(t)
}

// Test para DeleteUser en UserServiceImpl
func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userID := uint(1)

	// Configurar el comportamiento esperado del mock
	mockRepo.On("DeleteUser", userID).Return(nil)

	// Ejecutar el método DeleteUser del servicio
	deleteOut, err := service.DeleteUser(userID)

	// Verificar que no se produzca un error
	assert.NoError(t, err)
	// Verificar que el resultado de eliminación sea exitoso
	assert.True(t, deleteOut.Success)

	mockRepo.AssertExpectations(t)
}

// Testear los errores
func TestCreateUserError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	// Configurar el error que quieres simular
	expectedErr := errors.New("error creating user")

	// Configurar el comportamiento del mock para devolver un error
	mockRepo.On("CreateUser", mock.Anything).Return(expectedErr)

	// Llamar al método CreateUser del servicio
	_, err := userService.CreateUser(input.CreateUserIn{Name: "John", LastName: "Doe"})

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}

func TestGetUserByIDError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	// Configurar el error que quieres simular
	expectedErr := errors.New("error getting user by ID")

	// Configurar el comportamiento del mock para devolver un error
	mockRepo.On("GetUserByID", mock.Anything).Return(&models.User{}, expectedErr)

	// Llamar al método GetUserByID del servicio
	_, err := userService.GetUserByID(1)

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}

func TestGetAllUsersError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	// Configurar el error que quieres simular
	expectedErr := errors.New("error getting all users")

	users := []*models.User{}
	// Configurar el comportamiento del mock para devolver un error
	mockRepo.On("GetAllUsers").Return(users, expectedErr)

	// Llamar al método GetAllUsers del servicio
	_, err := userService.GetAllUsers()

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}

func TestUpdateUserError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	// Configurar el error que quieres simular al obtener el usuario
	expectedErr := errors.New("error getting user by ID")

	user := &models.User{}

	// Configurar el comportamiento del mock para devolver un error al obtener el usuario por ID
	mockRepo.On("GetUserByID", uint(1)).Return(user, expectedErr)

	// Llamar al método UpdateUser del servicio
	_, err := userService.UpdateUser(1, input.UpdateUserIn{Name: "John Updated", LastName: "Doe Updated"})

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}

func TestDeleteUserError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	// Configurar el error que quieres simular
	expectedErr := errors.New("error deleting user")

	// Configurar el comportamiento del mock para devolver un error
	mockRepo.On("DeleteUser", mock.Anything).Return(expectedErr)

	// Llamar al método DeleteUser del servicio
	_, err := userService.DeleteUser(1)

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}

func TestUpdateUserUpdateError(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := &MockUserRepository{}

	// Configurar el servicio con el mock
	userService := NewUserService(mockRepo)

	userID := uint(1)
	user := &models.User{}

	// Configurar el comportamiento del mock para devolver un error al obtener el usuario
	expectedErr := errors.New("error getting user")
	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", userID, mock.AnythingOfType("*models.User")).Return(expectedErr)

	// Llamar al método UpdateUser del servicio
	_, err := userService.UpdateUser(1, input.UpdateUserIn{Name: "John Updated", LastName: "Doe Updated"})

	// Verificar que se haya devuelto el error esperado
	assert.EqualError(t, err, expectedErr.Error())
}
