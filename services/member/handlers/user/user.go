package user

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
	"models/status_code"
	"models/user"
	"shareerrors"
	"validator"
)

func (h *Handlers) Me(c *fiber.Ctx) error {
	userDetail, err := fibercore.GetUserFromHeader(c)
	if err != nil {
		return err
	}
	me, err := h.userService.Me(c.UserContext(), *userDetail)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, me)
}

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

func (h *Handlers) Register(c *fiber.Ctx) error {
	req := new(user.RegisterUserRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	//println("email: " + req.Email)
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	resp, err := h.userService.Register(c.UserContext(), *req)
	if err != nil {
		return err
	}
	return fibercore.JSONSuccess(c, resp)
}

func (h *Handlers) VerifyEmail(c *fiber.Ctx) error {
	code := c.Params("code")
	resp, err := h.userService.VerifyEmail(c.UserContext(), code)
	if err != nil {
		return err
	}
	//return resp.Message
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	resText := "<!DOCTYPE html><header>" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />" +
		"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />" +
		"</header><body><h2>" + resp.Message + "</h2></body></html>"
	return c.SendString(resText)
	//return fibercore.JSONSuccess(c, resp)
}
