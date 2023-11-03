package repository

import (
	"github.com/delta/FestAPI/models"
)

type CollegeRepository interface {
	Insert(college models.College) error
	Delete(college models.College) error
	FindByName(collegeName string) (*models.College, error)
	Exists(collegeName string) error
}
