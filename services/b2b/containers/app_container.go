package containers

import (
	"b2b/configs"
	"b2b/handlers/callback"
	userHandler "b2b/handlers/user"
	"b2b/infrastructure/database"
	"b2b/infrastructure/server"
	"b2b/repo/user"
	userService "b2b/service/user"
	"fmt"
	"go.uber.org/dig"
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
		// handlers
		callback.New,
		userHandler.New,
		//providerHandler.New,
		//transactionService.New,
		// services
		userService.New,
		// repo
		//transaction.NewGormRepo,
		user.NewGormRepo,
		//awc.New,
		//// pkg
		//tracer.New,
		//rediscache.New,
		//// connectors
		//reportconnector.New,
		//memberConnector.New,
		//walletConnector.New,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}

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
