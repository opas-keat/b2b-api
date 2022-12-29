package containers

import (
	"fmt"
	"go.uber.org/dig"
	"golang.org/x/exp/slog"
	"os"
	"systemlink/configs"
	dealerHandler "systemlink/handlers/dealer"
	"systemlink/infrastructure/database"
	"systemlink/infrastructure/server"
	"systemlink/repo/dealer"
	dealerService "systemlink/service/dealer"
)

type AppContainer struct {
	container *dig.Container
}

func NewAppContainer() (*AppContainer, error) {
	container := &AppContainer{
		container: dig.New(),
	}

	if err := container.configure(); err != nil {
		return nil, err
	}

	return container, nil
}

func (c *AppContainer) configure() error {
	servicesConstructors := []interface{}{
		configs.NewAppConfig,
		// infra
		server.NewFiberServ,
		database.New,
		// repo
		dealer.NewGormRepo,
		// services
		dealerService.New,
		// handlers
		dealerHandler.New,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr)))

	return nil
}

func (c *AppContainer) Start() error {
	if err := c.container.Invoke(func(s server.Server, config *configs.AppConfig) {
		s.Start()
	}); err != nil {
		fmt.Printf("invoke %s", err)
		return err
	}

	return nil
}
