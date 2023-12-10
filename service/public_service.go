package service

import "github.com/delta/FestAPI/dto"

type PublicService interface {
	AllColleges() dto.Response
}
