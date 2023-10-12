package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

//TODO check for binding

type EventRegistrationRequest struct {
	EventID string `param:"event_id"`
}

func EventRegistration(c echo.Context) error {
	var req EventRegistrationRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	if len(req.EventID) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Event ID")
	}

	userInstance := c.Get("user").(*jwt.Token)
	claims := userInstance.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	db := config.GetDB()
	// Check if the event with the given event_id exists
	var event models.Event
	if err := db.Where("id = ?", req.EventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Event not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal server error")
	}

	// Check if the user has already registered for the event
	var eventRegData models.EventRegistration
	if err := db.Where("user_id = ? AND event_id = ?", userID, req.EventID).First(&eventRegData).Error; err == nil {
		return utils.SendResponse(c, http.StatusBadRequest, "You have already registered for the event")
	}
	eventID, _ := strconv.Atoi(req.EventID)

	registration := models.EventRegistration{
		EventID: uint(eventID),
		UserID:  userID,
	}

	if err := db.Create(&registration).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return utils.SendResponse(c, http.StatusOK, "You have registered successfully")
}
