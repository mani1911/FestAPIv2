package app

import "github.com/labstack/echo/v4"

type PRController interface {
	Register(c echo.Context) error
	RegisterStatus(c echo.Context) error
}
