package handlers

import (
	"go.uber.org/dig"
	"member/handlers/user"
)

type HandlerParams struct {
	dig.In
	User *user.Handlers
}
