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

func (repository *hospiRepositoryImpl) GetRooms(hostelID int, floor int, isFilled int) ([]*dto.GetRoomsResponse, error) {
	var res []*dto.GetRoomsResponse

	var query = repository.DB.
		Model(&models.Hostel{}).
		Select("rooms.id as room_id, rooms.name as room, hostels.id as hostel_id, hostels.name as hostel, hostels.gender as gender, rooms.capacity as capacity, rooms.occupied as occupied, rooms.floor as floor").
		Joins("RIGHT JOIN rooms ON hostels.id = rooms.hostel_id")

	if hostelID > 0 {
		query.Where("hostel_id = ?", hostelID)
	}

	if floor >= 0 {
		query.Where("floor = ?", floor)
	}

	if isFilled == 0 {
		query.Where("occupied < capacity")
	}

	if err := query.Find(&res).Error; err != nil {
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

func (repository *hospiRepositoryImpl) AddRoomReg(req *models.RoomReg) error {
	if err := repository.DB.Model(&models.RoomReg{}).Create(&req).Error; err != nil {
		return errors.New("Error registering room")
	}

	return nil
}

func (repository *hospiRepositoryImpl) UpdateRoomRegWithUserID(userEmail string, userID uint) error {
	if err := repository.DB.Model(&models.RoomReg{}).Where("email = ? ", userEmail).Update("user_id", userID).Error; err != nil {
		return errors.New("Error updating room registration")
	}
	return nil
}

func (repository *hospiRepositoryImpl) UpdateRoomRegWithRoomID(userEmail string, roomID uint) error {
	if err := repository.DB.Model(&models.RoomReg{}).Where("email = ? ", userEmail).Update("room_id", roomID).Error; err != nil {
		return errors.New("Error updating room registration")
	}
	return nil
}

func (repository *hospiRepositoryImpl) AddVisitor(req *models.Visitor) error {
	if err := repository.DB.Model(&models.Visitor{}).Create(&req).Error; err != nil {
		return errors.New("Error adding visitor")
	}
	return nil
}

func (repository *hospiRepositoryImpl) UpdateVisitor(req *models.Visitor) error {
	if err := repository.DB.Save(&req).Error; err != nil {
		return errors.New("Error updating visitor")
	}

	return nil
}

func (repository *hospiRepositoryImpl) FindRoomRegByUserID(userID uint) (*models.RoomReg, error) {
	var roomReg models.RoomReg

	if err := repository.DB.Model(&models.RoomReg{}).Where("user_id = ?", userID).First(&roomReg).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &roomReg, nil
}

func (repository *hospiRepositoryImpl) FindRoomRegByUserEmail(userEmail string) (*models.RoomReg, error) {
	var roomReg models.RoomReg

	if err := repository.DB.Model(&models.RoomReg{}).Where("email = ?", userEmail).First(&roomReg).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &roomReg, nil
}

func (repository *hospiRepositoryImpl) FindVisitorByUserID(userID uint) (*models.Visitor, error) {
	var visitor models.Visitor

	if err := repository.DB.Model(&models.Visitor{}).Where("user_id = ?", userID).First(&visitor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &visitor, nil
}
