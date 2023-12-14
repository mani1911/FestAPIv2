package impl

import (
	"net/http"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
)

type adminServiceImpl struct {
	adminRepository repository.AdminRepository
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository) service.AdminService {
	return &adminServiceImpl{
		adminRepository: adminRepository,
	}
}

func (impl *adminServiceImpl) Login(req dto.AuthAdminRequest) dto.Response {

	// Initialize Logger
	log := utils.GetServiceLogger("AdminService Login")

	// Checking if both email and password are present
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Username / Password cannot be empty"}
	}

	// Fetching admin details
	adminDetails, err := impl.adminRepository.FindByName(req.Username)
	// If admin doesn't exist
	if adminDetails == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Username or Password is Incorrect"}
	} else if err != nil {
		log.Error("Error fetching Admin Details. Error : ", err)
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error in searching for admins"}
	}

	// Comparing passwords
	err = utils.ComapareHashPassword(adminDetails.Password, req.Password)
	if err != nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Username or Password is Incorrect"}
	}

	// Creating JWT for the user along with role
	jwtToken, err := utils.GenerateToken(adminDetails.ID, true, adminDetails.Role)
	if err != nil {
		log.Error("Error Generating Token. Error : ", err)
		return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
	}

	return dto.Response{Code: http.StatusOK, Message: jwtToken}
}
