package repository

import (
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
)

type HospiRepository interface {
	GetHostels() ([]*dto.GetHostelsResponse, error)
	AddHostel(*models.Hostel) error
	UpdateHostel(*models.Hostel) error
	FindHostelByID(uint) (*models.Hostel, error)
	GetRooms(int, int, int) ([]*dto.GetRoomsResponse, error)
	AddRoom(*models.Room) error
	UpdateRoom(*models.Room) error
	DeleteRoom(uint) error
	FindRoomByID(uint) (*models.Room, error)
	RoomReg(*models.RoomReg) error
	UpdateRoomRegWithUserID(string, uint) error
	AddVisitor(*models.Visitor) error
	FindRoomRegByID(uint) *models.RoomReg
}
