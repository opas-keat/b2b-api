package handlers

import (
	"go.uber.org/dig"
	"systemlink/handlers/dealer"
	"systemlink/handlers/product"
)

type HandlerParams struct {
	dig.In
	Dealer  *dealer.Handlers
	Product *product.Handlers
}
