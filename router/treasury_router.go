package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func NewTreasuryRouter(e *echo.Group, controller app.TreasuryController) {
	//protected routes
	treasuryRoutes := e.Group("/treasury", middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN))
	treasuryRoutes.POST("/addBill", controller.AddBill)
	e.POST("/treasury/townscript", controller.Townscript) // should be public
}
