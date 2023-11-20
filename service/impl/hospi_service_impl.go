package impl

import (
	"net/http"

	dto "github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
)

type hospiServiceImpl struct {
	hospiRepository repository.HospiRepository
	adminRepository repository.AdminRepository
}

func NewHospiServiceImpl(
	hospiRepository repository.HospiRepository,
	adminRepository repository.AdminRepository) service.HospiService {
	return &hospiServiceImpl{
		hospiRepository: hospiRepository,
		adminRepository: adminRepository,
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

func (impl *hospiServiceImpl) GetRooms() dto.Response {

	res, err := impl.hospiRepository.GetRooms()
	if res == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Rooms not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: res}
}

func (impl *hospiServiceImpl) AddUpdateRoom(req dto.AddUpdateRoomRequest) dto.Response {
	var roomDetails models.Room

	if req.ID == 0 {
		if len(req.Name) == 0 ||
			(req.Capacity) == 0 ||
			(req.HostelID) == 0 {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		}

		roomDetails.Name = req.Name
		roomDetails.Capacity = req.Capacity
		roomDetails.HostelID = req.HostelID

		if err := impl.hospiRepository.AddRoom(&roomDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to add room"}
		}
	} else {
		room, err := impl.hospiRepository.FindRoomByID(req.ID)
		if room == nil && err == nil {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		} else if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}

		if len(req.Name) == 0 ||
			(req.Capacity) == 0 {
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
		}

		roomDetails.ID = req.ID
		roomDetails.Name = req.Name
		roomDetails.Capacity = req.Capacity
		roomDetails.HostelID = room.HostelID

		if err := impl.hospiRepository.UpdateRoom(&roomDetails); err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "failed to update room"}
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
