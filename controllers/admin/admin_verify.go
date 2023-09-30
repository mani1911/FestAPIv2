package controller

import (
	"net/http"

	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

func AdminVerify(c echo.Context) error {
	return utils.SendResponse(c, http.StatusOK, "Verified Admin")
}
