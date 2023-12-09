package impl

import (
	"fmt"

	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewEventRepositoryImpl(DB *gorm.DB) repository.EventRepository {
	return &eventRepositoryImpl{DB: DB}
}

type eventRepositoryImpl struct {
	*gorm.DB
}

func (repository *eventRepositoryImpl) CheckUserRegistered(details models.EventRegistration) bool {

	var eventRegData models.EventRegistration

	// Check if the user has already registered for the event
	if err := repository.DB.Where("user_id = ? AND event_id = ?", details.UserID, details.EventID).First(&eventRegData).Error; err == nil {
		return true
	}
	return false
}

func (repository *eventRepositoryImpl) Register(details models.EventRegistration) error {

	// Register User for Event
	if err := repository.DB.Create(&details).Error; err != nil {
		return fmt.Errorf("error creating User")
	}
	return nil
}

func (repository *eventRepositoryImpl) FindEventAbstractByID(eventID uint) (*models.EventAbstractDetails, error) {
	var eventSubmission models.EventAbstractDetails

	// Find event abstract by Id
	if err := repository.DB.Where("event_id = ?", eventID).First(&eventSubmission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &eventSubmission, nil
}

func (repository *eventRepositoryImpl) FindEventByID(eventID uint) (*models.Event, error) {
	var event models.Event

	// Find event by Id
	if err := repository.DB.Where("id = ?", eventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (repository *eventRepositoryImpl) GetUserRegisteredEvents(userID uint) ([]*models.EventRegistration, error) {
	var res []*models.EventRegistration

	err := repository.DB.Preload("Event").Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return res, err

}

func (repository *eventRepositoryImpl) IsTeamEvent(eventID uint) bool {
	var event models.Event

	// Find event by Id
	if err := repository.DB.Where("id = ?", eventID).First(&event).Error; err != nil {
		return false
	}

	return event.IsTeam
}

func (repository *eventRepositoryImpl) AddTeam(eventID uint, members []uint, teamName string, teamLeaderID uint) error {
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create the team
	team := models.EventTeam{
		EventID:      eventID,
		TeamName:     teamName,
		TeamLeaderID: teamLeaderID,
	}

	if err := tx.Create(&team).Error; err != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()
		return err
	}

	// Add team members
	for _, member := range members {
		teamMember := models.EventTeamMember{
			TeamID: team.ID,
			UserID: member,
		}
		if err := tx.Create(&teamMember).Error; err != nil {
			// Rollback the transaction if there's an error
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	err := tx.Commit().Error

	return err
}

func (repository *eventRepositoryImpl) AreUsersInTeam(eventID uint, userIDList []uint) bool {
	// Check if user is in a team whose event id is the same as
	// the given event id

	// Get all the teams in which a user inside userIDList is present
	var teamIDs []uint

	err := repository.DB.Model(&models.EventTeamMember{}).
		Where("user_id IN (?)", userIDList).
		Pluck("DISTINCT team_id", &teamIDs).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		fmt.Println("AreUsersInTeam: ", err.Error())
		return true

	}

	var team models.EventTeam

	err = repository.DB.Model(&models.EventTeam{}).
		Where("team_id IN (?)", teamIDs).
		Where("event_id = ?", eventID).
		First(&team).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// No matching team found
			return false
		}
		// Handle the error
		fmt.Println("AreUsersInTeam: ", err.Error())
		return true

	}
	// A team with the same EventID exists
	return true

}

func (repository *eventRepositoryImpl) GetTeamID(eventID uint, userID uint) (*uint, error) {
	var team models.EventTeam
	result := repository.DB.Joins("JOIN event_team_members on event_team_members.team_id = event_teams.team_id").
		Where("event_team_members.user_id = ? AND event_teams.event_id = ?", userID, eventID).
		First(&team)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error

	}

	teamID := team.TeamID

	return &teamID, nil
}

func (repository *eventRepositoryImpl) GetTeamMembers(teamID uint) ([]uint, error) {
	var teamMembers []models.EventTeamMember
	var teamMemberIDs []uint

	result := repository.DB.Where("team_id = ?", teamID).Find(&teamMembers)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error

	}

	for _, member := range teamMembers {
		teamMemberIDs = append(teamMemberIDs, member.UserID)
	}

	return teamMemberIDs, nil
}

func (repository *eventRepositoryImpl) AddEvent(event models.Event, eventAbstractDetails models.EventAbstractDetails) error {

	tx := repository.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating event")
	}

	if err := tx.Create(&eventAbstractDetails).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating event abstract details")
	}

	err := tx.Commit().Error

	return err
}
