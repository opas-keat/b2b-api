package product

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetProductByCode(c *fiber.Ctx) error {
	code := c.Params("product_code")
	resp, err := h.productService.GetProductByCode(c.UserContext(), code)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
