package dto

type UserInfoRequest struct {
	Info     string `json:"user_info" binding:"required"`
	InfoType string `json:"info_type" binding:"required"`
}
type AuthAdminRequest struct {
	Username string `json:"admin_username" binding:"required"`
	Password string `json:"admin_password" binding:"required"`
}
