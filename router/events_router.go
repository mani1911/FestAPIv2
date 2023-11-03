package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewEventRouter(e *echo.Group, controller app.EventController) {
	eventsRoutes := e.Group("/events")
	// Protected Routes
	eventsRoutes.Use(middleware.UserAuth())
	eventsRoutes.GET("/abstract/details/:event_id", controller.AbstractDetails)
	eventsRoutes.POST("/register", controller.Register)
	eventsRoutes.GET("/user/registered", controller.UserEventDetails)
}
