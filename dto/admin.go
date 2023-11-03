package dto

type AuthAdminRequest struct {
	Username string `json:"admin_username" binding:"required"`
	Password string `json:"admin_password" binding:"required"`
}
