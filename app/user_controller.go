package app

import "github.com/labstack/echo/v4"

type UserController interface {
	DAuthLogin(c echo.Context) error
	Login(c echo.Context) error
	Register(c echo.Context) error
	Update(c echo.Context) error
	ProfileDetails(c echo.Context) error
	QRgeneration(c echo.Context) error
}
