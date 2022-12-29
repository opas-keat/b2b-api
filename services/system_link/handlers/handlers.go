package handlers

import (
	"go.uber.org/dig"
	"systemlink/handlers/dealer"
)

type HandlerParams struct {
	dig.In
	Dealer *dealer.Handlers
}
