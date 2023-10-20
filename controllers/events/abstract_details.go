package controller

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type AbstractDetailsResponse struct {
	ForwardEmail    string `json:"forward_email"`
	MaxParticipants uint   `json:"max_participants"`
}

type AbstractDetailsRequest struct {
	EventID uint `param:"event_id"`
}

func AbstractDetails(c echo.Context) error {
	var req AbstractDetailsRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	db := config.GetDB()
	var eventSubmission models.EventAbstractDetails
	if err := db.Where("event_id = ?", req.EventID).First(&eventSubmission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Event not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	// Create a response struct
	response := AbstractDetailsResponse{
		ForwardEmail:    eventSubmission.ForwardEmail,
		MaxParticipants: eventSubmission.MaxParticipants,
	}

	return utils.SendResponse(c, http.StatusOK, response)
}
