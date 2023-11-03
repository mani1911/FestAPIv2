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
		return fmt.Errorf("Error creating User")
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
