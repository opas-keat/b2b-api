package handlers

import (
	"go.uber.org/dig"
	"user/handlers/callback"
	"user/handlers/user"
)

type HandlerParams struct {
	dig.In
	//Provider *provider.Handlers
	Callback *callback.Handlers
	User     *user.Handlers
}
