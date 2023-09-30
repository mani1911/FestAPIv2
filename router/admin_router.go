package router

import (
	controller "github.com/delta/FestAPI/controllers/admin"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func adminRouter(e *echo.Group) {

	adminRoutes := e.Group("/admin")
	// Public Routes
	adminRoutes.POST("/login", controller.AuthAdminLogin)
	//Protected Routes
	adminRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN))
	adminRoutes.GET("/verify", controller.AdminVerify)
}
