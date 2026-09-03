package controllers

import (
	"money_management/base"
	"money_management/dtos"
	"money_management/models"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

type TransactionController interface {
	All(c fiber.Ctx) error
	Detail(c fiber.Ctx) error
	CreateTransaction(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type transactionController struct {
	transactionService services.TransactionService
}

func (t transactionController) Delete(c fiber.Ctx) error {
	id := fiber.Params[string](c, "id")
	user := fiber.Locals[models.CurrentUser](c, "user")

	err := t.transactionService.DeleteTransaction(user, id)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[*bool]{
		Message: "Transaction deleted",
		Data:    nil,
	})
}

func (t transactionController) All(c fiber.Ctx) error {
	user := fiber.Locals[models.CurrentUser](c, "user")
	res, err := t.transactionService.GetAll(user)
	if err != nil {
		return err
	}
	return c.JSON(base.Response[map[string][]models.Transaction]{
		Message: "Transactions retrieved",
		Data:    res,
	})
}

func (t transactionController) CreateTransaction(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.TransactionCreateDto](c.Body())
	if err != nil {
		return err
	}
	err = utils.ValidateStruct(body)
	if err != nil {
		return err
	}
	user := fiber.Locals[models.CurrentUser](c, "user")
	var transaction models.Transaction
	if *body.IsIncome == true {
		result, err := t.transactionService.AddIncome(user, body.Amount, body.Description)
		if err != nil {
			return err
		}
		transaction = result
	} else {
		result, err := t.transactionService.AddExpense(user, body.Amount, body.Description)
		if err != nil {
			return err
		}
		transaction = result
	}
	return c.Status(fiber.StatusCreated).JSON(base.Response[models.Transaction]{
		Message: "Transaction Created",
		Data:    transaction,
	})
}

func (t transactionController) Detail(c fiber.Ctx) error {
	id := fiber.Params[string](c, "id")
	errId := utils.ValidateVariable(id, "id", "required,uuid")
	if errId != nil {
		return errId
	}
	user := fiber.Locals[models.CurrentUser](c, "user")
	result, err := t.transactionService.GetDetail(user, id)
	if err != nil {
		return err
	}

	return c.JSON(base.Response[models.Transaction]{
		Message: "Transaction retrieved",
		Data:    result,
	})
}

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return &transactionController{transactionService: transactionService}
}
