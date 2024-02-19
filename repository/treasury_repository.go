package repository

import (
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
)

type TreasuryRepository interface {
	AddBill(*dto.AddBillRequest) error
	Townscript(*dto.TownScriptRequest) error
	GetBillByEmailAndPaidTo(string, string) *models.Bill
	GetBillByUserIDAndPaidTo(uint, string) (*models.Bill, error)
	UpdateBillWithUserID(string, uint) error
	AddBillByModel(*models.Bill) error
}
