package service

import "github.com/delta/FestAPI/dto"

type AdminService interface {
	Login(req dto.AuthAdminRequest) dto.Response
}
