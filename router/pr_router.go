package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func NewPRRouter(e *echo.Group, controller app.PRController) {
	prRoutes := e.Group("/pr")

	prRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN, models.PR))
	prRoutes.GET("/registerStatus", controller.RegisterStatus)
	prRoutes.POST("/register", controller.Register)
}
