package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Group, controller app.UserController) {

	userRoutes := e.Group("/user")
	// Public Routes
	userRoutes.GET("/dauth/callback", controller.DAuthLogin)
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/login", controller.Login)
	userRoutes.POST("/verify", controller.VerifyEmail)
	userRoutes.POST("/changePassword", controller.ChangePassword)
	//Protected Routes
	userRoutes.Use(middleware.UserAuth())
	userRoutes.GET("/details", controller.ProfileDetails)
	userRoutes.GET("/qr", controller.QRgeneration)
	userRoutes.PATCH("/update", controller.Update)
}
