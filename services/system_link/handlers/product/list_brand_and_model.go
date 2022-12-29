package product

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ListBrandAndModel(c *fiber.Ctx) error {
	productType := c.Params("product_type")
	resp, err := h.productService.ListBrandAndModel(c.UserContext(), productType)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
