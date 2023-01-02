package product

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
	"models/product"
	"models/status_code"
	"shareerrors"
	"validator"
)

func (h *Handlers) Create(c *fiber.Ctx) error {
	req := new(product.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.productService.Create(c.UserContext(), *req)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
