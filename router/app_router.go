package router

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/registry"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, registry registry.Registry) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Enable swagger docs for Dev mode
	if config.Target == "dev" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	apiRouter := e.Group("/api")
	NewUserRouter(apiRouter, registry.NewAppController().User)
	NewAdminRouter(apiRouter, registry.NewAppController().Admin)
	NewEventRouter(apiRouter, registry.NewAppController().Event)
	NewHospiRouter(apiRouter, registry.NewAppController().Hospi)
	NewCMSRouter(apiRouter, registry.NewAppController().CMS)
	NewPublicRouter(apiRouter, registry.NewAppController().Public)
	NewTShirtsRouter(apiRouter, registry.NewAppController().TShirts)
}
