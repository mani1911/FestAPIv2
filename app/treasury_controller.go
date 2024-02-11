package app

import "github.com/labstack/echo/v4"

type TreasuryController interface {
	AddBill(c echo.Context) error
	Townscript(c echo.Context) error
}
