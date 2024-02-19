package impl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"gorm.io/gorm"
)

func NewTreasuryRepositoryImpl(DB *gorm.DB) repository.TreasuryRepository {
	return &treasuryRepositoryImpl{DB: DB}
}

type treasuryRepositoryImpl struct {
	*gorm.DB
}

func (repository *treasuryRepositoryImpl) GetBillByEmailAndPaidTo(UserEmail string, PaidTo string) *models.Bill {
	var bill models.Bill
	err := repository.DB.Model(&models.Bill{}).Where("paid_to = ? ", PaidTo).Where("email = ? ", UserEmail).First(&bill).Error
	if err != nil {
		return nil
	}
	return &bill
}

func (repository *treasuryRepositoryImpl) GetBillByUserIDAndPaidTo(userID uint, PaidTo string) (*models.Bill, error) {
	var bill models.Bill
	err := repository.DB.Model(&models.Bill{}).Where("paid_to = ? ", PaidTo).Where("user_id = ? ", userID).First(&bill).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &bill, nil
}

func (repository *treasuryRepositoryImpl) UpdateBillWithUserID(UserEmail string, UserID uint) error {
	err := repository.DB.Model(&models.Bill{}).Where("email = ? ", UserEmail).Update("user_id", UserID).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *treasuryRepositoryImpl) AddBill(req *dto.AddBillRequest) error {
	var user models.User
	if err := repository.DB.Where("ID = ? ", req.UserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return nil
	}

	bill := models.Bill{
		UserID: req.UserID,
		Email:  user.Email,
		Time:   req.Time,
		Mode:   req.Mode,
		Amount: req.Amount,
		RefID:  req.RefID,
		PaidTo: req.PaidTo,
	}

	if err := repository.DB.Model(&models.Bill{}).Create(&bill).Error; err != nil {
		return errors.New("Error adding visitor")
	}

	switch strings.ToLower(strings.Trim(req.Type, " ")) {
	case "checkin":
		repository.DB.Model(&models.Visitor{}).Where("user_id = ?", req.UserID).Update("check_in_bill_id", bill.ID)
	case "checkout":
		repository.DB.Model(&models.Visitor{}).Where("user_id = ?", req.UserID).Update("check_out_bill_id", bill.ID)
	case "eventpass":
		repository.DB.Model(&models.Visitor{}).Where("user_id = ?", req.UserID).Update("event_pass_bill_id", bill.ID)
	case "discount":
		repository.DB.Model(&models.Visitor{}).Where("user_id = ?", req.UserID).Update("discount_bill_id", bill.ID)
	case "fine":
		repository.DB.Model(&models.Visitor{}).Where("user_id = ?", req.UserID).Update("fine_bill_id", bill.ID)
	}

	return nil
}

func (repository *treasuryRepositoryImpl) Townscript(req *dto.TownScriptRequest) error {

	userEmail := req.UserEmailID
	Mode := "Online"

	Amount, err := strconv.ParseFloat(fmt.Sprint(req.TicketPrice), 32)
	if err != nil {
		Amount = 0
	}

	RefID := req.RegistrationID
	var PaidTo models.AdminRole = models.AdminRole("townScript")
	var Days uint
	startDate := ""

	if req.EventName == "Hospitality - 1 Day" {
		Days = 1
	} else if req.EventName == "Hospitality - 2 Days" {
		Days = 2
	} else if req.EventName == "Hospitality - 3 Days" {
		Days = 3
	} else if req.EventName == "Hospitality - 4 Days" {
		Days = 4
	}

	for _, answer := range req.AnswerList {
		if answer.Question == "Start Date" {
			startDate = answer.Answer
		}
	}

	bill := models.Bill{
		Email:  userEmail,
		Mode:   Mode,
		Amount: float32(Amount),
		RefID:  fmt.Sprint(RefID),
		PaidTo: PaidTo,
		Time:   time.Now(),
	}

	if err := repository.DB.Model(&models.Bill{}).Create(&bill).Error; err != nil {
		return errors.New("Error saving payment bill")
	}

	var user models.User

	err1 := repository.DB.Where(" Email = ? ", userEmail).First(&user).Error
	var userID uint

	if err1 != nil && err1 != gorm.ErrRecordNotFound {
		return err1
	}

	if err1 == nil {
		userID = user.ID
	}

	if err := repository.DB.Model(&models.RoomReg{}).Create(&models.RoomReg{
		UserID:     userID,
		Email:      userEmail,
		NoOfDays:   Days,
		StartDate:  startDate,
		EventCode:  req.EventCode,
		CheckedOut: false,
	}).Error; err != nil {
		return errors.New("Error registering room")
	}
	return nil
}

func (repository *treasuryRepositoryImpl) AddBillByModel(bill *models.Bill) error {
	if err := repository.DB.Model(&models.Bill{}).Create(&bill).Error; err != nil {
		return errors.New("Error saving payment bill")
	}

	return nil
}
