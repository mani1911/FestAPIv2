package service

import "github.com/delta/FestAPI/dto"

type TShirtsService interface {
	UpdateSize(dto.TShirtsUpdateDTO) dto.Response
}
