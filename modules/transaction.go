package modules

import (
	"money_management/controllers"
	"money_management/repositories"
	"money_management/services"
)

type TransactionModule struct {
	Repository repositories.TransactionRepository
	Service    services.TransactionService
	Controller controllers.TransactionController
}

func NewTransactionModule() TransactionModule {
	repository := repositories.NewTransactionRepository()
	service := services.NewTransactionService(repository)
	controller := controllers.NewTransactionController(service)
	return TransactionModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}
