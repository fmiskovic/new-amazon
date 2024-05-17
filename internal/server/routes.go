package server

import (
	doc "github.com/fmiskovic/new-amz/docs/v1"
	"github.com/fmiskovic/new-amz/internal/validators"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	"net/http"
)

func initRouter() http.Handler {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// validator
	e.Validator = validators.New()

	// routes
	initRoutes(e)

	// Open API
	e.GET("/docs/*", echoSwagger.WrapHandler)
	_ = doc.Doc{} // this is a hacks to make echoSwagger happy

	return e
}

func initRoutes(r *echo.Echo) {
	dep := bootstrap()

	v1 := r.Group("/api/v1")

	account := v1.Group("/account")
	account.POST("", dep.createAccountHandler.Handle)
	account.GET("/:id", dep.getAccountByIdHandler.Handle)
	account.GET("/:id/orders", dep.searchAccountOrdersHandler.Handle)

	item := v1.Group("/item")
	item.GET("/:id", dep.getItemByIdHandler.Handle)
	item.GET("", dep.getItemsPageHandler.Handle)

	order := v1.Group("/order")
	order.POST("", dep.createOrderHandler.Handle)
	order.GET("/:id", dep.getOrderByIdHandler.Handle)
}
