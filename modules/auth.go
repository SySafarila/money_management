package modules

import (
	"money_management/controllers"
	"money_management/services"
)

type AuthModule struct {
	Service     services.AuthService
	Controller  controllers.AuthController
}

func NewAuthModule(userService services.UserService) AuthModule {
	service := services.NewAuthService(userService)
	controller := controllers.NewAuthController(service)
	return AuthModule{
		Service:     service,
		Controller:  controller,
	}
}
