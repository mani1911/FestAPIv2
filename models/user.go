package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint    `gorm:"primaryKey;autoIncrement;not null"`
	Name         string  `gorm:"size:255;not null"`
	CollegeID    uint    `gorm:"foreignKey:CollegeID"`
	College      College `gorm:"foreignKey:CollegeID"`
	OtherCollege string  `gorm:"size:255"`
	Email        string  `gorm:"size:255;unique;not null"`
	Gender       Gender  `gorm:"size:10;not null"`
	Country      string  `gorm:"size:255;not null"`
	State        string  `gorm:"size:255;not null"`
	City         string  `gorm:"size:255;not null"`
	Address      string  `gorm:"size:255;not null"`
	Pincode      string  `gorm:"size:255;not null"`
	Phone        string  `gorm:"size:255;not null"`
}
