package models

import "gorm.io/gorm"

type TShirts struct {
	gorm.Model

	ID             uint   `gorm:"primaryKey;autoIncrement;not null"`
	UserID         uint   `gorm:"index;not null"`
	RollNo         string `gorm:"size:255;not null"`
	Size           string `gorm:"size:255;not null"`
	Code           string `gorm:"size:255;not null"`
	ScreenshotLink string `gorm:"text;not null"`
}
