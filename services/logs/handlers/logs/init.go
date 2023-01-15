package logs

import (
	"logs/configs"
	logsService "logs/service/logs"
)

type Handlers struct {
	config      *configs.AppConfig
	logsService logsService.Service
}

func New(config *configs.AppConfig, logsService logsService.Service) *Handlers {
	return &Handlers{
		config,
		logsService,
	}
}
