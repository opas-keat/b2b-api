package dealer

import (
	"member/configs"
	dealerService "member/service/dealer"
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
