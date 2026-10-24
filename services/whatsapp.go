package services

import (
	"encoding/json"
	"fmt"
	"money_management/dtos"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
)

var whatsappApiVersion = "VERSION"
var whatsappBusinessPhoneNumberId = "BUSINESS_PHONE_NUMBER_ID"

type WhatsAppService interface {
	ReceiveMessage(message dtos.WhatsAppReceiveMessagePayload) []string
	SendMessage(payload dtos.WhatsAppSendMessagePayload) error
}

type whatsAppService struct {
	//userRepository repositories.UserRepository
}

func (w whatsAppService) SendMessage(payload dtos.WhatsAppSendMessagePayload) error {
	reqByte, errJson := json.Marshal(payload)
	if errJson != nil {
		return fiber.ErrInternalServerError
	}

	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", whatsappApiVersion, whatsappBusinessPhoneNumberId)
	token := ""
	authorizationValue := fmt.Sprintf("Bearer %s", token)
	cc := client.New()
	_, errHttpClient := cc.Post(url, client.Config{
		Header: map[string]string{
			"Authorization": authorizationValue,
		},
		Body: string(reqByte),
	})
	if errHttpClient != nil {
		return fiber.ErrInternalServerError
	}
	return nil
}

func (w whatsAppService) ReceiveMessage(data dtos.WhatsAppReceiveMessagePayload) []string {
	var texts []string
	for i := range data.Entry {
		changes := data.Entry[i].Changes
		for i2 := range changes {
			messages := changes[i2].Value.Messages
			for i3 := range messages {
				// TODO: ack message
				messageBody := messages[i3].Text.Body
				texts = append(texts, messageBody)
			}
		}
	}
	return texts
}

func NewWhatsAppService() WhatsAppService {
	return whatsAppService{}
}
