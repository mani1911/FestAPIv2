package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Group, controller app.UserController) {

	userRoutes := e.Group("/user")
	// Public Routes
	userRoutes.GET("/dauth/callback/", controller.DAuthLogin)
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	//Protected Routes
	userRoutes.Use(middleware.UserAuth())
	userRoutes.PATCH("/update", controller.Update)
}
