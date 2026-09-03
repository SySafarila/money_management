package services

import (
	"money_management/models"
	"money_management/repositories"
)

type UserService interface {
	Find(id string) (models.User, error)
	FindAll(queries models.UserQueryParams) ([]models.User, error)
	FindByUsername(username string) (models.User, error)
	RegisterUser(data models.User) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func (u userService) RegisterUser(data models.User) (models.User, error) {
	return u.userRepository.RegisterUser(data)
}

func (u userService) FindByUsername(username string) (models.User, error) {
	return u.userRepository.FindByUsername(username)
}

func (u userService) FindAll(queries models.UserQueryParams) ([]models.User, error) {
	return u.userRepository.FindAll(queries)
}

func (u userService) Find(id string) (models.User, error) {
	return u.userRepository.Find(id)
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return userService{userRepository: userRepository}
}
