package impl

import (
	"errors"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewHospiRepositoryImpl(DB *gorm.DB) repository.HospiRepository {
	return &hospiRepositoryImpl{DB: DB}
}

type hospiRepositoryImpl struct {
	*gorm.DB
}

func (repository *hospiRepositoryImpl) GetHostels() ([]*dto.GetHostelsResponse, error) {
	var res []*dto.GetHostelsResponse

	if err := repository.DB.Model(&models.Hostel{}).Select("id, name, gender").Find(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}

func (repository *hospiRepositoryImpl) AddHostel(hostel *models.Hostel) error {
	if err := repository.DB.Create(&hostel).Error; err != nil {
		return errors.New("Error creating hostel")
	}
	return nil
}

func (repository *hospiRepositoryImpl) UpdateHostel(hostel *models.Hostel) error {
	if err := repository.DB.Save(&hostel).Where("ID = ?", hostel.ID).Error; err != nil {
		return errors.New("Cannot update hostel details")
	}
	return nil
}

func (repository *hospiRepositoryImpl) FindHostelByID(id uint) (*models.Hostel, error) {
	var res models.Hostel

	if err := repository.DB.Where("ID = ? ", id).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (repository *hospiRepositoryImpl) GetRooms() ([]*dto.GetRoomsResponse, error) {
	var res []*dto.GetRoomsResponse

	if err := repository.DB.
		Model(&models.Hostel{}).
		Select("rooms.id as room_id, rooms.name as room, hostels.id as hostel_id, hostels.name as hostel, hostels.gender as gender").
		Joins("RIGHT JOIN rooms ON hostels.id = rooms.hostel_id").
		Find(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}

func (repository *hospiRepositoryImpl) AddRoom(room *models.Room) error {
	if err := repository.DB.Create(&room).Error; err != nil {
		return errors.New("Error creating room")
	}
	return nil
}

func (repository *hospiRepositoryImpl) UpdateRoom(room *models.Room) error {
	if err := repository.DB.Save(&room).Where("ID = ?", room.ID).Error; err != nil {
		return errors.New("Error updating room details")
	}
	return nil
}

func (repository *hospiRepositoryImpl) DeleteRoom(id uint) error {
	// hard delete
	if err := repository.DB.Unscoped().Delete(&models.Room{ID: id}).Error; err != nil {
		return errors.New("Error deleting the room")
	}
	return nil
}

func (repository *hospiRepositoryImpl) FindRoomByID(id uint) (*models.Room, error) {
	var res models.Room

	if err := repository.DB.Where("ID = ? ", id).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}
