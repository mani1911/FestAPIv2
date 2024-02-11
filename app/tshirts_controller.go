package app

import "github.com/labstack/echo/v4"

type TShirtsController interface {
	UpdateSize(c echo.Context) error
}
