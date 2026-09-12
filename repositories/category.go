package repositories

import (
	"money_management/database"
	"money_management/models"
)

type CategoryRepository interface {
	Find(user models.CurrentUser, id string) (models.Category, error)
	FindAll(user models.CurrentUser) ([]models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(user models.CurrentUser, id string, category models.Category) (models.Category, error)
	Delete(user models.CurrentUser, id string) error
}

type categoryRepository struct{}

func (t categoryRepository) Find(user models.CurrentUser, id string) (models.Category, error) {
	var category models.Category
	err := database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).First(&category).Error
	return category, err
}

func (t categoryRepository) FindAll(user models.CurrentUser) ([]models.Category, error) {
	var categorys []models.Category
	err := database.DB.Where("user_id = ?", user.UserId).Order("name ASC").Find(&categorys).Error
	return categorys, err
}

func (t categoryRepository) Create(category models.Category) (models.Category, error) {
	err := database.DB.Create(&category).Error
	return category, err
}

func (t categoryRepository) Update(user models.CurrentUser, id string, category models.Category) (models.Category, error) {
	err := database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).Updates(&category).Error
	return category, err
}

func (t categoryRepository) Delete(user models.CurrentUser, id string) error {
	return database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).Delete(&models.Category{}).Error
}

func NewCategoryRepository() CategoryRepository {
	return categoryRepository{}
}
