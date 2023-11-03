package service

import (
	"github.com/delta/FestAPI/dto"
)

type EventService interface {
	Register(dto.EventRegistrationDTO) dto.Response
	AbstractDetails(dto.AbstractDetailsRequest) dto.Response
	UserEventDetails(userID uint) dto.Response
}
