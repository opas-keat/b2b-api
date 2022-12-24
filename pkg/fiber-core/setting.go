package fibercore

import (
	"models"

	"github.com/gofiber/fiber/v2"
)

func SetServiceName(serviceName string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals(models.ServiceNameContextKey, serviceName)
		return c.Next()
	}
}
