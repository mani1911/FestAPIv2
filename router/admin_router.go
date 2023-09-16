package router

import (
	"github.com/delta/FestAPI/middleware"
	"github.com/labstack/echo/v4"
)

func adminRouter(e *echo.Group) {
	adminRoutes := e.Group("/admin/")
	adminRoutes.Use(middleware.AdminAuth())
}
