package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/middleware"
	"github.com/delta/FestAPI/models"
	"github.com/labstack/echo/v4"
)

func NewHospiRouter(e *echo.Group, controller app.HospiController) {
	hospiRoutes := e.Group("/hospi")
	// Admin protected Routes
	hospiRoutes.Use(middleware.UserAuth(), middleware.AdminRoleAuth(models.ADMIN))
	// Hostel routes
	hospiRoutes.GET("/getHostels", controller.GetHostels)
	hospiRoutes.POST("/updateHostel", controller.AddUpdateHostel) // add and update hostel
	// Room routes
	hospiRoutes.GET("/getRooms", controller.GetRooms)
	hospiRoutes.POST("/updateRoom", controller.AddUpdateRoom) // add and update room
	hospiRoutes.DELETE("/deleteRoom", controller.DeleteRoom)
}
