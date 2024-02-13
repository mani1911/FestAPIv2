package impl

import (
	"fmt"
	"net/http"
	"time"

	dto "github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
)

type treasuryServiceImpl struct {
	treasuryRepository repository.TreasuryRepository
	adminRepository    repository.AdminRepository
	collegeRepository  repository.CollegeRepository
	userRepository     repository.UserRepository
}

func NewTreasuryServiceImpl(
	treasuryRepository repository.TreasuryRepository,
	adminRepository repository.AdminRepository,
	userRepository repository.UserRepository,
	collegeRepository repository.CollegeRepository) service.TreasuryService {
	return &treasuryServiceImpl{
		collegeRepository:  collegeRepository,
		adminRepository:    adminRepository,
		userRepository:     userRepository,
		treasuryRepository: treasuryRepository,
	}
}

func (impl *treasuryServiceImpl) AddBill(req dto.AddBillRequest) dto.Response {
	if req.Amount == 0 || req.Mode == "" || req.PaidTo == "" || req.UserID == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
	} else if req.Type == "" {
		return dto.Response{Code: http.StatusBadRequest, Message: "Purpose of payment not found"}
	} else if req.Time == "" {
		req.Time = time.Now().Format("2006-01-02 15:04:05")
	}

	if err := impl.treasuryRepository.AddBill(&req); err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to add bill"}
	}
	return dto.Response{Code: http.StatusOK, Message: "Bill added!"}
}

func (impl *treasuryServiceImpl) Townscript(req dto.TownScriptRequest) dto.Response {
	log := utils.GetControllerLogger("TreasuryController TownScript")
	if req.UserEmailID == "" || req.Currency == "" || req.EventName == "" || req.EventCode == "" || req.TicketPrice == "" {
		log.Fatal("Malformed Request", fmt.Sprint(req))
		return dto.Response{Code: http.StatusBadRequest, Message: "Malformed Request"}
	}

	if err := impl.treasuryRepository.Townscript(&req); err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to register payment"}
	}
	log.Println("Payment made!", req.UserEmailID, req.TicketPrice, req.EventName, req.EventCode, req.RegistrationID, req.RegistrationTimestamp)
	return dto.Response{Code: http.StatusOK, Message: "Payment made!"}
}
