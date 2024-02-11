package models

import (
	"time"

	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model
	ID              uint      `gorm:"not null;primarykey;autoIncrement"`
	UserID          uint      `gorm:"not null"`
	User            User      `gorm:"foreignkey:ID;references:UserID"`
	CheckInTime     time.Time `gorm:"not null"`
	CheckOutTime    time.Time
	CheckInBillID   uint
	CheckOutBillID  uint
	FineBillID      uint
	DiscountBillID  uint
	EventPassBillID uint
	RoomID          uint    `gorm:"not null"`
	Room            RoomReg `gorm:"foreignkey:UserID;references:UserID"`
}
