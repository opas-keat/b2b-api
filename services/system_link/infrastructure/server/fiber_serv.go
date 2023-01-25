package server

import (
	"fibercore"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"systemlink/configs"
	"systemlink/constant"
	"systemlink/handlers"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
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

func (f *FiberServ) presaleordersHandler(router fiber.Router) {
	router.Post("/", f.handler.PreSaleOrder.CreatePreSaleOrder)
}

func (f *FiberServ) dealerHandler(router fiber.Router) {
	router.Post("/", f.handler.Dealer.ListDealerByCode)
	router.Post("/internal", f.handler.Dealer.ListDealerByCodeInternal)
	router.Get("/:dealer_code", f.handler.Dealer.GetDealerByCode)

}

func (f *FiberServ) productHandler(router fiber.Router) {
	router.Get("/bams", f.handler.Product.ListBrandAndModel)
	router.Get("/:product_code", f.handler.Product.GetProductByCode)

}

func (f *FiberServ) shippingsHandler(router fiber.Router) {
	router.Post("/", f.handler.Shipping.ListShippingByRegion)
	router.Get("/:code", f.handler.Shipping.GetShippingByCode)

}

func (f *FiberServ) configHandler() {
	f.app.Use(fibercore.SetServiceName(constant.ServiceName))
	v1 := f.app.Group("/api/v1")

	dealers := v1.Group("/dealers")
	f.dealerHandler(dealers)

	products := v1.Group("/products")
	f.productHandler(products)

	shippings := v1.Group("/shippings")
	f.shippingsHandler(shippings)

	presaleorders := v1.Group("/presaleorders")
	f.presaleordersHandler(presaleorders)
}

func (f *FiberServ) Start() {
	go func() {
		if err := f.app.Listen(fmt.Sprintf(":%v", f.config.Port)); err != nil {
			slog.Error("Error Start Fiber Server", err, "status", 500)
		}
	}()

	// go func() {
	// 	if err := f.app.Listen(fmt.Sprintf(":%v", f.config.Port)); err != nil {
	// 		log.Panic(err)
	// 	}
	// }()

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
