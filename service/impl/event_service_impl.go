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

	if !event.IsTeam {
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
	if len(req.TeamName) == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Team Name cannot be empty"}
	}

	if len(req.TeamMembers) == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Team Members cannot be empty"}
	}

	if len(req.TeamMembers) > int(event.MaxTeamSize) {
		return dto.Response{Code: http.StatusBadRequest, Message: "Team Members cannot be more than max team size"}
	}

	isUserRegistered := impl.eventRepository.CheckUserRegistered(eventRegistrationDetails)
	if isUserRegistered {
		return dto.Response{Code: http.StatusBadRequest, Message: "You have already registered for the event"}
	}

	userIDs := make([]uint, len(req.TeamMembers))
	for i, member := range req.TeamMembers {
		user, err := impl.userRepository.FindByEmail(member)
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}

		if user == nil {
			return dto.Response{Code: http.StatusNotFound, Message: "User not found"}
		}

		userIDs[i] = user.ID
	}

	// This is to ensure the invariant that the team leader must be present in
	// the team also. Forgetting this check lead to some bugs, thats why we
	// add it here to enforce it.

	// Check if req.UserID is already present in the list of userIDs
	found := false // flag to check if req.UserID is present in userIDs
	for _, id := range userIDs {
		if id == req.UserID {
			found = true
		}
	}

	if !found {
		// Then add req.UserID to userIDs
		userIDs = append(userIDs, req.UserID)
		if len(userIDs) > int(event.MaxTeamSize) {
			return dto.Response{Code: http.StatusBadRequest, Message: "Team Members cannot be more than max team size"}
		}
	}

	isTeamRegistered := impl.eventRepository.AreUsersInTeam(req.EventID, userIDs)
	if isTeamRegistered {
		return dto.Response{Code: http.StatusBadRequest, Message: "Team Members have already registered for the event"}
	}

	// Register the ids of the team members to the event
	for _, id := range userIDs {
		err = impl.eventRepository.Register(models.EventRegistration{
			EventID: req.EventID,
			UserID:  id,
		})

		if err != nil {
			return dto.Response{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
		}
	}

	err = impl.eventRepository.AddTeam(req.EventID, userIDs, req.TeamName, req.UserID)
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

func (impl *eventServiceImpl) Status(eventID uint, userID uint) dto.Response {
	event, err := impl.eventRepository.FindEventByID(eventID)

	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	if event == nil {
		return dto.Response{Code: http.StatusNotFound, Message: "Event Not Found"}
	}

	isUserRegistered := impl.eventRepository.CheckUserRegistered(models.EventRegistration{
		EventID: eventID,
		UserID:  userID,
	})

	if !isUserRegistered {
		return dto.Response{Code: http.StatusOK, Message: dto.EventStatusResponse{
			IsRegistered: false,
			IsTeam:       false,
			TeamID:       0,
			TeamMembers:  nil,
		}}
	}

	if isUserRegistered && !event.IsTeam {
		return dto.Response{Code: http.StatusOK, Message: dto.EventStatusResponse{
			IsRegistered: true,
			IsTeam:       false,
			TeamID:       0,
			TeamMembers:  nil,
		}}
	}

	teamID, err := impl.eventRepository.GetTeamID(eventID, userID)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	if teamID == nil {
		return dto.Response{Code: http.StatusOK, Message: dto.EventStatusResponse{
			IsRegistered: false,
			IsTeam:       false,
			TeamID:       0,
			TeamMembers:  nil,
		}}
	}

	teamMembers, err := impl.eventRepository.GetTeamMembers(*teamID)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	userNames := make([]string, len(teamMembers))

	for i, userID := range teamMembers {
		user, err := impl.userRepository.FindByID(userID)
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}

		userNames[i] = user.Name
	}

	return dto.Response{Code: http.StatusOK, Message: dto.EventStatusResponse{
		IsRegistered: true,
		IsTeam:       true,
		TeamID:       *teamID,
		TeamMembers:  userNames,
	}}
}
