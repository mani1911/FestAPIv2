package models

import (
	"gorm.io/gorm"
)

type RoomReg struct {
	gorm.Model
	ID         uint `gorm:"not null;primarykey;autoIncrement"`
	RoomID     uint
	Email      string
	UserID     uint
	EventCode  string `gorm:"not null"`
	NoOfDays   uint   `gorm:"not null"`
	StartDate  string `gorm:"not null"`
	CheckedOut bool
}
