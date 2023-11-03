package models

import (
	"gorm.io/gorm"
)

type EventAbstractDetails struct {
	gorm.Model
	ID              uint `gorm:"primaryKey;autoIncrement;not null"`
	EventID         uint
	Event           Event  `gorm:"foreignKey:EventID"`
	ForwardEmail    string `gorm:"size:255;not null"`
	MaxParticipants uint   `gorm:"not null"`
}
