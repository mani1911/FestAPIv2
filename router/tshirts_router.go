package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewTShirtsRouter(e *echo.Group, controller app.TShirtsController) {

	tshirtRoutes := e.Group("/tshirt")

	//Protected Routes
	tshirtRoutes.Use(middleware.UserAuth())
	tshirtRoutes.POST("/updateSize", controller.UpdateSize)
}
