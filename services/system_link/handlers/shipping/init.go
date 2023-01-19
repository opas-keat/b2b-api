package shipping

import (
	"systemlink/configs"
	shippingService "systemlink/service/shipping"
)

type Handlers struct {
	config          *configs.AppConfig
	shippingService shippingService.Service
}

func New(config *configs.AppConfig, shippingService shippingService.Service) *Handlers {
	return &Handlers{
		config,
		shippingService,
	}
}
