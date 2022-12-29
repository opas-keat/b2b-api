package dealer

import (
	"systemlink/configs"
	dealerService "systemlink/service/dealer"
)

type Handlers struct {
	config        *configs.AppConfig
	dealerService dealerService.Service
}

func New(config *configs.AppConfig, dealerService dealerService.Service) *Handlers {
	return &Handlers{
		config,
		dealerService,
	}
}
