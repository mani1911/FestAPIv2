package models

import (
	"gorm.io/gorm"
)

type EventTeam struct {
	gorm.Model
	TeamID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	EventID      uint   `gorm:"not null"`
	TeamName     string `gorm:"size:255;not null"`
	TeamLeaderID uint   `gorm:"not null"`
}
