package repository

import (
	"github.com/delta/FestAPI/models"
)

type EventRepository interface {
	Register(event models.EventRegistration) error
	CheckUserRegistered(event models.EventRegistration) bool
	FindEventByID(eventID uint) (*models.Event, error)
	FindEventAbstractByID(eventID uint) (*models.EventAbstractDetails, error)
	GetUserRegisteredEvents(userID uint) ([]*models.EventRegistration, error)
}
