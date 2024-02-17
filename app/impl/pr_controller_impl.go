package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

type prControllerImpl struct {
	PRService service.PRService
}

func NewPRControllerImpl(prService service.PRService) app.PRController {
	return &prControllerImpl{PRService: prService}
}

// @Summary		Register PR
// @Description	Register PR for a person who has just arrived on campus
// @ID				Register
// @Tags			PR
// @Accept			json
// @Param			request	body		dto.RegisterRequest	true	"PR Register Request"
// @Success		200		{object}	string				"Success"
// @Failure		400		{object}	string				"Invalid Request"
// @Failure		500		{object}	string				"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/pr/register [post]
func (impl *prControllerImpl) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	res := impl.PRService.Register(req.UserID, req.RegAmount)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Check PR Registration Status.
// @Description	Check the PR Registration Status of a person who has just arrived on campus
// @ID				RegisterStatus
// @Tags			PR
// @Accept			json
// @Param			user_email	query		string	true	"User Email"
// @Success		200			{object}	string	"Success"
// @Failure		400			{object}	string	"Invalid Request"
// @Failure		500			{object}	string	"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/pr/registerStatus [get]
func (impl *prControllerImpl) RegisterStatus(c echo.Context) error {
	if !c.QueryParams().Has("user_email") {
		return utils.SendResponse(c, http.StatusBadRequest, "Bad Request")
	}
	userEmail := c.QueryParams().Get("user_email")
	res := impl.PRService.RegisterStatus(userEmail)
	return utils.SendResponse(c, res.Code, res.Message)
}
