package app

import "github.com/labstack/echo/v4"

type PublicController interface {
	Colleges(c echo.Context) error
}
