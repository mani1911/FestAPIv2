package service

import "github.com/delta/FestAPI/dto"

type HospiService interface {
	GetHostels() dto.Response
	AddUpdateHostel(dto.AddUpdateHostelRequest) dto.Response
	GetRooms() dto.Response
	AddUpdateRoom(dto.AddUpdateRoomRequest) dto.Response
	DeleteRoom(dto.DeleteRoomRequest) dto.Response
	CheckIn(dto.CheckInRequest) dto.Response
}
