package models

import (
	"gorm.io/gorm"
)

type EventTeamMember struct {
	gorm.Model
	ID     uint `gorm:"not null;primarykey;autoIncrement"`
	TeamID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
}
