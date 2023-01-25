package presaleorder

import (
	"systemlink/configs"
	presaleorderService "systemlink/service/presaleorder"
)

type Handlers struct {
	config              *configs.AppConfig
	presaleorderService presaleorderService.Service
}

func New(config *configs.AppConfig, presaleorderService presaleorderService.Service) *Handlers {
	return &Handlers{
		config,
		presaleorderService,
	}
}
