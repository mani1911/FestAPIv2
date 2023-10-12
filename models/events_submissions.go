package models

import (
	"time"

	"gorm.io/gorm"
)

type EventSubmission struct {
	gorm.Model
	EventID         uint      `gorm:"primaryKey;not null"`
	ForwardEmail    string    `gorm:"size:255;not null"`
	MaxParticipants int       `gorm:"not null"`
	NoOfFiles       int       `gorm:"not null"`
	Deadline        time.Time `gorm:"not null"`
}
