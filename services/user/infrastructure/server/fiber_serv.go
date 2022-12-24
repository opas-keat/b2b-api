package server

import (
	"fibercore"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user/configs"
	"user/constant"
	"user/handlers"
)

type FiberServ struct {
	config  *configs.AppConfig
	app     *fiber.App
	handler handlers.HandlerParams
	//db      *gorm.DB
}

//func NewFiberServ(c *configs.AppConfig, h handlers.HandlerParams, db *gorm.DB) Server {
//	app := fibercore.NewApp()
//	app.Use(fibercore.OpentracingMiddleware(nil))
//	fiberServ := &FiberServ{
//		config: c,
//		app:    app,
//		//handler: h,
//		//db:      db,
//	}
//	fiberServ.configHandler()
//
//	return fiberServ
//}

func NewFiberServ(c *configs.AppConfig, h handlers.HandlerParams) Server {
	app := fibercore.NewApp()
	//app.Use(fibercore.OpentracingMiddleware(nil))
	fiberServ := &FiberServ{
		config:  c,
		app:     app,
		handler: h,
		//db:      db,
	}
	fiberServ.configHandler()

	return fiberServ
}

func (f *FiberServ) configHandler() {
	f.app.Use(fibercore.SetServiceName(constant.ServiceName))
	v1 := f.app.Group("/api/v1/users")

	callback := v1.Group("/callback")
	callback.Post("/", f.handler.Callback.Corridor)
	//
	//provider := v1.Group("/provider")
	//provider.Get("/launch_game", f.handler.Provider.LaunchGame)

}

func (f FiberServ) Start() {
	go func() {
		if err := f.app.Listen(fmt.Sprintf(":%v", f.config.Port)); err != nil {
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
