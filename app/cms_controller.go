package app

import "github.com/labstack/echo/v4"

type CMSController interface {
	AddEvent(c echo.Context) error
}
