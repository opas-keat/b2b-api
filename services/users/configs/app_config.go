package configs

import (
	"models"
	"users/constant"

	"github.com/caarlos0/env/v6"
	"go.uber.org/dig"
)

type AppConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`

	Mode models.Mode `env:"MODE"`
	Port int         `env:"PORT"`
}

type ExposeAppConfig struct {
	dig.Out
	AppConfig *AppConfig

	Mode        models.Mode
	ServiceName models.ServiceName
}

func NewAppConfig() (ExposeAppConfig, error) {
	cfg := new(AppConfig)
	opts := []env.Options{
		{RequiredIfNoDef: true},
	}

	if err := env.Parse(cfg, opts...); err != nil {
		return ExposeAppConfig{}, err
	}

	return ExposeAppConfig{
		AppConfig: cfg,

		Mode:        cfg.Mode,
		ServiceName: constant.ServiceName,
	}, nil
}
