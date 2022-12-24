package fibercore

import (
	"github.com/gofiber/fiber/v2"
	"models"
)

func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}

func JSONSuccess[T any](c *fiber.Ctx, resp T) error {
	return c.Status(fiber.StatusOK).JSON(new(models.Response[T]).Success(c.Context(), &resp))
}
