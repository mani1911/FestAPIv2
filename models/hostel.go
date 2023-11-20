package models

import (
	"gorm.io/gorm"
)

type Hostel struct {
	gorm.Model
	ID     uint   `gorm:"not null;primarykey;autoIncrement"`
	Name   string `gorm:"size:255;not null;index"`
	Gender Gender `gorm:"size:10;not null;type:gender"`
}
