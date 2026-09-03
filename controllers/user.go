package controllers

import (
	"money_management/base"
	"money_management/models"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

type UserController interface {
	DetailUser(c fiber.Ctx) error
	GetAllUsers(c fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

func (u userController) GetAllUsers(c fiber.Ctx) error {
	queries, errQuery := utils.ParseQueries[models.UserQueryParams](c.Queries())
	if errQuery != nil {
		return errQuery
	}
	users, err := u.userService.FindAll(queries)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[[]models.User]{
		Message: "Users retrieved",
		Data:    users,
	})
}

func (u userController) DetailUser(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.userService.Find(id)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[models.User]{
		Message: "User found",
		Data:    user,
	})
}

func NewUserController(userService services.UserService) UserController {
	return userController{userService: userService}
}
