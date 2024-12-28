package impl

import (
	"application/models"
	"application/persistence/repositories"
)

type UserRepositoryImpl struct {
	db repositories.GormDB
}

func NewUserRepository(db repositories.GormDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) UpdateUser(id uint, updatedUser *models.User) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	user.Name = updatedUser.Name
	user.LastName = updatedUser.LastName

	return r.db.Save(&user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
