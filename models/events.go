package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID        int    `gorm:"not null;primarykey;autoIncrement"`
	EventName string `gorm:"size:255;not null;index"`
}
