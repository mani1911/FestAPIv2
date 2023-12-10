package router

import (
	"github.com/delta/FestAPI/app"
	"github.com/labstack/echo/v4"
)

func NewPublicRouter(e *echo.Group, controller app.PublicController) {
	e.GET("/colleges", controller.Colleges)
}
