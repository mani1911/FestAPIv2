package impl

import (
	"net/http"
	"strconv"
	"time"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
)

type prServiceImpl struct {
	userRepository     repository.UserRepository
	treasuryRepository repository.TreasuryRepository
	hospiRepository    repository.HospiRepository
}

func NewPRServiceImpl(
	userRepository repository.UserRepository, treasuryRepository repository.TreasuryRepository, hospiRepository repository.HospiRepository) service.PRService {
	return &prServiceImpl{
		userRepository:     userRepository,
		treasuryRepository: treasuryRepository,
		hospiRepository:    hospiRepository,
	}
}

func (impl *prServiceImpl) RegisterStatus(userEmail string) dto.Response {
	user, _ := impl.userRepository.FindByEmail(userEmail)
	if user == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found. Ask user to register on the Pragyan Site with the same email as Townscript Payments"}
	}
	roomReg := impl.hospiRepository.FindRoomRegByID(user.ID)

	return dto.Response{Code: http.StatusAccepted, Message: dto.RegisterStatusResponse{
		User:    *user,
		RoomReg: *roomReg,
	}}
}

func (impl *prServiceImpl) Register(userID uint, registerAmount string) dto.Response {
	logger := utils.GetServiceLogger("PRService Register")
	user, _ := impl.userRepository.FindByID(userID)
	if user == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found. Ask user to register on the Pragyan Site with the same email as Townscript Payments"}
	}
	bill := impl.treasuryRepository.GetBillByEmailAndPaidTo(user.Email, "townScript")
	if bill != nil {
		err := impl.treasuryRepository.UpdateBillWithUserID(user.Email, user.ID)
		if err != nil {
			logger.Warn("Couldnt Update Bill for User : ", user.Email)
			return dto.Response{Code: http.StatusInternalServerError, Message: "Online Bill Email and Pragyan Email don't match"}
		}
		err = impl.hospiRepository.UpdateRoomRegWithUserID(user.Email, user.ID)
		if err != nil {
			logger.Warn("Couldnt Update RoomReg for User : ", user.Email)
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}
	}
	parsedAmount, err := strconv.ParseInt(registerAmount, 10, 64)
	if err != nil {
		logger.Warn("Amount not a Number for user : ", user.Email)
		return dto.Response{Code: http.StatusBadRequest, Message: "Bad Request"}
	}
	err = impl.treasuryRepository.AddBill(&dto.AddBillRequest{
		UserID: userID,
		Time:   time.Now(),
		Amount: uint(parsedAmount),
		RefID:  "",
		PaidTo: models.PR,
		Type:   "eventPass",
	})
	if err != nil {
		logger.Warn("Error Creating Bill for User : ", user.Email)
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	return dto.Response{Code: http.StatusAccepted, Message: "User Registered Successfully"}
}
