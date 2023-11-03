package dto

type AuthUserRequest struct {
	Code string `query:"code"`
}

type AuthUserLoginRequest struct {
	Email    string `json:"user_email" binding:"required"`
	Password string `json:"user_password" binding:"required"`
}

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
