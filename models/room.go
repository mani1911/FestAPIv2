package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint   `gorm:"not null;primarykey;autoIncrement"`
	Name     string `gorm:"size:255;not null;index"`
	HostelID uint   `gorm:"not null"`
	Capacity uint   `gorm:"not null"`
	Floor    uint   `gorm:"not null"`
	Occupied uint   `gorm:"not null"`
}
