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

type AuthUserRequest struct {
	Code string `query:"code"`
}

func AuthUserLogin(c echo.Context) error {
	var req AuthUserRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Server Error")
	}
	// Fetching code from header
	code := req.Code
	if len(code) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// Obtaining access token from dauth server
	token, err := utils.GetDAuthToken(code)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in Authenticating user")
	}
	// Obtaining user details from dauth server
	user, err := utils.GetDAuthUser(token.AccessToken)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in Authenticating User")
	}

	Name := user.Name
	Email := user.Email
	Gender := user.Gender
	Phone := user.Phone
	if len(Name) == 0 || len(Email) == 0 {
		return utils.SendResponse(c, http.StatusInternalServerError, "User not found")
	}

	db := config.GetDB()
	var userDetails models.User
	var collegeDetails models.College

	// Checking if user exist in db
	if err = db.Where("Email = ? ", Email).First(&userDetails).Error; err != nil {
		// If user doesn't exist i.e new user
		if err == gorm.ErrRecordNotFound {
			// Fetching college details
			if err = db.Where("Name = ?", "National Institute of Technology, Tiruchirapalli").First(&collegeDetails).Error; err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user")
			}
			// Creating a password for each user
			password := Email + "may4cebewithu"
			passHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
			if err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user")
			}
			// Creating new User object
			userReg := models.User{
				Email:     Email,
				Name:      Name,
				CollegeID: collegeDetails.ID,
				Gender:    models.Gender(Gender),
				Phone:     Phone,
				Password:  passHash,
			}
			// Storing new user in the database
			if err := db.Create(&userReg).Error; err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user")
			}
			// Creating JWT Token for the new user
			jwtToken, err := utils.GenerateToken(userReg.ID, false)
			if err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
			}
			return utils.SendResponse(c, http.StatusOK, jwtToken)
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in finding User")
	}

	// User already exists in the database
	// Creating JWT for the existing user
	jwtToken, err := utils.GenerateToken(userDetails.ID, false)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
	}
	return utils.SendResponse(c, http.StatusOK, jwtToken)
}
