package controller

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUserSigninRequest struct {
	Email    string `form:"user_email" query:"user_email" json:"user_email" binding:"required"`
	Password string `form:"user_password" query:"user_password" json:"user_password" binding:"required"`
}

func AuthUserSignin(c echo.Context) error {
	var req AuthUserSigninRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// Checking if both email and password are present
	if len(req.Email) == 0 || len(req.Password) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "enter username / password")
	}

	var userDetails models.User
	db := config.GetDB()

	// Checkig if user exists in the database
	if err := db.Where("Email = ? ", req.Email).First(&userDetails).Error; err != nil {
		// If user doesn't exist
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusBadRequest, "User not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in searching for user")
	}

	// Comparing passwords
	err := bcrypt.CompareHashAndPassword(userDetails.Password, []byte(req.Password))
	if err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Enter a valid password")
	}
	// Creating JWT for the user
	jwtToken, err := utils.GenerateToken(userDetails.ID, false)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
	}
	return utils.SendResponse(c, http.StatusOK, jwtToken)
}
