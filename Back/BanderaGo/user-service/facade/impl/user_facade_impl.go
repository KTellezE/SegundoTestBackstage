package impl

import (
	"application/dtos/input"
	"application/dtos/output"
	"application/services"
)

type UserFacadeImpl struct {
	UserService services.UserService
}

func NewUserFacade(service services.UserService) *UserFacadeImpl {
	return &UserFacadeImpl{UserService: service}
}

func (f *UserFacadeImpl) CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error) {
	return f.UserService.CreateUser(userIn)
}

func (f *UserFacadeImpl) GetUserByID(id uint) (output.GetUserOut, error) {
	return f.UserService.GetUserByID(id)
}

func (f *UserFacadeImpl) GetAllUsers() ([]output.GetUsersOut, error) {
	return f.UserService.GetAllUsers()
}

func (f *UserFacadeImpl) UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error) {
	return f.UserService.UpdateUser(id, userIn)
}

func (f *UserFacadeImpl) DeleteUser(id uint) (output.DeleteUserOut, error) {
	return f.UserService.DeleteUser(id)
}
