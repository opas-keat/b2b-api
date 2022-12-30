package user

import (
	"member/configs"
	userService "member/service/user"
)

type Handlers struct {
	config      *configs.AppConfig
	userService userService.Service
}

func New(config *configs.AppConfig, userService userService.Service) *Handlers {
	return &Handlers{
		config,
		userService,
	}
}
