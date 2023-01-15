package logs

import (
	"fibercore"
	"models"
	"models/logs"
	"models/status_code"
	"shareerrors"
	"validator"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) Create(c *fiber.Ctx) error {
	req := new(logs.CreateLogsRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.logsService.Create(c.UserContext(), *req)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}

func (h *Handlers) List(c *fiber.Ctx) error {
	req := new(models.ListRequest[logs.ListLogsRequest])
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.logsService.List(c.UserContext(), pointer.Get[logs.ListLogsRequest](req.Criteria), req.Pagination)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
