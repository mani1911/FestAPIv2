package controller

import (
	"net/http"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type SubmissionDetailsResponse struct {
	EventSubmission models.EventSubmission `json:"event_submission" binding:"required"`
	UserEmail       string                 `json:"user_email" binding:"required"`
}

type SubmissionDetailsRequest struct {
	EventID string `json:"event_id" required:"true"`
}

func SubmissionDetails(c echo.Context) error {
	var req SubmissionDetailsRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	if len(req.EventID) == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Event ID")
	}

	db := config.GetDB()
	var eventSubmission models.EventSubmission
	if err := db.Where("event_id = ?", req.EventID).First(&eventSubmission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "Event not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	userInstance := c.Get("user").(*jwt.Token)
	claims := userInstance.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	// Get user_email
	var user models.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.SendResponse(c, http.StatusNotFound, "User not found")
		}
		return utils.SendResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	// Create a response struct
	response := SubmissionDetailsResponse{
		EventSubmission: eventSubmission,
		UserEmail:       user.Email,
	}

	return utils.SendResponse(c, http.StatusOK, response)
}
