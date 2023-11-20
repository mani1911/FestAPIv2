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
	GetRooms() ([]*dto.GetRoomsResponse, error)
	AddRoom(room *models.Room) error
	UpdateRoom(room *models.Room) error
	DeleteRoom(id uint) error
	FindRoomByID(id uint) (*models.Room, error)
}
