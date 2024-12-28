package services

import (
	"application/dtos/input"
	"application/dtos/output"
)

type UserService interface {
	CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error)
	GetUserByID(id uint) (output.GetUserOut, error)
	GetAllUsers() ([]output.GetUsersOut, error)
	UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error)
	DeleteUser(id uint) (output.DeleteUserOut, error)
}
