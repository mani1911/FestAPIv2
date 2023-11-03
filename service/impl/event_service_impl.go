package impl

import (
	"net/http"

	dto "github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type eventServiceImpl struct {
	eventRepository repository.EventRepository
	userRepository  repository.UserRepository
}

func NewEventServiceImpl(
	eventRepository repository.EventRepository,
	userRepository repository.UserRepository) service.EventService {
	return &eventServiceImpl{
		eventRepository: eventRepository,
		userRepository:  userRepository,
	}
}

func (impl *eventServiceImpl) Register(req dto.EventRegistrationDTO) dto.Response {

	event, err := impl.eventRepository.FindEventByID(req.EventID)

	if event == nil && err == nil {
		return dto.Response{Code: http.StatusNotFound, Message: "Event Not Found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	eventRegistrationDetails := models.EventRegistration{
		EventID: req.EventID,
		UserID:  req.UserID,
	}
	isUserRegistered := impl.eventRepository.CheckUserRegistered(eventRegistrationDetails)
	if isUserRegistered {
		return dto.Response{Code: http.StatusBadRequest, Message: "You have already registered for the event"}
	}

	err = impl.eventRepository.Register(eventRegistrationDetails)
	if err != nil {
		return dto.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
	return dto.Response{Code: http.StatusOK, Message: "User Registered to Event Successfully"}
}

func (impl *eventServiceImpl) AbstractDetails(req dto.AbstractDetailsRequest) dto.Response {

	eventSubmission, err := impl.eventRepository.FindEventAbstractByID(req.EventID)
	if eventSubmission == nil && err == nil {
		return dto.Response{Code: http.StatusNotFound, Message: "Event Not Found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	// Create a Response
	res := dto.AbstractDetailsResponse{
		ForwardEmail:    eventSubmission.ForwardEmail,
		MaxParticipants: eventSubmission.MaxParticipants}

	return dto.Response{Code: http.StatusOK, Message: res}
}

func (impl *eventServiceImpl) UserEventDetails(userID uint) dto.Response {

	var res []dto.GetEventDetailsResponse

	// checking if user has registered for any event
	userDetails, err := impl.userRepository.FindByID(userID)
	if userDetails == nil && err == nil {
		return dto.Response{Code: http.StatusOK, Message: "Invalid User"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	// checking if user has registered for any event

	userEventDetails, err := impl.eventRepository.GetUserRegisteredEvents(userID)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	for _, event := range userEventDetails {

		eventDetail := dto.GetEventDetailsResponse{
			EventID:   event.EventID,
			EventName: event.Event.EventName,
		}
		res = append(res, eventDetail)
	}

	return dto.Response{Code: http.StatusAccepted, Message: res}

}
