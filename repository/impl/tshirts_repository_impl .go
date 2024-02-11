package impl

import (
	"errors"

	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewTShirtsRepositoryImpl(DB *gorm.DB) repository.TShirtsRepository {
	return &tshirtsRepositoryImpl{DB: DB}
}

type tshirtsRepositoryImpl struct {
	*gorm.DB
}

func (repository *tshirtsRepositoryImpl) UpdateSize(userID uint, size string, rollNo string) error {
	var tshirtsDetails models.TShirts

	if err := repository.DB.Where("user_id = ?", userID).First(&tshirtsDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tshirtsDetails.UserID = userID
			tshirtsDetails.Size = size
			tshirtsDetails.RollNo = rollNo
			if errSave := repository.DB.Save(&tshirtsDetails).Error; errSave != nil {
				return errors.New("Cannot update t-shirt details")
			}
		} else {
			return errors.New("Cannot update t-shirt details")
		}
	} else {
		tshirtsDetails.Size = size
		if errSave := repository.DB.Save(&tshirtsDetails).Where("user_id = ?", tshirtsDetails.UserID).Error; errSave != nil {
			return errors.New("Cannot update t-shirt details")
		}
	}
	return nil
}
