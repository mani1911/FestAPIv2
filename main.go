package main

import (
	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/router"
	"github.com/delta/FestAPI/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())

	config.InitConfig()
	config.ConnectDB()
	config.MigrateDB()

	router.NewRouter(server)

	utils.InitLogger(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))

}
