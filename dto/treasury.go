package dto

import (
	"github.com/delta/FestAPI/models"
)

type AddBillRequest struct {
	UserID uint             `json:"user_id"`
	Time   string           `json:"time"`
	Mode   string           `json:"mode"`
	Amount uint             `json:"amount"`
	RefID  string           `json:"ref_id"`
	PaidTo models.AdminRole `json:"paid_to"`
	Type   string           `json:"type"`
}

type TownScriptRequest struct {
	UserEmailID      string `json:"userEmailId"`
	UserName         string `json:"userName"`
	Currency         string `json:"currency"`
	TicketName       string `json:"ticketName"`
	EventName        string `json:"eventName"`
	EventCode        string `json:"eventCode"`
	TicketPrice      uint   `json:"ticketPrice"`
	DiscountCode     string `json:"discountCode"`
	DiscountAmount   uint   `json:"discountAmount"`
	CustomQuestion1  string `json:"customQuestion1"`
	CustomQuestion20 string `json:"customQuestion20"`
	AnswerList       []struct {
		UniqueQuestionID uint   `json:"uniqueQuestionId"`
		Question         string `json:"question"`
		Answer           string `json:"answer"`
	} `json:"answerList"`
	UniqueOrderID         string `json:"uniqueOrderId"`
	RegistrationTimestamp string `json:"registrationTimestamp"`
}
