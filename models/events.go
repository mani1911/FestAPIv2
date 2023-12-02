package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uint   `gorm:"not null;primarykey;autoIncrement"`
	EventName   string `gorm:"size:255;not null;index"`
	IsTeam      bool   `gorm:"not null;default:false"`
	MaxTeamSize uint   `gorm:"not null;default:1"`
}
