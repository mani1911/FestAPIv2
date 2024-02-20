package impl

import (
	"net/http"
	"time"

	dto "github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type hospiServiceImpl struct {
	hospiRepository    repository.HospiRepository
	adminRepository    repository.AdminRepository
	userRepository     repository.UserRepository
	treasuryRepository repository.TreasuryRepository
}

func NewHospiServiceImpl(
	hospiRepository repository.HospiRepository,
	adminRepository repository.AdminRepository,
	userRepository repository.UserRepository,
	treasuryRepository repository.TreasuryRepository) service.HospiService {
	return &hospiServiceImpl{
		hospiRepository:    hospiRepository,
		adminRepository:    adminRepository,
		userRepository:     userRepository,
		treasuryRepository: treasuryRepository,
	}
}

func (impl *hospiServiceImpl) GetHostels() dto.Response {

	hostels, err := impl.hospiRepository.GetHostels()
	if hostels == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Hostels not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: hostels}
}

func (impl *hospiServiceImpl) AddUpdateHostel(req dto.AddUpdateHostelRequest) dto.Response {
	var hostelDetails models.Hostel
	if req.ID == 0 {
		if len(req.Name) == 0 ||
			len(req.Gender) == 0 {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		}

		hostelDetails.Name = req.Name
		hostelDetails.Gender = req.Gender

		if err := impl.hospiRepository.AddHostel(&hostelDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to create Hostel"}
		}
	} else {
		hostel, err := impl.hospiRepository.FindHostelByID(req.ID)
		if hostel == nil && err == nil {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		} else if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}

		if len(req.Name) == 0 ||
			len(req.Gender) == 0 {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		}

		hostelDetails.ID = req.ID
		hostelDetails.Name = req.Name
		hostelDetails.Gender = req.Gender

		if err := impl.hospiRepository.UpdateHostel(&hostelDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to update Hostel"}
		}
	}
	return dto.Response{Code: http.StatusOK, Message: "Successfully updated the hostel"}
}

func (impl *hospiServiceImpl) GetRooms(req dto.GetRoomRequest) dto.Response {
	res, err := impl.hospiRepository.GetRooms(req.HostelID, req.Floor, req.IsFilled)

	if res == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Rooms not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: res}
}

func (impl *hospiServiceImpl) AddUpdateRoom(req dto.AddUpdateRoomRequest) dto.Response {
	var roomDetails models.Room

	if len(req.Name) == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Room Name"}
	}

	if req.Capacity == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Human require space bro :)"}
	}

	if req.HostelID == 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "How they gonna have room without hostel?"}
	}

	if req.ID != 0 {
		room, err := impl.hospiRepository.FindRoomByID(req.ID)
		if room == nil && err == nil {
			return dto.Response{Code: http.StatusBadRequest, Message: "Room not found.."}
		} else if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}

		if req.Capacity < room.Occupied {
			return dto.Response{Code: http.StatusBadRequest, Message: "Capacity can't reduce capacity if room is already occupied"}
		}

		roomDetails = *room
	}
	roomDetails.Name = req.Name
	roomDetails.Capacity = req.Capacity
	roomDetails.HostelID = req.HostelID
	roomDetails.Floor = req.Floor

	if req.ID == 0 {
		if err := impl.hospiRepository.AddRoom(&roomDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to add room"}
		}
	} else {
		if err := impl.hospiRepository.UpdateRoom(&roomDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to update room"}
		}
	}

	return dto.Response{Code: http.StatusOK, Message: "Successfully updated the room"}
}

func (impl *hospiServiceImpl) DeleteRoom(req dto.DeleteRoomRequest) dto.Response {
	room, err := impl.hospiRepository.FindRoomByID(req.ID)
	if room == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Room ID"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	if err := impl.hospiRepository.DeleteRoom(req.ID); err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: "Successfully deleted the room"}
}

