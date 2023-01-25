package handlers

import (
	"systemlink/handlers/dealer"
	"systemlink/handlers/presaleorder"
	"systemlink/handlers/product"
	"systemlink/handlers/shipping"

	"go.uber.org/dig"
)

type HandlerParams struct {
	dig.In
	Dealer       *dealer.Handlers
	Product      *product.Handlers
	Shipping     *shipping.Handlers
	PreSaleOrder *presaleorder.Handlers
}
