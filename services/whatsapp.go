package services

import (
	"encoding/json"
	"fmt"
	"log"
	"money_management/dtos"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"github.com/spf13/viper"
)

var whatsappApiVersion = "v24.0"
var whatsappBusinessPhoneNumberId = "876979305497650"

type WhatsAppService interface {
	ReceiveMessage(message dtos.WhatsAppReceiveMessagePayload) []string
	SendMessage(payload dtos.WhatsAppSendMessageWithUrlPayload) error
	verify(from string, merchantCode string, verifyCode string) error
}

type whatsAppService struct {
	//userRepository repositories.UserRepository
}

func (w whatsAppService) verify(from string, merchantCode string, verifyCode string) error {
	reply := fmt.Sprintf("Verifying user for merchant code: %s with code: %s", merchantCode, verifyCode)
	callbackUrl := fmt.Sprintf("https://fsm.web.id/vouchers/%s", "voucherCode")
	callBackTitle := "Continue"
	req := new(dtos.WhatsAppSendMessageWithUrlPayload)
	req.MessagingProduct = "whatsapp"
	req.RecipientType = "individual"
	req.To = from
	req.Type = "interactive"
	req.Interactive.Type = "cta_url"
	req.Interactive.Body.Text = reply
	req.Interactive.Action.Name = "cta_url"
	req.Interactive.Action.Parameters.DisplayText = callBackTitle
	req.Interactive.Action.Parameters.Url = callbackUrl
	req.Interactive.Footer.Text = "Sent automatically by system."
	errReply := w.SendMessage(*req)
	if errReply != nil {
		return errReply
	}
	return nil
}

func (w whatsAppService) SendMessage(payload dtos.WhatsAppSendMessageWithUrlPayload) error {
	reqByte, errJson := json.Marshal(payload)
	if errJson != nil {
		return fiber.ErrInternalServerError
	}

	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", whatsappApiVersion, whatsappBusinessPhoneNumberId)
	token := viper.GetString("WHATSAPP_ACCESS_TOKEN")
	authorizationValue := fmt.Sprintf("Bearer %s", token)
	cc := client.New()
	resp, errHttpClient := cc.Post(url, client.Config{
		Header: map[string]string{
			"Authorization": authorizationValue,
			"Content-Type":  "application/json",
			"Accept":        "application/json",
		},
		Body: string(reqByte),
	})
	if errHttpClient != nil {
		return fiber.ErrInternalServerError
	}
	defer resp.Close()
	if resp.StatusCode() > 300 {
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
				split := strings.Split(messageBody, " ")
				if len(split) >= 3 && split[0] == "/verify" {
					merchantCode := split[1]
					verifyCode := split[2]
					errVerify := w.verify(messages[i2].From, merchantCode, verifyCode)
					if errVerify != nil {
						log.Println(errVerify)
					}
				}
				texts = append(texts, messageBody)
			}
		}
	}
	return texts
}

func NewWhatsAppService() WhatsAppService {
	return whatsAppService{}
}
