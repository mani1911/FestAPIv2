package controller

import (
	"net/http"

	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

// @Summary		Verify Admin status.
// @Description	Verifies the status of an admin.
// @ID				AdminVerify
// @Produce		json
// @Success		200	{object}	utils.SendResponse.DefaultResponse	"Success"
// @Security		ApiKeyAuth
// @Security		RoleAuth
// @Router			/admin/verify [get]
func AdminVerify(c echo.Context) error {
	return utils.SendResponse(c, http.StatusOK, "Verified Admin")
}
