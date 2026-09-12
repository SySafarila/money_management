package services

import (
	"money_management/dtos"
	"money_management/models"
	"money_management/repositories"
)

type CategoryService interface {
	GetAll(user models.CurrentUser) ([]models.Category, error)
	GetDetail(user models.CurrentUser, id string) (models.Category, error)
	NewCategory(user models.CurrentUser, data dtos.CategoryCreateDto) (models.Category, error)
	UpdateCategory(user models.CurrentUser, id string, data dtos.CategoryCreateDto) (models.Category, error)
	DeleteCategory(user models.CurrentUser, id string) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func (t categoryService) UpdateCategory(user models.CurrentUser, id string, data dtos.CategoryCreateDto) (models.Category, error) {
	category, errDetailCategory := t.GetDetail(user, id)
	if errDetailCategory != nil {
		return models.Category{}, errDetailCategory
	}
	category.Name = data.Name
	result, err := t.categoryRepository.Update(user, category.Id, category)
	return result, err
}

func (t categoryService) GetDetail(user models.CurrentUser, id string) (models.Category, error) {
	return t.categoryRepository.Find(user, id)
}

func (t categoryService) GetAll(user models.CurrentUser) ([]models.Category, error) {
	return t.categoryRepository.FindAll(user)
}

func (t categoryService) DeleteCategory(user models.CurrentUser, id string) error {
	category, err := t.GetDetail(user, id)
	if err != nil {
		return err
	}
	return t.categoryRepository.Delete(user, category.Id)
}

func (t categoryService) NewCategory(user models.CurrentUser, data dtos.CategoryCreateDto) (models.Category, error) {
	category := models.Category{
		UserId: user.UserId,
		Name:   data.Name,
	}
	return t.categoryRepository.Create(category)
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: repo}
}
