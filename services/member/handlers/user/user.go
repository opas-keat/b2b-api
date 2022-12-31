package user

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
	"models/status_code"
	"models/user"
	"shareerrors"
	"validator"
)

func (h *Handlers) Login(c *fiber.Ctx) error {
	req := new(user.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	resp, err := h.userService.Login(c.UserContext(), req.Email, req.Password)
	if err != nil {
		return err
	}
	//c.Cookie(&fiber.Cookie{
	//	Name:  "access_token",
	//	Value: resp.AccessToken,
	//})
	return fibercore.JSONSuccess(c, resp)
}
