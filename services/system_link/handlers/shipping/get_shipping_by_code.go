package shipping

import (
	"fibercore"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetShippingByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	println(code)

	resp, err := h.shippingService.GetShippingByCode(c.UserContext(), code)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
