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

type AuthAdminAdminRequest struct {
	Username string `json:"admin_username"`
	Password string `json:"admin_password"`
}

// @Summary		Authenticate and log in an admin.
// @Description	Authenticates an admin using username and password, and returns a JWT token for authentication.
// @ID				AuthAdminLogin
// @Accept			json
// @Produce		json
// @Param			request	body		AuthAdminAdminRequest	true	"Admin authentication request"
// @Success		200		{string}	string					"Success"
// @Failure		400		{string}	string					"Invalid Request"
// @Failure		500		{string}	string					"Internal Server Error"
// @Router			/admin/login [post]
func AuthAdminLogin(c echo.Context) error {
	var req AuthAdminAdminRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// Checking if both email and password are present
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "enter username / password")
	}

	var adminDetails models.Admin
	db := config.GetDB()

	// Checking if admin exists in the database
	if err := db.Where("Username = ? ", req.Username).First(&adminDetails).Error; err != nil {
		// If admin doesn't exist
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusBadRequest, "Admin not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in searching for admin")
	}

	// Comparing passwords
	err := bcrypt.CompareHashAndPassword(adminDetails.Password, []byte(req.Password))
	if err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Enter a valid password")
	}
	// Creating JWT for the user along with role
	jwtToken, err := utils.GenerateToken(adminDetails.ID, true, adminDetails.Role)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
	}

	// Creating HTTPOnly Cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwtToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return utils.SendResponse(c, http.StatusOK, "User Authenticated Successfully")
}
