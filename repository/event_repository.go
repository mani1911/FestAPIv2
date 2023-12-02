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
	IsTeamEvent(eventID uint) bool
	AddTeam(eventID uint, members []uint, teamName string, teamLeaderID uint) error
	AreUsersInTeam(eventID uint, userIDList []uint) bool
	GetTeamID(eventID uint, userID uint) (*uint, error)
	GetTeamMembers(teamID uint) ([]uint, error)
}
