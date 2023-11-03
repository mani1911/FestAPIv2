package impl

import (
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewAdminRepositoryImpl(DB *gorm.DB) repository.AdminRepository {
	return &adminRepositoryImpl{DB: DB}
}

type adminRepositoryImpl struct {
	*gorm.DB
}

func (repository *adminRepositoryImpl) FindByName(name string) (*models.Admin, error) {
	var adminDetails models.Admin
	// Checking if admin exists in the database
	if err := repository.DB.Where("Username = ? ", name).First(&adminDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &adminDetails, nil
}
