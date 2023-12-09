package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewCMSRouter(e *echo.Group, controller app.CMSController) {
	CMSRoutes := e.Group("/cms")
	// Protected Routes
	CMSRoutes.Use(middleware.CMSAuth())
	CMSRoutes.POST("/add_event", controller.AddEvent)
}
