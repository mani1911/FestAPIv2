package router

import (
	controllers "github.com/delta/FestAPI/controllers/general"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	apiRouter := e.Group("/api")
	apiRouter.GET("/ping", controllers.Ping)
	userRouter(apiRouter)
	adminRouter(apiRouter)
	eventsRouter(apiRouter)
}
