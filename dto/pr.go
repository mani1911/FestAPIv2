package dto

import "github.com/delta/FestAPI/models"

type RegisterRequest struct {
	UserID    uint   `json:"user_id"`
	RegAmount string `json:"reg_amount"`
}

type RegisterStatusResponse struct {
	User            models.User    `json:"user"`
	RoomReg         models.RoomReg `json:"roomReg"`
	TownScriptBills []models.Bill  `json:"townscriptBills"`
}
