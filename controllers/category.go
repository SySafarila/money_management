package controllers

import (
	"money_management/base"
	"money_management/dtos"
	"money_management/models"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

type CategoryController interface {
	All(c fiber.Ctx) error
	Detail(c fiber.Ctx) error
	CreateCategory(c fiber.Ctx) error
	UpdateCategory(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type categoryController struct {
	categoryService services.CategoryService
}

func (t categoryController) UpdateCategory(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.CategoryCreateDto](c.Body())
	if err != nil {
		return err
	}
	err = utils.ValidateStruct(body)
	if err != nil {
		return err
	}
	user := fiber.Locals[models.CurrentUser](c, "user")
	id := fiber.Params[string](c, "id")
	err = utils.ValidateVariable(id, "id", "required,uuid")
	if err != nil {
		return err
	}
	category, err := t.categoryService.UpdateCategory(user, id, body)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(base.Response[models.Category]{
		Message: "Category updated",
		Data:    category,
	})
}

func (t categoryController) Delete(c fiber.Ctx) error {
	id := fiber.Params[string](c, "id")
	err := utils.ValidateVariable(id, "id", "required,uuid")
	if err != nil {
		return err
	}
	user := fiber.Locals[models.CurrentUser](c, "user")

	err = t.categoryService.DeleteCategory(user, id)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[*bool]{
		Message: "Category deleted",
		Data:    nil,
	})
}

func (t categoryController) All(c fiber.Ctx) error {
	user := fiber.Locals[models.CurrentUser](c, "user")
	res, err := t.categoryService.GetAll(user)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[[]models.Category]{
		Message: "Categories retrieved",
		Data:    res,
	})
}

func (t categoryController) CreateCategory(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.CategoryCreateDto](c.Body())
	if err != nil {
		return err
	}
	err = utils.ValidateStruct(body)
	if err != nil {
		return err
	}
	user := fiber.Locals[models.CurrentUser](c, "user")
	category, err := t.categoryService.NewCategory(user, body)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(base.Response[models.Category]{
		Message: "Category Created",
		Data:    category,
	})
}

func (t categoryController) Detail(c fiber.Ctx) error {
	id := fiber.Params[string](c, "id")
	errId := utils.ValidateVariable(id, "id", "required,uuid")
	if errId != nil {
		return errId
	}
	user := fiber.Locals[models.CurrentUser](c, "user")
	result, err := t.categoryService.GetDetail(user, id)
	if err != nil {
		return err
	}

	return c.JSON(base.Response[models.Category]{
		Message: "Category retrieved",
		Data:    result,
	})
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &categoryController{categoryService: categoryService}
}
