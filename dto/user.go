package dto

import (
	"time"

	"github.com/delta/FestAPI/models"
)

type AuthUserRequest struct {
	Code string `query:"code"`
}

type AuthUserLoginRequest struct {
	Email    string `json:"user_email" binding:"required"`
	Password string `json:"user_password" binding:"required"`
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type AuthUserRegisterRequest struct {
	Username      string `json:"user_name"`
	Email         string `json:"user_email"`
	Fullname      string `json:"user_fullname"`
	Password      string `json:"user_password"`
	Sex           string `json:"user_sex"`
	Nationality   string `json:"user_nationality"`
	Address       string `json:"user_address"`
	Pincode       string `json:"user_pincode"`
	State         string `json:"user_state"`
	City          string `json:"user_city"`
	Phone         string `json:"user_phone"`
	Degree        string `json:"user_degree"`
	Year          string `json:"user_year"`
	College       string `json:"user_college"`
	OtherCollege  string `json:"user_othercollege"`
	Sponsor       string `json:"user_sponsor"`
	VoucherName   string `json:"user_voucher_name"`
	ReferralCode  string `json:"user_referral_code"`
	Country       string `json:"user_country"`
	RecaptchaCode string `json:"user_recaptcha_code"`
}

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

type UserInfoResponse struct {
	ID           uint
	Name         string
	FullName     string
	CollegeID    uint
	OtherCollege string
	Email        string
	College      CollegeResponse
	Gender       models.Gender
	Country      string
	State        string
	City         string
	Address      string
	Pincode      string
	Phone        string
	Password     []byte
	Sponsor      string
	VoucherName  string
	ReferralCode string
	Degree       string
	Year         string
	Nationality  string
	IsDauth      bool
	TShirtSize   string
}
