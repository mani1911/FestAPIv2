package service

import (
	"github.com/delta/FestAPI/dto"
)

type UserService interface {
	DAuthLogin(dto.AuthUserRequest) dto.Response
	Login(dto.AuthUserLoginRequest) dto.Response
	Register(dto.AuthUserRegisterRequest) dto.Response
	Update(dto.AuthUserUpdateRequest, uint) dto.Response
}
