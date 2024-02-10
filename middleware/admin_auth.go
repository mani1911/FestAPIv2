package middleware

import (
	"net/http"
	"strings"

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
			roleMap := make(map[string]bool)
			for _, role := range strings.Split(claims.Role, ",") {
				roleMap[role] = true
			}
			for _, requiredRole := range roles {
				if foundRole := roleMap[string(requiredRole)]; !foundRole {
					return utils.SendResponse(c, http.StatusForbidden, "Prohibited")
				}
			}
			return next(c)
		}
	}
}
