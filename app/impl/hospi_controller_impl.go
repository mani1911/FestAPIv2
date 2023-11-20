package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

type hospiControllerImpl struct {
	HospiService service.HospiService
}

func NewHospiControllerImpl(hospiService service.HospiService) app.HospiController {
	return &hospiControllerImpl{HospiService: hospiService}
}

// @Summary		Get all the Hostels
// @Description	Retrieve the details of the hostels.
// @ID				GetHostels
// @Tags			Hospi
// @Produce		json
// @Success		200	{object}	dto.GetHostelsResponse	"Success"
// @Failure		400	{string}	string					"Hostels not found"
// @Failure		500	{string}	string					"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/getHostels [get]
func (impl *hospiControllerImpl) GetHostels(c echo.Context) error {
	res := impl.HospiService.GetHostels()
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Add/Update a new hostel.
// @Description	Add/Update a new hostel with the provided details.
// @ID				AddUpdateHostel
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.AddUpdateHostelRequest	true	"Add/update hostel request"
// @Success		200		{object}	string						"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		500		{object}	string						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/updateHostel [post]
func (impl *hospiControllerImpl) AddUpdateHostel(c echo.Context) error {
	var req dto.AddUpdateHostelRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.AddUpdateHostel(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Get all the Rooms
// @Description	Retrieve the details of the rooms along with the hostel name.
// @ID				GetRooms
// @Tags			Hospi
// @Produce		json
// @Success		200	{object}	dto.GetRoomsResponse	"Success"
// @Failure		400	{string}	string					"Rooms not found"
// @Failure		500	{string}	string					"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/getRooms [get]
func (impl *hospiControllerImpl) GetRooms(c echo.Context) error {
	res := impl.HospiService.GetRooms()
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Add/Update a new room.
// @Description	Add/Update a new room with the provided details.
// @ID				AddUpdateRoom
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.AddUpdateRoomRequest	true	"Add/update room request"
// @Success		200		{object}	string						"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		500		{object}	string						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/updateRoom [post]
func (impl *hospiControllerImpl) AddUpdateRoom(c echo.Context) error {
	var req dto.AddUpdateRoomRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.AddUpdateRoom(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Delete a room.
// @Description	Delete a room with the provided ID.
// @ID				DeleteRoom
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.DeleteRoomRequest	true	"Delete room request"
// @Success		200		{object}	string					"Success"
// @Failure		400		{object}	string					"Invalid Request"
// @Failure		500		{object}	string					"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/deleteRoom [delete]
func (impl *hospiControllerImpl) DeleteRoom(c echo.Context) error {
	var req dto.DeleteRoomRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.DeleteRoom(req)
	return utils.SendResponse(c, res.Code, res.Message)
}
