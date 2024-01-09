package models

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPasswordUser struct {
	gorm.Model
	ID             uint      `gorm:"not null;primarykey;autoIncrement"`
	Email          string    `gorm:"size:255;not null;index"`
	Code           string    `gorm:"size:255;not null;index"`
	ExpirationDate time.Time `json:"expiration_date"`
}
