package shipping

import (
	"fibercore"
	"models"
	"models/shipping"
	"models/status_code"
	"shareerrors"
	"validator"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ListShippingByRegion(c *fiber.Ctx) error {
	req := new(models.ListRequest[shipping.ListShippingRequest])
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	//userDetail, err := fibercore.GetUserFromHeader(c)
	//if err != nil {
	//	return err
	//}
	resp, err := h.shippingService.ListShippingByRegion(c.UserContext(), pointer.Get[shipping.ListShippingRequest](req.Criteria), req.Pagination)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
