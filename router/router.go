package router

import "github.com/labstack/echo/v4"

func NewRouter(e *echo.Echo) {
	// Open Routes
	_ = e.Group("/api/v1")
}
