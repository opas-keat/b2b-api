package server

import (
	"fibercore"
	"fmt"
	"log"
	"logs/configs"
	"logs/constant"
	"logs/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FiberServ struct {
	config  *configs.AppConfig
	app     *fiber.App
	handler handlers.HandlerParams
	db      *gorm.DB
}

func NewFiberServ(c *configs.AppConfig, h handlers.HandlerParams, db *gorm.DB) Server {
	app := fibercore.NewApp()
	fiberServ := &FiberServ{
		config:  c,
		app:     app,
		handler: h,
		db:      db,
	}
	fiberServ.configHandler()

	return fiberServ
}

func dbTransactionMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return db.Transaction(func(tx *gorm.DB) error {
			c.Locals(constant.TxContextKey, tx)
			return c.Next()
		})
	}
}

func (f *FiberServ) logsHandler(router fiber.Router) {
}

func (f *FiberServ) configHandler() {
	f.app.Use(fibercore.SetServiceName(constant.ServiceName))
	v1 := f.app.Group("/api/v1")

	logs := v1.Group("/logs")
	f.logsHandler(logs)

}

func (f *FiberServ) Start() {
	go func() {
		if err := f.app.Listen(fmt.Sprintf("127.0.0.1:%v", f.config.Port)); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = f.app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
