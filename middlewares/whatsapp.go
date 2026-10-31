package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

func WhatsAppMiddleware(c fiber.Ctx) error {
	appSecret := viper.GetString("whatsapp_secret")
	if appSecret == "" {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"WhatsApp app secret is not configured",
		)
	}

	signature := c.Get("X-Hub-Signature-256")

	if !strings.HasPrefix(signature, "sha256=") {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Invalid webhook signature",
		)
	}

	signatureHash, err := hex.DecodeString(
		strings.TrimPrefix(signature, "sha256="),
	)

	if err != nil || len(signatureHash) != sha256.Size {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Invalid webhook signature",
		)
	}

	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write(c.Body())

	expectedHash := mac.Sum(nil)

	if !hmac.Equal(signatureHash, expectedHash) {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Invalid webhook signature",
		)
	}

	return c.Next()
}
