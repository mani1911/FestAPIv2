package middleware

import (
	"net/http"

	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminRoleAuth(roles ...models.AdminRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*utils.JWTCustomClaims)
			// if user is not an admin, prohibit access
			if !claims.Admin {
				return utils.SendResponse(c, http.StatusForbidden, "Prohibited")
			}
			// check if user has the appropriate role
			for _, requiredRole := range roles {
				if requiredRole == claims.Role {
					// User has the required role, allow access
					return next(c)
				}
			}
			return utils.SendResponse(c, http.StatusForbidden, "Prohibited")
		}
	}
}
