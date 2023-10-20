package controller

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthUserUpdateRequest struct {
	Sex          string `json:"user_sex"`
	Nationality  string `json:"user_nationality"`
	Address      string `json:"user_address"`
	Pincode      string `json:"user_pincode"`
	State        string `json:"user_state"`
	City         string `json:"user_city"`
	Phone        string `json:"user_phone"`
	Degree       string `json:"user_degree"`
	Year         string `json:"user_year"`
	College      string `json:"user_college"`
	OtherCollege string `json:"user_othercollege"`
	Sponsor      string `json:"user_sponsor"`
	VoucherName  string `json:"user_voucher_name"`
	ReferralCode string `json:"user_referral_code"`
	Country      string `json:"user_country"`
}

// @Summary		Update user information.
// @Description	Update user information with the provided details.
// @ID				AuthUserUpdate
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Param			request	body		AuthUserUpdateRequest	true	"User update request"
// @Success		200		{string}	string					"Success"
// @Failure		400		{string}	string					"Invalid Request"
// @Failure		500		{string}	string					"Internal Server Error"
// @Router			/user/update [put]
func AuthUserUpdate(c echo.Context) error {
	// obtaining user id from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	var req AuthUserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// Check if required fields are present
	if len(req.Sex) == 0 ||
		len(req.Nationality) == 0 ||
		len(req.Address) == 0 ||
		len(req.Pincode) == 0 ||
		len(req.State) == 0 ||
		len(req.City) == 0 ||
		len(req.Phone) == 0 ||
		len(req.Degree) == 0 ||
		len(req.Year) == 0 ||
		len(req.College) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	var userDetails models.User
	var collegeDetails models.College
	db := config.GetDB()

	// Checking if user exists in DB
	if err := db.Where("ID = ? ", userID).First(&userDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusBadRequest, "Invalid user id")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server error")
	}
	// Checking if college exists in DB
	if err := db.Where("Name = ?", req.College).First(&collegeDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusBadRequest, "Invalid college name")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Error in finding College")
	}

	// Updating Details
	userDetails.College = collegeDetails
	userDetails.OtherCollege = req.OtherCollege
	userDetails.Gender = models.Gender(req.Sex)
	userDetails.Country = req.Country
	userDetails.State = req.State
	userDetails.City = req.City
	userDetails.Address = req.Address
	userDetails.Pincode = req.Pincode
	userDetails.Phone = req.Phone
	userDetails.Sponsor = req.Sponsor
	userDetails.VoucherName = req.VoucherName
	userDetails.ReferralCode = req.ReferralCode
	userDetails.Degree = req.Degree
	userDetails.Year = req.Year
	userDetails.Nationality = req.Nationality
	// Storing updates in DB
	if err := db.Save(&userDetails).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Failed to update user")
	}
	return utils.SendResponse(c, http.StatusOK, "Account Updated")
}
