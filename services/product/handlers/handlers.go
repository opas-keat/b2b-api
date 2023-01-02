package handlers

import (
	"go.uber.org/dig"
	"product/handlers/product"
)

type HandlerParams struct {
	dig.In
	Product *product.Handlers
}
