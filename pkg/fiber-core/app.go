package fibercore

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"os"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{})
	app.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUIDv4,
		ContextKey: fiber.HeaderXRequestID,
	}))

	app.Use(cors.New(cors.ConfigDefault))
	app.Use(setUserLocal())

	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			switch c.Path() {
			case "/health", "/ping", "/metrics", "/api/v1/file/upload", "/api/v1/members/verify-session":
				return true
			default:
				return false
			}
		},
		Format: fmt.Sprintf(
			`{"time":"${time}","request_id":"${locals:X-Request-ID}","member_id":"${locals:X-Member-ID}","status":"${status}","latency":"${latency}","method":"${method}","path":"${path}","body":${body},"res_body":${resBody}}`,
		) + fmt.Sprintf("\n"),
		TimeFormat: "2006-01-02T15:04:05",
		Output:     os.Stdout,
	}))

	app.Use(pprof.New())
	app.Get("metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	app.Get("ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("pong")
	})
	app.Get("health", Health)

	return app
}
