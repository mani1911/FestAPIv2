package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type eventControllerImpl struct {
	eventService service.EventService
}

func NewEventControllerImpl(eventService service.EventService) app.EventController {
	return &eventControllerImpl{eventService: eventService}
}

// @Summary		Register the user for an event.
// @Description	Register the user for the specified event.
// @ID				EventRegister
// @Tags			Event
// @Produce		json
// @Param			request	body		dto.EventRegistrationRequest	true	"Event Registration Request"
// @Success		200		{object}	string							"Success"
// @Failure		400		{string}	string							"Invalid Request"
// @Failure		401		{object}	string							"Unauthorized"
// @Failure		500		{string}	string							"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/events/register [post]
func (impl *eventControllerImpl) Register(c echo.Context) error {
	var req dto.EventRegistrationRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	userInstance := c.Get("user").(*jwt.Token)
	claims := userInstance.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	res := impl.eventService.Register(dto.EventRegistrationDTO{
		EventID: req.EventID,
		UserID:  userID,
	})
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Get Event's Abstract Details
// @Description	Retrieve the details of the abstract for the specified event.
// @ID				EventAbstractDetails
// @Tags			Event
// @Produce		json
// @Param			event_id	path		int							true	"EventID"
// @Success		200			{object}	dto.AbstractDetailsResponse	"Success"
// @Failure		400			{string}	string						"Invalid Request"
// @Failure		401			{object}	string						"Unauthorized"
// @Failure		404			{string}	string						"Event not found"
// @Failure		500			{string}	string						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/events/abstract/details/{event_id} [get]
func (impl *eventControllerImpl) AbstractDetails(c echo.Context) error {
	var req dto.AbstractDetailsRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.eventService.AbstractDetails(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Get details of events registered by a user.
// @Description	Retrieve a list of events registered by the user.
// @ID				UserEventDetails
// @Tags			Event, User
// @Produce		json
// @Success		200	{object}	[]dto.GetEventDetailsResponse	"Success"
// @Failure		400	{string}	string							"Invalid Request"
// @Failure		500	{string}	string							"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/events/user/registered [get]
func (impl *eventControllerImpl) UserEventDetails(c echo.Context) error {

	// obtaining user id from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	res := impl.eventService.UserEventDetails(userID)

	return utils.SendResponse(c, res.Code, res.Message)
}
