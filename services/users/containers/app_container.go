package containers

import (
	"users/configs"

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
		// // infra
		// server.NewFiberServ,
		// database.New,
		// // handlers
		// callback.New,
		// providerHandler.New,
		// transactionService.New,
		// // repo
		// transaction.NewGormRepo,
		// user.NewGormRepo,
		// awc.New,
		// // pkg
		// tracer.New,
		// rediscache.New,
		// // connectors
		// reportconnector.New,
		// memberConnector.New,
		// walletConnector.New,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}

	return nil
}

func (c *AppContainer) Start() error {
	// if err := c.container.Invoke(func(s server.Server) {
	// 	s.Start()
	// }); err != nil {
	// 	fmt.Printf("%s", err)
	// 	return err
	// }

	return nil
}