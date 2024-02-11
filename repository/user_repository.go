package repository

import "github.com/delta/FestAPI/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) error
	FindByCollegeID(id uint) (*models.College, error)
	FindTShirtSize(id uint) (*models.TShirts, error)
	SetDauth(user *models.User) error
}
