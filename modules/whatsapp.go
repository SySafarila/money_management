package modules

import (
	"money_management/controllers"
	"money_management/services"
)

type WhatsAppModule struct {
	Service    services.WhatsAppService
	Controller controllers.WhatsAppController
}

func NewWhatsAppModule() WhatsAppModule {
	whatsAppService := services.NewWhatsAppService()
	whatsAppController := controllers.NewWhatsAppController(whatsAppService)
	return WhatsAppModule{
		Service:    whatsAppService,
		Controller: whatsAppController,
	}
}
