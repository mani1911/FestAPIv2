package utils

import (
	"net/http"
	"time"

	"github.com/delta/FestAPI/config"
)

func GenerateCookie(message string) *http.Cookie {
	// Creating HTTPOnly Cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = message
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Domain = config.CookieDomain
	cookie.SameSite = http.SameSiteNoneMode

	return cookie
}
