package configs

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/dig"
	"models"
)

type AppConfig struct {
	Mode models.Mode `env:"MODE"`
	Port int         `env:"PORT"`

	DBHost     string `env:"DB_HOST"`
	DBUser     string `env:"DB_USER"`
	DBName     string `env:"DB_NAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     string `env:"DB_PORT"`

	//Redis rediscache.RedisConfig `envPrefix:"REDIS_"`
	//
	//WalletURL   models.WalletURL   `env:"WALLET_URL"`
	//ExternalURL models.ExternalURL `env:"EXTERNAL_URL"`
	//ReportURL   models.ReportURL   `env:"REPORT_URL"`
	//ConfigURL   models.ConfigURL   `env:"CONFIG_URL"`

	//AccessTokenPrivate  string `env:"ACCESS_TOKEN_PRI"`
	//RefreshTokenPrivate string `env:"REFRESH_TOKEN_PRI"`
	//RefreshTokenPublic  string `env:"REFRESH_TOKEN_PUB"`

	//JaegerEndpoint models.JaegerEndpoint `env:"JAEGER_ENDPOINT"`
}

type ExposeAppConfig struct {
	dig.Out
	AppConfig *AppConfig
	//RedisConfig    *rediscache.RedisConfig
	//Mode           models.Mode
	//WalletURL      models.WalletURL
	//ExternalURL    models.ExternalURL
	//ReportURL      models.ReportURL
	//ConfigURL      models.ConfigURL
	//ServiceName    models.ServiceName
	//JaegerEndpoint models.JaegerEndpoint
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
		//RedisConfig:    &appCfg.Redis,
		//Mode:           appCfg.Mode,
		//WalletURL:      appCfg.WalletURL,
		//ExternalURL:    appCfg.ExternalURL,
		//ReportURL:      appCfg.ReportURL,
		//ConfigURL:      appCfg.ConfigURL,
		//JaegerEndpoint: appCfg.JaegerEndpoint,
		//ServiceName:    constant.ServiceName,
	}, nil
}
