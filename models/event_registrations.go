package models

import (
	"gorm.io/gorm"
)

type EventRegistration struct {
	gorm.Model
	ID      uint  `gorm:"primaryKey;autoIncrement;not null"`
	EventID uint  `gorm:"foreignKey:EventID"`
	Event   Event `gorm:"foreignKey:EventID"`
	UserID  uint  `gorm:"index;not null"`
}
