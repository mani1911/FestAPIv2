package service

import (
	"github.com/delta/FestAPI/dto"
)

type CMSService interface {
	AddEvent(dto.AddEventRequest) dto.Response
}
