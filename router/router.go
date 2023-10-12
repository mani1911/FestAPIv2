package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	apiRouter := e.Group("/api")
	userRouter(apiRouter)
	adminRouter(apiRouter)
	eventsRouter(apiRouter)
}
