package repositories

import (
	"money_management/database"
	"money_management/dtos"
	"money_management/models"
)

type UserRepository interface {
	Find(id string) (models.User, error)
	FindAll(q models.UserQueryParams) ([]models.User, error)
	FindByUsername(username string) (models.User, error)
	RegisterUser(data models.User) (models.User, error)
}

type userRepository struct{}

func (u userRepository) RegisterUser(data models.User) (models.User, error) {
	user := dtos.RegisterRequest{
		Username: data.Username,
		Password: data.Password,
		FullName: data.FullName,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return data, nil
}

func (u userRepository) FindByUsername(username string) (models.User, error) {
	user := models.User{}
	err := database.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u userRepository) FindAll(q models.UserQueryParams) ([]models.User, error) {
	var users []models.User
	db := database.DB
	if q.Id != "" {
		db = db.Where("id = ?", q.Id)
	}
	if q.Username != "" {
		db = db.Where("username iLike ?", "%"+q.Username+"%")
	}
	if q.Verified == "true" {
		db = db.Where("verified_at IS NOT NULL")
	} else if q.Verified == "false" {
		db = db.Where("verified_at IS NULL")
	}
	err := db.Find(&users).Error
	return users, err
}

func (u userRepository) Find(id string) (models.User, error) {
	user := models.User{}
	err := database.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func NewUserRepository() UserRepository {
	return userRepository{}
}
