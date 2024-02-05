package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

type adminControllerImpl struct {
	adminService service.AdminService
}

func NewAdminControllerImpl(adminService service.AdminService) app.AdminController {
	return &adminControllerImpl{adminService: adminService}
}

// @Summary		Verify Admin status.
// @Description	Verifies the status of an admin.
// @ID				AdminVerify
// @Tags			Admin
// @Produce		json
// @Success		200	{object}	string	"Success"
// @Failure		401	{object}	string	"Unauthorized"
// @Security		ApiKeyAuth
// @Router			/api/admin/verify [get]
func (impl *adminControllerImpl) Verify(c echo.Context) error {
	return utils.SendResponse(c, http.StatusOK, "Verified Admin")
}

// @Summary		Authenticate and log in an admin.
// @Description	Authenticates an admin user and returns a JWT token for authentication.
// @ID				AuthAdminLogin
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			request	body		dto.AuthAdminRequest	true	"Admin authentication request"
// @Success		200		{object}	string					"Success"
// @Failure		400		{object}	string					"Invalid Request"
// @Failure		500		{object}	string					"Internal Server Error"
// @Router			/api/admin/login [post]
func (impl *adminControllerImpl) Login(c echo.Context) error {
	var req dto.AuthAdminRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.adminService.Login(req)

	message := res.Message.(string)

	return utils.SendResponse(c, res.Code, message)
}

// @Summary		Get the User details.
// @Description	If providing emailid, info_type should be "email" or "jwt" for QR.
// @ID				VerifyUser
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			request	body		dto.UserInfoRequest	true	"User Info request"
// @Success		200		{object}	dto.UserInfoResponse
// @Failure		400		{object}	string	"Invalid Request"
// @Failure		500		{object}	string	"Internal Server Error"
// @Router			/api/admin/verify_user [post]
func (impl *adminControllerImpl) VerifyUser(c echo.Context) error {
	var req dto.UserInfoRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.adminService.VerifyUser(req)

	return utils.SendResponse(c, res.Code, res.Message)
}
