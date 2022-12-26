package user

import (
	"fibercore"
	"github.com/gofiber/fiber/v2"
	"models/status_code"
	"models/user"
	"shareerrors"
	"validator"
)

type Handlers struct {
}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Me(c *fiber.Ctx) error {
	//userDetail, err := fibercore.GetUserFromHeader(c)
	//if err != nil {
	//	return err
	//}

	//me, err := h.memberService.Me(c.UserContext(), *userDetail)
	//if err != nil {
	//	return err
	//}
	var me = "me"
	return fibercore.JSONSuccess(c, me)
}

func (h *Handlers) Logout(c *fiber.Ctx) error {
	var response = "Logout"
	return fibercore.JSONSuccess(c, response)
}

func (h *Handlers) Login(c *fiber.Ctx) error {
	var response = "Login"
	return fibercore.JSONSuccess(c, response)
}

func (h *Handlers) VerifyEmail(c *fiber.Ctx) error {
	var response = "VerifyEmail"
	return fibercore.JSONSuccess(c, response)
}

func (h *Handlers) Register(c *fiber.Ctx) error {
	//ipAddress, err := fibercore.GetIPAddress(c)
	//if err != nil {
	//	return err
	//}
	//req.IPAddress = *ipAddress
	req := new(user.RegisterUserRequest)
	if err := c.BodyParser(req); err != nil {
		return shareerrors.NewError(status_code.BadRequest, err.Error())
	}
	//println("pass: " + req.Password)
	println("email: " + req.Email)
	//println("code: " + req.DealerCode)
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}
	var response = "Register"
	return fibercore.JSONSuccess(c, response)
}
