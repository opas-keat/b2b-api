package containers

import (
	"fmt"
	"os"
	"systemlink/configs"
	dealerHandler "systemlink/handlers/dealer"
	presaleorderHandler "systemlink/handlers/presaleorder"
	productHandler "systemlink/handlers/product"
	shippingsHandler "systemlink/handlers/shipping"
	"systemlink/infrastructure/database"
	"systemlink/infrastructure/server"
	"systemlink/repo/dealer"
	"systemlink/repo/presaleorder"
	"systemlink/repo/product"
	"systemlink/repo/shipping"

	dealerService "systemlink/service/dealer"
	presaleorderService "systemlink/service/presaleorder"
	productService "systemlink/service/product"
	shippingService "systemlink/service/shipping"

	"go.uber.org/dig"
	"golang.org/x/exp/slog"
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
		product.NewGormRepo,
		shipping.NewGormRepo,
		presaleorder.NewGormRepo,
		// services
		dealerService.New,
		productService.New,
		shippingService.New,
		presaleorderService.New,
		// handlers
		dealerHandler.New,
		productHandler.New,
		shippingsHandler.New,
		presaleorderHandler.New,
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
