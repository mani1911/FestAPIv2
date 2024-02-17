package service

import "github.com/delta/FestAPI/dto"

type PRService interface {
	Register(uint, string) dto.Response
	RegisterStatus(string) dto.Response
}
