package handlers

import (
	"logs/handlers/logs"

	"go.uber.org/dig"
)

type HandlerParams struct {
	dig.In
	Logs *logs.Handlers
}
