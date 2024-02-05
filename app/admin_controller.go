package app

import "github.com/labstack/echo/v4"

type AdminController interface {
	Verify(c echo.Context) error
	Login(c echo.Context) error
	VerifyUser(c echo.Context) error
}
