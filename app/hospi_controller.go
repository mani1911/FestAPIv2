package app

import "github.com/labstack/echo/v4"

type HospiController interface {
	GetHostels(c echo.Context) error
	AddUpdateHostel(c echo.Context) error
	GetRooms(c echo.Context) error
	AddUpdateRoom(c echo.Context) error
	DeleteRoom(c echo.Context) error
}
