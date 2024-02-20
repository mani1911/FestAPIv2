package dto

import (
	"github.com/delta/FestAPI/models"
)

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
	Capacity uint
	Occupied uint
	Floor    uint
}

type GetRoomRequest struct {
	HostelID int `json:"hostel_id"`
	Floor    int `json:"floor"`
	IsFilled int `json:"is_filled"`
}

type AddUpdateRoomRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	HostelID uint   `json:"hostel_id"`
	Capacity uint   `json:"capacity"`
	Floor    uint   `json:"floor"`
}

type DeleteRoomRequest struct {
	ID uint `json:"id"`
}

type CheckInRequest struct {
	CheckInTime string `json:"time"`
	UserID      uint   `json:"user_id"`
	RoomID      uint   `json:"room_id"`
	NoOfDays    uint   `json:"no_of_days"`
}

type CheckInStatusRequest struct {
	EmailID string `json:"email_id"`
}

type CheckInStatusResponse struct {
	NoOfDays   uint
	StartDate  string
	CheckedOut bool
	RoomID     uint
}

type AllocateRoomRequest struct {
	UserID       uint    `json:"user_id"`
	RoomID       uint    `json:"room_id"`
	NumberOfDays uint    `json:"number_of_days"`
	Amount       float32 `json:"amount"`
}

type CheckOutRequest struct {
	UserID        uint `json:"user_id"`
	Fine          uint `json:"fine"`
	FineReqID     uint `json:"fine_req_id"`
	Discount      uint `json:"discount"`
	DiscountReqID uint `json:"discount_req_id"`
}
