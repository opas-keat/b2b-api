package dealer

import (
	"fibercore"
	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"models"
	"models/dealer"
	"models/status_code"
	"shareerrors"
	"validator"
)

func (h *Handlers) ListDealerByCodeInternal(c *fiber.Ctx) error {
	req := new(models.ListRequest[dealer.ListDealerRequest])
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
	resp, err := h.dealerService.ListDealerByCodeInternal(c.UserContext(), pointer.Get[dealer.ListDealerRequest](req.Criteria), req.Pagination)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}
