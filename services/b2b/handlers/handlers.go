package handlers

import (
	"b2b/handlers/user"
	"go.uber.org/dig"
)

type HandlerParams struct {
	dig.In
	//Provider *provider.Handlers
	//Callback *callback.Handlers
	User *user.Handlers
}
