package impl

import (
	"net/http"
	"strings"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type tShirtsServiceImpl struct {
	tShirtsRepository repository.TShirtsRepository
	userRepository    repository.UserRepository
}

func NewTShirtsServiceImpl(
	tShirtsRepository repository.TShirtsRepository,
	userRepository repository.UserRepository) service.TShirtsService {
	return &tShirtsServiceImpl{
		tShirtsRepository: tShirtsRepository,
		userRepository:    userRepository,
	}
}

func (impl *tShirtsServiceImpl) UpdateSize(req dto.TShirtsUpdateDTO) dto.Response {

	userDetails, err := impl.userRepository.FindByID(req.UserID)
	if err != nil || !userDetails.IsDauth {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Invalid User"}
	}

	err = impl.tShirtsRepository.UpdateSize(req.UserID, req.Size, strings.Split(userDetails.Email, "@")[0])
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "There seems to be an issue updating the t-shirt size"}
	}

	return dto.Response{Code: http.StatusOK, Message: "T-Shirt Size has been updated successfully!"}
}
