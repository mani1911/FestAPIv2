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

type AuthUserRegisterRequest struct {
	Username     string `json:"user_name"`
	Email        string `json:"user_email"`
	Fullname     string `json:"user_fullname"`
	Password     string `json:"user_password"`
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

// @Summary		Register a new user.
// @Description	Register a new user with the provided details.
// @ID				AuthUserRegister
// @Accept			json
// @Produce		json
// @Param			request	body		AuthUserRegisterRequest	true	"User registration request"
// @Success		200		{string}	string					"Success"
// @Failure		400		{string}	string					"Invalid Request"
// @Failure		500		{string}	string					"Internal Server Error"
// @Router			/user/register [post]
func AuthUserRegister(c echo.Context) error {
	var req AuthUserRegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	// checking if required fields are present
	if len(req.Username) == 0 ||
		len(req.Email) == 0 ||
		len(req.Fullname) == 0 ||
		len(req.Password) == 0 ||
		len(req.Sex) == 0 ||
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
	// Checking if user exists in the database or not
	if err := db.Where("Email = ? ", req.Email).First(&userDetails).Error; err != nil {
		// Check if record doesn't exist => new user
		if err == gorm.ErrRecordNotFound {
			// Creating password hash
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
			if err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server error")
			}
			// Fetching College details
			if err := db.Where("Name = ?", req.College).First(&collegeDetails).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return utils.SendResponse(c, http.StatusBadRequest, "Enter a valid college name")
				}
				return utils.SendResponse(c, http.StatusInternalServerError, "Error in finding College")
			}
			// Creating new user
			userReg := models.User{
				Name:         req.Username,
				FullName:     req.Fullname,
				College:      collegeDetails,
				OtherCollege: req.OtherCollege,
				Email:        req.Email,
				Gender:       models.Gender(req.Sex),
				Country:      req.Country,
				State:        req.State,
				City:         req.City,
				Address:      req.Address,
				Pincode:      req.Pincode,
				Phone:        req.Phone,
				Password:     passwordHash,
				Sponsor:      req.Sponsor,
				VoucherName:  req.VoucherName,
				ReferralCode: req.ReferralCode,
				Degree:       req.Degree,
				Year:         req.Year,
				Nationality:  req.Nationality,
			}
			if err := db.Create(&userReg).Error; err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user")

			}
			return utils.SendResponse(c, http.StatusOK, "Account Created")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}
	return utils.SendResponse(c, http.StatusBadRequest, "Already Created Account")
}
