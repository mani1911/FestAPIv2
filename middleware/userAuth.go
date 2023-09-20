package middleware

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JWTCustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return utils.SendResponse(c, http.StatusForbidden, "Prohibited")
		},
	})
}
