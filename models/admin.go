package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey;autoIncrement;not null"`
	Username string    `gorm:"size:255;not null;unique;uniqueIndex"`
	Password []byte    `gorm:"size:255;not null"`
	Role     AdminRole `gorm:"size:255;not null;type:admin_roles"`
}
