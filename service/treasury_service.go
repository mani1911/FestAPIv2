package service

import "github.com/delta/FestAPI/dto"

type TreasuryService interface {
	AddBill(dto.AddBillRequest) dto.Response
	Townscript(dto.TownScriptRequest) dto.Response
}
