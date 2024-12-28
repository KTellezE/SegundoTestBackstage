package impl

import (
	"errors"
	"testing"

	"application/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type GormDBMock struct {
	mock.Mock
}

// Create implements repositories.GormDB.
func (m *GormDBMock) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Delete implements repositories.GormDB.
func (m *GormDBMock) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}

// Find implements repositories.GormDB.
func (m *GormDBMock) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

// First implements repositories.GormDB.
func (m *GormDBMock) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

// Save implements repositories.GormDB.
func (m *GormDBMock) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func TestCreateUser(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	user := &models.User{Name: "John", LastName: "Doe"}

	mockDB.On("Create", user).Return(&gorm.DB{})

	err := repo.CreateUser(user)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	user := &models.User{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"}

	mockDB.On("First", mock.Anything, []interface{}{uint(1)}).Return(&gorm.DB{
		// Simular que no hubo error en la operación
		Error: nil,
	}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.User)
		*arg = *user
	})

	result, err := repo.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockDB.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	users := []*models.User{
		{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"},
		{Model: gorm.Model{ID: 2}, Name: "Jane", LastName: "Doe"},
	}

	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que no hubo error en la operación
		Error: nil,
	}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]*models.User)
		*arg = users
	})

	result, err := repo.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockDB.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	user := &models.User{Model: gorm.Model{ID: 1}, Name: "John", LastName: "Doe"}

	mockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que la búsqueda del usuario existente no generó errores y devolvió el usuario
		Error: nil,
	}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.User)
		*arg = *user
	})
	mockDB.On("Save", mock.Anything).Return(&gorm.DB{
		// Simular que la actualización del usuario no generó errores
		Error: nil,
	})

	err := repo.UpdateUser(1, user)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que no hubo error en la operación
		Error: nil,
	})

	err := repo.DeleteUser(1)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

// Test Errors
func TestGetUserByIDError(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	// Simular un error al obtener el usuario por ID
	mockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que la búsqueda del usuario por ID generó un error
		Error: errors.New("error getting user by ID"),
	})

	// Llamar al método GetUserByID
	_, err := repo.GetUserByID(1)

	// Verificar que se haya producido un error y que sea el esperado
	assert.Error(t, err)
	assert.EqualError(t, err, "error getting user by ID")

	mockDB.AssertExpectations(t)
}

func TestGetAllUsersError(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	// Simular un error al obtener todos los usuarios
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que la obtención de todos los usuarios generó un error
		Error: errors.New("error getting all users"),
	})

	// Llamar al método GetAllUsers
	_, err := repo.GetAllUsers()

	// Verificar que se haya producido un error y que sea el esperado
	assert.Error(t, err)
	assert.EqualError(t, err, "error getting all users")

	mockDB.AssertExpectations(t)
}

func TestUpdateUserError(t *testing.T) {
	mockDB := new(GormDBMock)
	repo := NewUserRepository(mockDB)

	// Simular un error al actualizar el usuario
	mockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{
		// Simular que la búsqueda del usuario generó un error
		Error: errors.New("error getting user"),
	})

	// Llamar al método UpdateUser
	err := repo.UpdateUser(1, &models.User{})

	// Verificar que se haya producido un error y que sea el esperado
	assert.Error(t, err)
	assert.EqualError(t, err, "error getting user")

	mockDB.AssertExpectations(t)
}
