package utils

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CustomErrorHandler(ctx fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "internal server error"

	// Retrieve the custom status code if it's a *fiber.Error
	var eFiber *fiber.Error
	if errors.As(err, &eFiber) && eFiber != nil {
		code = eFiber.Code
		message = eFiber.Message
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = fiber.StatusNotFound
		message = "record not found"
	}

	if errors.Is(err, ErrNotImplemented) {
		code = fiber.StatusNotImplemented
		message = "not implemented"
	}

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		code = fiber.StatusUnauthorized
		message = "credentials are not valid"
	}

	var eValidation *ValidationError
	if errors.As(err, &eValidation) && eValidation != nil {
		code = fiber.StatusBadRequest
		return ctx.Status(code).JSON(fiber.Map{
			"message": eValidation.Error(),
			"errors":  eValidation.Errors,
		})
	}

	// Send custom error page
	if err != nil {
		if code >= fiber.StatusInternalServerError {
			log.Println(err)
		}
		return ctx.Status(code).JSON(fiber.Map{
			"message": message,
		})
	}

	// Return from handler
	return nil
}

var (
	ErrNotImplemented   = fiber.ErrNotImplemented
	ErrUserNotFound     = fiber.NewError(fiber.StatusNotFound, "user not found")
	ErrUserAlreadyExist = fiber.NewError(fiber.StatusConflict, "user already exist")
	ErrInvalidToken     = fiber.NewError(fiber.StatusUnauthorized, "invalid token")
)

type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return "validation error"
}
