package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewTShirtsRouter(e *echo.Group, controller app.TShirtsController) {

	userRoutes := e.Group("/tshirt")

	//Protected Routes
	userRoutes.Use(middleware.UserAuth())
	userRoutes.POST("/updateSize", controller.UpdateSize)
}
