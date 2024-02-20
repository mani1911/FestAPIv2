package impl

import (
	"net/http"
	"strconv"

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
// @Param			hostel_id	query		string					false	"Hostel ID"
// @Param			floor		query		string					false	"Floor Number"
// @Param			is_filled	query		string					false	"If 0, returns only free rooms. If 1, returns all."
// @Success		200			{object}	dto.GetRoomsResponse	"Success"
// @Failure		400			{string}	string					"Rooms not found"
// @Failure		500			{string}	string					"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/getRooms [get]
func (impl *hospiControllerImpl) GetRooms(c echo.Context) error {
	var hostelID int64 = -1
	var floor int64 = -1
	var isFilled int64 = -1

	var err error

	if c.QueryParams().Has("hostel_id") {
		if hostelID, err = strconv.ParseInt(c.QueryParams().Get("hostel_id"), 10, 32); err != nil {
			return utils.SendResponse(c, http.StatusBadRequest, "Invalid hostel id")
		}
	}

	if c.QueryParams().Has("floor") {
		if floor, err = strconv.ParseInt(c.QueryParams().Get("floor"), 10, 32); err != nil {
			return utils.SendResponse(c, http.StatusBadRequest, "Invalid floor")
		}
	}

	if c.QueryParams().Has("is_filled") {
		if isFilled, err = strconv.ParseInt(c.QueryParams().Get("is_filled"), 10, 32); err != nil {
			return utils.SendResponse(c, http.StatusBadRequest, "Invalid isFilled field")
		}
	}

	req := dto.GetRoomRequest{
		HostelID: int(hostelID),
		Floor:    int(floor),
		IsFilled: int(isFilled),
	}

	res := impl.HospiService.GetRooms(req)
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

// @Summary		Check-In status of a visitor.
// @Description	returns check-details of user if they've paid online
// @ID				CheckInStatus
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.CheckInStatusRequest	true	"Check in status request"
// @Success		200		{object}	dto.CheckInStatusResponse	"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		500		{object}	string						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/checkInStatus [post]
func (impl *hospiControllerImpl) CheckInStatus(c echo.Context) error {
	var req dto.CheckInStatusRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.CheckInStatus(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Allocate room for user
// @Description	Allocates room for a given user if available
// @ID				AllocateRoom
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.AllocateRoomRequest	true	"Room allocation request"
// @Success		200		{object}	string					"Success"
// @Failure		400		{object}	string					"Invalid Request"
// @Failure		500		{object}	string					"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/allocate/room [post]
func (impl *hospiControllerImpl) AllocateRoom(c echo.Context) error {
	var req dto.AllocateRoomRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.AllocateRoom(req)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Check Out
// @Description	Checks the user out of the room
// @ID				CheckOut
// @Tags			Hospi
// @Accept			json
// @Param			request	body		dto.CheckOutRequest	true	"Check Out Request"
// @Success		200		{object}	string				"Success"
// @Failure		400		{object}	string				"Invalid Request"
// @Failure		500		{object}	string				"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/hospi/checkout [post]
func (impl *hospiControllerImpl) CheckOut(c echo.Context) error {
	var req dto.CheckOutRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.HospiService.CheckOut(req)
	return utils.SendResponse(c, res.Code, res.Message)
}
