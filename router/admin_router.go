package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func NewAdminRouter(e *echo.Group, controller app.AdminController) {

	adminRoutes := e.Group("/admin")
	// Public Routes
	adminRoutes.POST("/login", controller.Login)
	//Protected Routes
	adminRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN))
	adminRoutes.GET("/verify", controller.Verify)
}
