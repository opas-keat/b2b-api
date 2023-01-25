package presaleorder

import (
	"fibercore"
	"models/presaleorder"
	"models/status_code"
	"shareerrors"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) CreatePreSaleOrder(c *fiber.Ctx) error {
	req := new(presaleorder.PreSaleOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	resp, err := h.presaleorderService.CreatePreSaleOrder(c.UserContext(), pointer.Get(req))
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
