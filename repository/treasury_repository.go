package repository

import (
	"github.com/delta/FestAPI/dto"
)

type TreasuryRepository interface {
	AddBill(*dto.AddBillRequest) error
	Townscript(*dto.TownScriptRequest) error
}
