package impl

import (
	"net/http"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type cmsServiceImpl struct {
	eventRepository repository.EventRepository
}

func NewCMSServiceImpl(eventRepository repository.EventRepository) service.CMSService {
	return &cmsServiceImpl{eventRepository: eventRepository}
}

func (impl *cmsServiceImpl) AddEvent(req dto.AddEventRequest) dto.Response {
	event, err := impl.eventRepository.FindEventByID(req.EventID)
	if event != nil && err == nil {
		return dto.Response{Code: http.StatusNotFound, Message: "Event Already Exists"}
	} else if event != nil && err != nil {
		return dto.Response{Code: http.StatusNotFound, Message: err}
	}

	eventDetail := models.Event{
		ID:          req.EventID,
		EventName:   req.EventName,
		IsTeam:      req.IsTeam,
		MaxTeamSize: req.MaxTeamSize,
	}

	eventAbstractDetail := models.EventAbstractDetails{
		EventID:         req.EventID,
		ForwardEmail:    req.ForwardEmail,
		MaxParticipants: req.MaxParticipants,
	}

	err = impl.eventRepository.AddEvent(eventDetail, eventAbstractDetail)

	if err != nil {
		return dto.Response{Code: http.StatusNotFound, Message: err}
	}

	return dto.Response{Code: http.StatusOK, Message: "Event Added Successfully"}
}
