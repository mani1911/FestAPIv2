package repository

import "github.com/delta/FestAPI/models"

type AdminRepository interface {
	FindByName(username string) (*models.Admin, error)
}
