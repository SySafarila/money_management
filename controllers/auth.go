package controllers

import (
	"money_management/base"
	"money_management/dtos"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

type AuthController interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
}

type authController struct {
	authService services.AuthService
}

func (a authController) Login(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.LoginRequest](c.Body())
	if err != nil {
		return err
	}
	err = utils.ValidateStruct(body)
	if err != nil {
		return err
	}
	token, eat, err := a.authService.Login(body)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[dtos.LoginResponse]{
		Message: "Login success",
		Data: dtos.LoginResponse{
			Token:      token,
			ValidUntil: eat,
		},
	})
}

func (a authController) Register(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.RegisterRequest](c.Body())
	if err != nil {
		return err
	}
	username, err := a.authService.Register(body)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[dtos.RegisterResponse]{
		Message: "Register success",
		Data: dtos.RegisterResponse{
			Username: username,
		},
	})
}

func (a authController) Logout(c fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthController(authService services.AuthService) AuthController {
	return authController{authService: authService}
}
