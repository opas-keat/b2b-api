package fibercore

import (
	"github.com/gofiber/fiber/v2"
	"models"
	"models/gateway"
	"models/status_code"
	"shareerrors"
	"strings"
)

func setUserLocal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, header := range models.ArrHeader {
			c.Locals(header, c.Get(header))
		}
		if strings.Contains(c.Path(), "public") {
			c.Locals(models.HeaderIPAddress, c.Get("x-real-ip", c.IP()))
		}
		return c.Next()
	}
}

func GetIPAddress(c *fiber.Ctx) (*string, error) {
	headerIPAddress := c.Locals(models.HeaderIPAddress).(string)
	if headerIPAddress == "" {
		return nil, shareerrors.NewError(status_code.BadRequest, "Header X-IP-Address is require")
	}
	return &headerIPAddress, nil
}

func GetUserFromHeader(c *fiber.Ctx) (*gateway.User, error) {
	var (
		//headerXUserName     = c.Locals(models.HeaderXUsername).(string)
		//headerXMemberId     = c.Locals(models.HeaderXMemberID).(string)
		//headerXRole         = c.Locals(models.HeaderXRole).(string)
		//headerXAgentID      = c.Locals(models.HeaderXAgentID).(string)
		//headerXMasterID     = c.Locals(models.HeaderXMasterID).(string)
		//headerXReferrer     = c.Locals(models.HeaderReferrer).(string)
		headerAuthorization = c.Locals(models.HeaderAuthorization).(string)
	)

	//if headerXUserName == "" {
	//	return nil, shareerrors.NewError(status_code.BadRequest, "Header X-Username is require")
	//}
	//
	//if headerXMemberId == "" {
	//	return nil, shareerrors.NewError(status_code.BadRequest, "Header X-Member-ID is require")
	//}
	//
	//if headerXRole == "" {
	//	return nil, shareerrors.NewError(status_code.BadRequest, "Header X-Role is require")
	//}
	//
	//ipAddress, err := GetIPAddress(c)
	//if err != nil {
	//	return nil, err
	//}
	//
	//role, ok := models.MappingToRole[c.Get(models.HeaderXRole)]
	//if !ok {
	//	return nil, shareerrors.NewError(status_code.BadRequest, "Header X-Role invalid")
	//}
	//
	//agentID := &headerXAgentID
	//if headerXAgentID == "" {
	//	agentID = nil
	//}
	//
	//masterID := &headerXMasterID
	//if headerXMasterID == "" {
	//	masterID = nil
	//}

	var token string
	fields := strings.Fields(headerAuthorization)
	if len(fields) != 0 && fields[0] == "Bearer" {
		token = fields[1]
	} else {
		return nil, shareerrors.NewError(status_code.Forbidden, "Authorization is require")
	}
	if token == "" {
		return nil, shareerrors.NewError(status_code.Forbidden, "Authorization is require")
	}
	//sub, err := utils.ValidateToken(token, constant.PrivateKey)
	//if err != nil {
	//	return nil, shareerrors.NewError(status_code.Forbidden, "token is invalid")
	//}
	//fields = strings.Split(fmt.Sprint(sub), "|")
	//return &gateway.User{
	//	Username:         fields[2],
	//	AgentID:          nil,
	//	MasterID:         nil,
	//	IPAddress:        "",
	//	MemberID:         fields[0],
	//	Role:             "",
	//	ReferrerMemberID: "",
	//	RoleName:         fields[1],
	//}, nil
	return &gateway.User{
		Username:         "",
		AgentID:          nil,
		MasterID:         nil,
		IPAddress:        "",
		MemberID:         "",
		Role:             "",
		ReferrerMemberID: "",
		RoleName:         "",
	}, nil
}
