package router

import (
	controller "github.com/delta/FestAPI/controllers/events"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func eventsRouter(e *echo.Group) {

	eventsRoutes := e.Group("/events")
	eventsRoutes.Use(middleware.UserAuth())
	eventsRoutes.GET("/submissions/details/:event_id", controller.SubmissionDetails)
	eventsRoutes.POST("/register", controller.EventRegistration)

}
