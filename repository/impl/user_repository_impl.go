package impl

import (
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (repository *userRepositoryImpl) CreateUser(user *models.User) error {

	// Storing new user in the database
	if err := repository.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var userDetail models.User

	// Find User by Email
	if err := repository.DB.Where("Email = ? ", email).First(&userDetail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &userDetail, nil
}

func (repository *userRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var userDetail models.User

	// Find User by ID
	if err := repository.DB.Where("ID = ? ", id).First(&userDetail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &userDetail, nil
}

func (repository *userRepositoryImpl) FindByCollegeID(id uint) (*models.College, error) {
	var collegeDetail models.College

	// Find College by ID
	if err := repository.DB.Where("ID = ? ", id).First(&collegeDetail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &collegeDetail, nil
}

func (repository *userRepositoryImpl) Update(userDetails *models.User) error {

	// Update User
	if err := repository.DB.Save(&userDetails).Error; err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryImpl) CreateForgotPasswordUser(user *models.ForgotPasswordUser) error {

	// Create Forgot Password User
	if err := repository.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryImpl) FindForgotPasswordUserByEmail(email string) (*models.ForgotPasswordUser, error) {
	var userDetail models.ForgotPasswordUser

	// Find Forgot User by Email
	if err := repository.DB.Where("Email = ? ", email).Last(&userDetail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &userDetail, nil
}
