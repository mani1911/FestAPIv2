package utils

import "github.com/labstack/echo/v4"

func SendResponse(c echo.Context, code int, message interface{}) error {

	// DefaultResponse
	//	@Description	message
	type DefaultResponse struct {
		// Default Response
		Message interface{} `json:"message"`
	}

	return c.JSON(code, DefaultResponse{Message: message})
}
