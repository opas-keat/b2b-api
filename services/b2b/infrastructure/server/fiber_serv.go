package server

import (
	"b2b/configs"
	"b2b/constant"
	"b2b/handlers"
	"fibercore"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type FiberServ struct {
	config  *configs.AppConfig
	app     *fiber.App
	handler handlers.HandlerParams
	db      *gorm.DB
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

func NewFiberServ(c *configs.AppConfig, h handlers.HandlerParams, db *gorm.DB) Server {
	app := fibercore.NewApp()
	//app.Use(fibercore.OpentracingMiddleware(nil))
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

func (f *FiberServ) publicHandler(router fiber.Router) {
	//router.Post("/register", dbTransactionMiddleware(f.db), f.handler.Users.Register)
	router.Post("/register", dbTransactionMiddleware(f.db), f.handler.User.Register)
	router.Get("/verifyemail/:code", f.handler.User.VerifyEmail)
}

func (f *FiberServ) authHandler(router fiber.Router) {
	router.Post("/login", f.handler.User.Login)
	router.Get("/logout", f.handler.User.Logout)
}

func (f *FiberServ) userHandler(router fiber.Router) {
	router.Get("/me", f.handler.User.Me)
}

func (f *FiberServ) configHandler() {
	f.app.Use(fibercore.SetServiceName(constant.ServiceName))
	v1 := f.app.Group("/api/v1")

	f.publicHandler(v1)

	auth := v1.Group("/auth")
	f.authHandler(auth)

	user := v1.Group("/user")
	f.userHandler(user)

}

func (f FiberServ) Start() {
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
