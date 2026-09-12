package modules

import (
	"money_management/controllers"
	"money_management/repositories"
	"money_management/services"
)

type CategoryModule struct {
	Repository repositories.CategoryRepository
	Service    services.CategoryService
	Controller controllers.CategoryController
}

func NewCategoryModule() CategoryModule {
	repository := repositories.NewCategoryRepository()
	service := services.NewCategoryService(repository)
	controller := controllers.NewCategoryController(service)
	return CategoryModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}
