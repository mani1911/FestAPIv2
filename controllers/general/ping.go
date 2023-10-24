package controllers

import (
	"net/http"

	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

// @Summary		Ping
// @Description	Checks if the server is up and running
// @Produce		json
// @Param			code	query		string	true	"DAuth code"
// @Success		200		{object}	utils.SendResponse.DefaultResponse	"Success"
// @Router			/ping [get]
func Ping(c echo.Context) error {
	return utils.SendResponse(c, http.StatusOK, "Pong")
}
