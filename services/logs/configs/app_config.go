package configs

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/dig"
)

type AppConfig struct {
	Mode string `env:"MODE"`
	Port int    `env:"PORT"`

	DBHost     string `env:"DB_HOST"`
	DBUser     string `env:"DB_USER"`
	DBName     string `env:"DB_NAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     string `env:"DB_PORT"`
}

type ExposeAppConfig struct {
	dig.Out
	AppConfig *AppConfig
}

func NewAppConfig() (ExposeAppConfig, error) {
	appCfg := new(AppConfig)
	opts := []env.Options{
		{RequiredIfNoDef: true},
	}

	if err := env.Parse(appCfg, opts...); err != nil {
		return ExposeAppConfig{}, err
	}

	return ExposeAppConfig{
		AppConfig: appCfg,
	}, nil
}
