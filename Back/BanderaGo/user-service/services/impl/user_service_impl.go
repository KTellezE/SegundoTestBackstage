package impl

import (
	"application/dtos/input"
	"application/dtos/output"
	"application/models"
	"application/persistence/repositories"
)

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(userIn input.CreateUserIn) (output.CreateUserOut, error) {
	user := models.User{
		Name:     userIn.Name,
		LastName: userIn.LastName,
	}
	if err := s.repo.CreateUser(&user); err != nil {
		return output.CreateUserOut{}, err
	}
	userOut := output.CreateUserOut{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}
	return userOut, nil
}

func (s *UserServiceImpl) GetUserByID(id uint) (output.GetUserOut, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return output.GetUserOut{}, err
	}
	userOut := output.GetUserOut{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
	}
	return userOut, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]output.GetUsersOut, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var usersOut []output.GetUsersOut
	for _, user := range users {
		userOut := output.GetUsersOut{
			ID:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
		}
		usersOut = append(usersOut, userOut)
	}
	return usersOut, nil
}

func (s *UserServiceImpl) UpdateUser(id uint, userIn input.UpdateUserIn) (output.UpdateUserOut, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return output.UpdateUserOut{}, err
	}

	user.Name = userIn.Name
	user.LastName = userIn.LastName

	if err := s.repo.UpdateUser(id, user); err != nil {
		return output.UpdateUserOut{}, err
	}

	userOut := output.UpdateUserOut{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		UpdatedAt: user.UpdatedAt,
	}
	return userOut, nil
}

func (s *UserServiceImpl) DeleteUser(id uint) (output.DeleteUserOut, error) {
	if err := s.repo.DeleteUser(id); err != nil {
		return output.DeleteUserOut{Success: false}, err
	}
	return output.DeleteUserOut{Success: true}, nil
}
