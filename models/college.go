package models

import "gorm.io/gorm"

type College struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"size:255;not null"`
}
