package handlers

import (
	"go.uber.org/dig"
	"member/handlers/dealer"
	"member/handlers/user"
)

type HandlerParams struct {
	dig.In
	User   *user.Handlers
	Dealer *dealer.Handlers
}