func (impl *hospiServiceImpl) CheckInStatus(req dto.CheckInStatusRequest) dto.Response {
	user, err := impl.userRepository.FindByEmail(req.EmailID)

	if user == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User with this email not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	visitor, err := impl.hospiRepository.FindVisitorByUserID(user.ID)

	if visitor == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Please complete registration process at PR desk."}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	roomReg, err := impl.hospiRepository.FindRoomRegByUserID(visitor.UserID)

	if roomReg == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User has not checked in"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: dto.CheckInStatusResponse{
		NoOfDays:   roomReg.NoOfDays,
		StartDate:  roomReg.StartDate,
		CheckedOut: roomReg.CheckedOut,
		RoomID:     roomReg.RoomID,
	}}
}

func (impl *hospiServiceImpl) AllocateRoom(req dto.AllocateRoomRequest) dto.Response {
	user, err := impl.userRepository.FindByID(req.UserID)

	if user == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	visitor, err := impl.hospiRepository.FindVisitorByUserID(user.ID)

	if visitor == nil && err == nil {
		return dto.Response{Code: http.StatusBadGateway, Message: "Go back to pr desk"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	room, err := impl.hospiRepository.FindRoomByID(req.RoomID)

	if room == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Room not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	if room.Occupied >= room.Capacity {
		return dto.Response{Code: http.StatusBadRequest, Message: "Room is already filled!"}
	}

	roomReg, err := impl.hospiRepository.FindRoomRegByUserID(user.ID)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	if roomReg != nil && roomReg.RoomID != 0 {
		return dto.Response{Code: http.StatusBadRequest, Message: "Room already Allocated to User."}
	}

	bill, err := impl.treasuryRepository.GetBillByUserIDAndPaidTo(user.ID, "townScript")
	if bill == nil && err == nil {
		// room registration not found
		bill = &models.Bill{
			UserID: user.ID,
			Email:  user.Email,
			Time:   time.Now(),
			Mode:   "OFFLINE",
			Amount: req.Amount,
			PaidTo: models.HOSPI,
		}

		err := impl.treasuryRepository.AddBillByModel(bill)

		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Unable to create bills for user"}
		}

		roomReg = &models.RoomReg{
			RoomID:    room.ID,
			Email:     user.Email,
			UserID:    user.ID,
			NoOfDays:  req.NumberOfDays,
			StartDate: time.Now().Format("15-01-2006"),
		}

		err = impl.hospiRepository.AddRoomReg(roomReg)

		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Unable to create room registration for user"}
		}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	} else {
		if roomReg == nil && err == nil {
			return dto.Response{Code: http.StatusBadRequest, Message: "Room registration not found for townscript user!"}
		} else if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Unable to find room registration for user"}
		}
	}

	room.Occupied++
	err = impl.hospiRepository.UpdateRoom(room)

	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Unable to update room, Internal Server Error"}
	}

	visitor.CheckInTime = time.Now()
	visitor.CheckInBillID = bill.ID
	visitor.RoomRegID = roomReg.ID

	err = impl.hospiRepository.UpdateVisitor(visitor)

	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Unable to add visitor"}
	}

	err = impl.hospiRepository.UpdateRoomRegWithRoomID(user.Email, room.ID)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: "Room allocated!"}
}

func (impl *hospiServiceImpl) CheckOut(req dto.CheckOutRequest) dto.Response {

	user, err := impl.userRepository.FindByID(req.UserID)

	if user == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	visitor, err := impl.hospiRepository.FindVisitorByUserID(user.ID)

	if visitor == nil && err == nil {
		return dto.Response{Code: http.StatusBadGateway, Message: "User has never registered as a Visitor"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	err = impl.hospiRepository.CheckoutVisitor(visitor)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "There seems to be an issue checking out the user"}
	}

	err = impl.treasuryRepository.AddCheckOutBill(user, &req)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "There seems to be an error in adding user discount/fine bills"}
	}

	return dto.Response{Code: http.StatusOK, Message: "Checked Out!"}
}
