package dto

import (
	"time"

	"github.com/delta/FestAPI/models"
)

type AddBillRequest struct {
	UserID uint             `json:"user_id"`
	Time   time.Time        `json:"time"`
	Mode   string           `json:"mode"`
	Amount uint             `json:"amount"`
	RefID  string           `json:"ref_id"`
	PaidTo models.AdminRole `json:"paid_to"`
	Type   string           `json:"type"`
}

type TownScriptRequest struct {
	AnswerList            []CustomQuestion `json:"answerList"`
	RegistrationTimestamp string           `json:"registrationTimestamp"`
	UserEmailID           string           `json:"userEmailId"`
	RegistrationID        int              `json:"registrationId"`
	EventName             string           `json:"eventName"`
	EventCode             string           `json:"eventCode"`
	TicketPrice           float32          `json:"ticketPrice"`
}

type CustomQuestion struct {
	Question         string `json:"question"`
	Answer           string `json:"answer"`
	UniqueQuestionID int    `json:"uniqueQuestionId"`
}
