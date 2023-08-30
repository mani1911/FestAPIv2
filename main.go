package main

import (
	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := echo.New()

	utils.InitLogger(server)
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())

	config.InitConfig()
	config.ConnectDB()
	config.MigrateDB()

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
