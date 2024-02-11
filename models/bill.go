package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ID     uint `gorm:"not null;primarykey;autoIncrement"`
	UserID uint
	Email  string    `gorm:"not null"`
	User   User      `gorm:"foreignkey:ID;references:UserID"`
	Time   time.Time `gorm:"not null"`
	Mode   string    `gorm:"not null"`
	Amount uint      `gorm:"not null"`
	RefID  string
	PaidTo AdminRole `gorm:"not null"`
}
