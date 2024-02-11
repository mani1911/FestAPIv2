package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func NewHospiRouter(e *echo.Group, controller app.HospiController) {
	hospiRoutes := e.Group("/hospi")

	hospiRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN, models.PR))
	hospiRoutes.GET("/getHostels", controller.GetHostels)
	hospiRoutes.GET("/getRooms", controller.GetRooms)

	hospiRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN, models.PR, models.CORE))
	hospiRoutes.POST("/updateHostel", controller.AddUpdateHostel) // add and update hostel
	hospiRoutes.POST("/updateRoom", controller.AddUpdateRoom)     // add and update room
	hospiRoutes.DELETE("/deleteRoom", controller.DeleteRoom)

	hospiRoutes.POST("/checkIn", controller.CheckIn)
}
