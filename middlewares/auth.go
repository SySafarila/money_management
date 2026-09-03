package middlewares

import (
	"money_management/models"
	"money_management/utils"

	"github.com/gofiber/fiber/v3"
)

func AuthCheck(c fiber.Ctx) error {
	var apiKey string
	if c.Get("X-Api-Key") != "" {
		apiKey = c.Get("x-api-key")
	}
	currentUser, err := utils.ValidateAndParseJwt(apiKey)
	if err != nil {
		return err
	}
	fiber.Locals[models.CurrentUser](c, "user", currentUser)
	return c.Next()
}
