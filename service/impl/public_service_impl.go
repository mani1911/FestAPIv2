package impl

import (
	"net/http"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type publicServiceImpl struct {
	collegeRepository repository.CollegeRepository
}

func NewPublicServiceImpl(collegeRepository repository.CollegeRepository) service.PublicService {
	return &publicServiceImpl{collegeRepository: collegeRepository}
}

func (impl *publicServiceImpl) AllColleges() dto.Response {
	var colleges []dto.CollegeResponse
	collegeDetails, err := impl.collegeRepository.GetAllColleges()
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error fetching colleges"}
	}

	for _, college := range collegeDetails {
		newRespObject := dto.CollegeResponse{
			ID:   college.ID,
			Name: college.Name,
		}
		colleges = append(colleges, newRespObject)
	}

	return dto.Response{Code: http.StatusOK, Message: colleges}
}
