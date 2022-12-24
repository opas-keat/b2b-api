package handlers

import (
	"go.uber.org/dig"
	"user/handlers/callback"
)

type HandlerParams struct {
	dig.In
	//Provider *provider.Handlers
	Callback *callback.Handlers
}
