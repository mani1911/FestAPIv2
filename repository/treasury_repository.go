package repository

import (
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
)

type TreasuryRepository interface {
	AddBill(*dto.AddBillRequest) error
	Townscript(*dto.TownScriptRequest) error
	GetBillByEmailAndPaidTo(string, string) *models.Bill
	UpdateBillWithUserID(string, uint) error
}
