package app

import "github.com/labstack/echo/v4"

type EventController interface {
	Register(c echo.Context) error
	AbstractDetails(c echo.Context) error
	UserEventDetails(c echo.Context) error
	Status(c echo.Context) error
}
