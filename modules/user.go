package modules

import (
	"money_management/controllers"
	"money_management/repositories"
	"money_management/services"
)

type UserModule struct {
	Respository repositories.UserRepository
	Service     services.UserService
	Controller  controllers.UserController
}

func NewUserModule() UserModule {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	return UserModule{
		Respository: userRepo,
		Service:     userService,
		Controller:  userController,
	}
}
