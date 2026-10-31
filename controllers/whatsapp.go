package controllers

import (
	"money_management/base"
	"money_management/dtos"
	"money_management/services"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

type WhatsAppController interface {
	ReceiveMessage(c fiber.Ctx) error
	Verify(c fiber.Ctx) error
}

type whatsAppController struct {
	whatsAppService services.WhatsAppService
}

func (u whatsAppController) Verify(c fiber.Ctx) error {
	hubChallenge := c.Query("hub.challenge")
	token := c.Query("hub.verify_token")

	if token != "SyahrulSafarila" {
		return fiber.ErrUnauthorized
	}

	return c.Status(fiber.StatusOK).SendString(hubChallenge)
}

func (u whatsAppController) ReceiveMessage(c fiber.Ctx) error {
	body, err := utils.ParseBody[dtos.WhatsAppReceiveMessagePayload](c.Body())
	if err != nil {
		return err
	}

	handle := u.whatsAppService.ReceiveMessage(body)

	c.Status(fiber.StatusOK)
	return c.JSON(base.Response[interface{}]{
		Message: "WhatsApps's message retrieved",
		Data:    handle,
	})
}

func NewWhatsAppController(whatsAppService services.WhatsAppService) WhatsAppController {
	return whatsAppController{whatsAppService: whatsAppService}
}
