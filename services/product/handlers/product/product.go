package product

import (
	"fibercore"
	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"models"
	"models/product"
	"models/status_code"
	"shareerrors"
	"validator"
)

func (h *Handlers) Create(c *fiber.Ctx) error {
	userDetail, err := fibercore.GetUserFromHeader(c)
	if err != nil {
		return err
	}
	req := new(product.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.productService.Create(c.UserContext(), *userDetail, *req)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}

func (h *Handlers) ListBrandAndModel(c *fiber.Ctx) error {
	req := new(models.ListRequest[product.ListBrandAndModelRequest])
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.productService.ListBrandAndModel(c.UserContext(), pointer.Get[product.ListBrandAndModelRequest](req.Criteria), req.Pagination)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
