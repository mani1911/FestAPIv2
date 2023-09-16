package router

import (
	controller "github.com/delta/FestAPI/controllers/user"
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func userRouter(e *echo.Group) {

	userRoutes := e.Group("/user/")
	// Public Routes
	userRoutes.GET("dauth/callback/", controller.AuthUserLogin)
	userRoutes.POST("register", controller.AuthUserRegister)
	userRoutes.POST("signin", controller.AuthUserSignin)
	//Protected Routes
	userRoutes.Use(middleware.UserAuth())
	userRoutes.PUT("update", controller.AuthUserUpdate)
}
