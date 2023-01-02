package product

import (
	"product/configs"
	productService "product/service/product"
)

type Handlers struct {
	config         *configs.AppConfig
	productService productService.Service
}

func New(config *configs.AppConfig, productService productService.Service) *Handlers {
	return &Handlers{
		config,
		productService,
	}
}
