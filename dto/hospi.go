package dto

import "github.com/delta/FestAPI/models"

type GetHostelsResponse struct {
	ID     uint
	Name   string
	Gender models.Gender
}

type AddUpdateHostelRequest struct {
	ID     uint          `json:"id"`
	Name   string        `json:"name"`
	Gender models.Gender `json:"gender"`
}

type GetRoomsResponse struct {
	RoomID   uint
	Room     string
	Hostel   string
	HostelID uint
	Gender   models.Gender
}

type AddUpdateRoomRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	HostelID uint   `json:"hostel_id"`
	Capacity uint   `json:"capacity"`
}

type DeleteRoomRequest struct {
	ID uint `json:"id"`
}
