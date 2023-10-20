package controller

import (
	"net/http"
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUserLoginRequest struct {
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

// @Summary		Authenticate and log in a user.
// @Description	Authenticates a user using email and password.
// @ID				AuthUserLogin
// @Accept			json
// @Produce		json
// @Param			request	body		AuthUserLoginRequest	true	"User authentication request"
// @Success		200		{string}	string					"Success"
// @Failure		400		{string}	string					"Invalid Request"
// @Failure		500		{string}	string					"Internal Server Error"
// @Router			/user/login [post]
func AuthUserLogin(c echo.Context) error {
	var req AuthUserLoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// Check if both email and password are present
	if len(req.Email) == 0 || len(req.Password) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "enter username / password")
	}

	var userDetails models.User
	db := config.GetDB()

	// Check if user exists in the database
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
	jwtToken, err := utils.GenerateToken(userDetails.ID, false, "")
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in generating token")
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwtToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return utils.SendResponse(c, http.StatusOK, "user authenticated")
}
