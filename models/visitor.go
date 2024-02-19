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
	CheckInTime     time.Time // Hospi Checkin Time
	CheckOutTime    time.Time // Hospi Checkout Time
	CheckInBillID   uint
	CheckOutBillID  uint
	FineBillID      uint
	DiscountBillID  uint
	EventPassBillID uint
	RoomRegID       uint
	RoomReg         RoomReg `gorm:"foreignkey:ID;references:RoomRegID"`
}
