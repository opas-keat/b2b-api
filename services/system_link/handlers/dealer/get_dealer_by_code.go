package dealer

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetDealerByCode(c *fiber.Ctx) error {
	code := c.Params("dealer_code")
	println(code)

	resp, err := h.dealerService.GetDealerByCode(c.UserContext(), code)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
