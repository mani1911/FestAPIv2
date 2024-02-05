package impl

import (
	"net/http"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
)

type adminServiceImpl struct {
	adminRepository repository.AdminRepository
	userRepository  repository.UserRepository
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository, userRepository repository.UserRepository) service.AdminService {
	return &adminServiceImpl{
		adminRepository: adminRepository,
		userRepository:  userRepository,
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
		log.Error("Error fetching Admin Details. Error : ", err.Error())
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
		log.Error("Error Generating Token. Error : ", err.Error())
		return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
	}

	return dto.Response{Code: http.StatusOK, Message: jwtToken}
}

func (impl *adminServiceImpl) VerifyUser(req dto.UserInfoRequest) dto.Response {
	if len(req.InfoType) == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Info Type cannot be empty"}
	}

	var userDetails *models.User
	var err error

	switch req.InfoType {
	case "jwt":
		claims, parseErr := utils.ParseToken(req.Info)
		if parseErr != nil {
			return dto.Response{Code: http.StatusBadRequest, Message: err.Error()}
		}
		userDetails, err = impl.userRepository.FindByEmail(claims.UserEmail)
	case "email":
		userDetails, err = impl.userRepository.FindByEmail(req.Info)
	}

	if err != nil {
		return dto.Response{Code: http.StatusBadRequest, Message: err.Error()}
	}

	college, _ := impl.userRepository.FindByCollegeID(userDetails.CollegeID)
	userDetails.College = *college

	return dto.Response{Code: http.StatusOK, Message: userDetails}
}
