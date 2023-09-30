package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint    `gorm:"primaryKey;autoIncrement;not null"`
	Name         string  `gorm:"size:255;not null"`
	FullName     string  `gorm:"size:255; not null"`
	CollegeID    uint    `gorm:"foreignKey:CollegeID"`
	College      College `gorm:"foreignKey:CollegeID"`
	OtherCollege string  `gorm:"size:255"`
	Email        string  `gorm:"size:255;unique;not null;uniqueIndex"`
	Gender       Gender  `gorm:"size:10;not null;type:gender"`
	Country      string  `gorm:"size:255;not null"`
	State        string  `gorm:"size:255;not null"`
	City         string  `gorm:"size:255;not null"`
	Address      string  `gorm:"size:255;not null"`
	Pincode      string  `gorm:"size:255;not null"`
	Phone        string  `gorm:"size:255;not null"`
	Password     []byte  `gorm:"size:255;not null"`
	Sponsor      string  `gorm:"size:255;not null"`
	VoucherName  string  `gorm:"size:255;not null"`
	ReferralCode string  `gorm:"size:255;not null"`
	Degree       string  `gorm:"size:255;not null"`
	Year         string  `gorm:"size:255;not null"`
	Nationality  string  `gorm:"size:255;not null"`
}
