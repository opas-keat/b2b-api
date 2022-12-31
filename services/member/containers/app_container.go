package containers

import (
	dealerconnector "connector/dealer"
	"fmt"
	"go.uber.org/dig"
	"golang.org/x/exp/slog"
	"member/configs"
	dealerHandler "member/handlers/dealer"
	userHandler "member/handlers/user"
	"member/infrastructure/database"
	"member/infrastructure/server"
	"member/repo/dealer"
	"member/repo/user"
	dealerService "member/service/dealer"
	userService "member/service/user"
	"os"
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
		//connector
		dealerconnector.New,
		// infra
		server.NewFiberServ,
		database.New,
		// repo
		user.NewGormRepo,
		dealer.NewGormRepo,
		// services
		dealerService.New,
		userService.New,
		// handlers
		dealerHandler.New,
		userHandler.New,
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
