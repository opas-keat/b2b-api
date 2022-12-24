package callback

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"strings"
	"user/entities"
)

func (h *Handlers) Corridor(c *fiber.Ctx) error {

	var (
		key             = c.FormValue("key")
		message         = c.FormValue("message")
		messageReplacer = strings.NewReplacer("\n", "", "\r", "", "\t", "", "\\", "")
	)
	if key == "" {
		return h.response(c, new(entities.Response).Create(entities.Invalidparameters, errors.New("key is required")))
	}
	//if key != h.config.AwcSecretKey {
	//	return h.response(c, new(entities.Response).Create(entities.InvalidCertificate, nil))
	//}

	cleanUpMessage := messageReplacer.Replace(message)
	action := gjson.Get(cleanUpMessage, "action").String()

	//if next, ok := h.router[action]; ok {
	//	return next(c, []byte(cleanUpMessage))
	//}
	return h.response(c, new(entities.Response).Create(entities.Invalidparameters, errors.New(fmt.Sprintf("unknow action %s", action))))

}
