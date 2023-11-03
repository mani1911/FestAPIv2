package models

import (
	"gorm.io/gorm"
)

type InformalsDetails struct {
	gorm.Model
	ID   uint   `gorm:"not null;primarykey;autoIncrement"`
	Name string `gorm:"size:255;not null;unique;uniqueIndex"`
}
