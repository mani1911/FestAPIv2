package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type cmsControllerImpl struct {
	cmsService service.CMSService
}

func NewCMSControllerImpl(cmsService service.CMSService) app.CMSController {
	return &cmsControllerImpl{cmsService: cmsService}
}

func (impl *cmsControllerImpl) AddEvent(c echo.Context) error {
	var req dto.AddEventRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.cmsService.AddEvent(dto.AddEventRequest{
		EventID:     req.EventID,
		EventName:   req.EventName,
		IsTeam:      req.IsTeam,
		MaxTeamSize: req.MaxTeamSize,
	})

	return utils.SendResponse(c, res.Code, res.Message)
}
