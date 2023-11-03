package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID        uint   `gorm:"not null;primarykey;autoIncrement"`
	EventName string `gorm:"size:255;not null;index"`
}
