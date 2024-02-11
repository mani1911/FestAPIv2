package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

type treasuryControllerImpl struct {
	treasuryService service.TreasuryService
}

func NewTreasuryControllerImpl(treasuryService service.TreasuryService) app.TreasuryController {
	return &treasuryControllerImpl{treasuryService: treasuryService}
}

// @Summary		Add a Bill.
// @Description	Add a bill with its purpose: checkIn/checkOut/discount/fine/eventPass.
// @ID				AddBill
// @Tags			Treasury
// @Accept			json
// @Param			request	body		dto.AddBillRequest	true	"Add bill request"
// @Success		200		{object}	string				"Success"
// @Failure		400		{object}	string				"Invalid Request"
// @Failure		500		{object}	string				"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/treasury/addBill [post]
func (impl *treasuryControllerImpl) AddBill(c echo.Context) error {
	var req dto.AddBillRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.treasuryService.AddBill(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Make a payment using Townscript.
// @Description	Make a payment using Townscript.
// @ID				Townscript
// @Tags			Treasury
// @Accept			json
// @Param			request	body		dto.TownScriptRequest	true	"Townscript request"
// @Param			secret	query		string					false	"Secret Token"
// @Success		200		{object}	string					"Success"
// @Failure		400		{object}	string					"Invalid Request"
// @Failure		401		{object}	string					"Unauthorized"
// @Failure		500		{object}	string					"Internal Server Error"
// @Router			/api/treasury/townscript [post]
func (impl *treasuryControllerImpl) Townscript(c echo.Context) error {

	var req dto.TownScriptRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	token := c.Request().URL.Query().Get("secret")
	if token != config.JWTSecret {
		return utils.SendResponse(c, http.StatusUnauthorized, "Unauthorized")
	}
	res := impl.treasuryService.Townscript(req)
	return utils.SendResponse(c, res.Code, res.Message)
}
