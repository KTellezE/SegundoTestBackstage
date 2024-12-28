package repositories

import "application/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}
