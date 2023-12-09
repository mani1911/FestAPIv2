package middleware

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

func CMSAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header["X-Cms-Token"][0] != config.CMSToken {
				return utils.SendResponse(c, http.StatusForbidden, "Prohibited")
			}
			return next(c)
		}
	}
}
