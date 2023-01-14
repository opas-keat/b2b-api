package logs

import (
	"logs/configs"
)

type Handlers struct {
	config *configs.AppConfig
}

func New(config *configs.AppConfig) *Handlers {
	return &Handlers{
		config,
	}
}
