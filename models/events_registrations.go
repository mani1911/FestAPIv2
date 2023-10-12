package models

import (
	"gorm.io/gorm"
)

type EventRegistration struct {
	gorm.Model
	EventID uint `gorm:"index;not null"`
	UserID  uint `gorm:"index;not null"`
}
