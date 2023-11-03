package impl

import (
	"errors"
	"fmt"

	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewCollegeRepositoryImpl(DB *gorm.DB) repository.CollegeRepository {
	return &collegeRepositoryImpl{DB: DB}
}

type collegeRepositoryImpl struct {
	*gorm.DB
}

func (repository *collegeRepositoryImpl) Insert(college models.College) error {
	err := repository.DB.Create(&college).Error
	if err != nil {
		return fmt.Errorf("Cannot insert college data")
	}
	return nil
}

func (repository *collegeRepositoryImpl) Delete(college models.College) error {
	err := repository.DB.Delete(&college).Error
	if err != nil {
		return fmt.Errorf("Cannot delete college data")
	}
	return nil
}

func (repository *collegeRepositoryImpl) FindByName(collegeName string) (*models.College, error) {
	var collegeDetail models.College
	if err := repository.DB.Where("Name =? ", collegeName).First(&collegeDetail).Error; err != nil {
		return nil, errors.New("Cannot find college")
	}
	return &collegeDetail, nil
}

func (repository *collegeRepositoryImpl) Exists(collegeName string) error {
	var collegeDetail models.College
	if err := repository.DB.Where("Name =? ", collegeName).First(&collegeDetail).Error; err == gorm.ErrRecordNotFound {
		return errors.New("Invalid College Name")
	}
	return nil
}
