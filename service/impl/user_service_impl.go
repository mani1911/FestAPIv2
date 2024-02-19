package impl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"net/http"

	"github.com/delta/FestAPI/config"
	dto "github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository    repository.UserRepository
	collegeRepository repository.CollegeRepository
}

func NewUserServiceImpl(
	userRepository repository.UserRepository,
	collegeRepository repository.CollegeRepository) service.UserService {
	return &userServiceImpl{
		userRepository:    userRepository,
		collegeRepository: collegeRepository,
	}
}

func (impl *userServiceImpl) DAuthLogin(req dto.AuthUserRequest) dto.Response {

	log := utils.GetServiceLogger("UserService DAuthLogin")

	// Fetching code from header
	code := req.Code
	site := req.Site

	// Obtaining access token from dauth server
	token, err := utils.GetDAuthToken(code, site)
	if err != nil {
		log.Error("Error getting Auth Token", err.Error())
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error in Authenticating user"}
	}

	// Obtaining user details from dauth server
	user, err := utils.GetDAuthUser(token.AccessToken)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error in Authenticating User"}
	}

	Name := user.Name
	Email := user.Email
	Gender := user.Gender
	Phone := user.Phone
	if len(Name) == 0 || len(Email) == 0 {
		return dto.Response{Code: http.StatusInternalServerError, Message: "User not found"}
	}

	userDetails, err := impl.userRepository.FindByEmail(Email)
	if userDetails == nil && err == nil {
		// Fetching college details
		collegeDetails, err := impl.collegeRepository.FindByName("National Institute of Technology, Tiruchirapalli")
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to create user"}
		}

		// Creating a password for each user
		password := Email + config.DAuthUserPassword
		passHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to create user"}
		}

		// Creating new User object
		userReg := models.User{
			Email:     Email,
			Name:      Name,
			CollegeID: collegeDetails.ID,
			Gender:    models.Gender(Gender),
			Phone:     Phone,
			Password:  passHash,
			IsDauth:   true,
			FullName:  Name,
		}

		//Creating new User
		err = impl.userRepository.CreateUser(&userReg)
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to Create User"}
		}

		// Creating JWT Token for the new user
		jwtToken, err := utils.GenerateToken(userReg.ID, false, "")
		if err != nil {
			return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
		}
		return dto.Response{Code: http.StatusOK, Message: jwtToken}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error in finding User"}
	}

	jwtToken, err := utils.GenerateToken(userDetails.ID, false, "")
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
	}
	err = impl.userRepository.SetDauth(userDetails)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "There seems to be an issue!"}
	}

	return dto.Response{Code: http.StatusOK, Message: jwtToken}
}

func (impl *userServiceImpl) Login(req dto.AuthUserLoginRequest) dto.Response {

	// Get User from database
	userDetails, err := impl.userRepository.FindByEmail(req.Email)

	if userDetails == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}

	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error in finding User"}
	}

	// Comparing passwords
	err = utils.ComapareHashPassword(userDetails.Password, req.Password)
	if err != nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Enter a valid password"}
	}

	// Creating JWT for the user
	jwtToken, err := utils.GenerateToken(userDetails.ID, false, "")
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
	}

	return dto.Response{Code: http.StatusOK, Message: jwtToken}
}

func CheckRecaptcha(response string) error {

	const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", config.RecaptchaSecret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response.
	var body dto.SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	// Check recaptcha verification success.

	if !body.Success {
		return errors.New("Unsuccessful Recaptcha Verify Request")
	}

	return nil
}

func (impl *userServiceImpl) Register(req dto.AuthUserRegisterRequest) dto.Response {

	log := utils.GetServiceLogger("UserService Register")

	// checking if required fields are present
	if len(req.Username) == 0 ||
		len(req.Email) == 0 ||
		len(req.Fullname) == 0 ||
		len(req.Password) == 0 ||
		len(req.Sex) == 0 ||
		len(req.Nationality) == 0 ||
		len(req.Address) == 0 ||
		len(req.Pincode) == 0 ||
		len(req.State) == 0 ||
		len(req.City) == 0 ||
		len(req.Phone) == 0 ||
		len(req.Degree) == 0 ||
		len(req.Year) == 0 ||
		len(req.RecaptchaCode) == 0 ||
		(len(req.College) == 0 && len(req.OtherCollege) == 0) {
		log.Error("User Registration Check Fail")
		return dto.Response{Code: http.StatusBadRequest, Message: "Invalid Request"}
	}
	// Checking if user exists in the database or not
	userDetails, err := impl.userRepository.FindByEmail(req.Email)
	if userDetails == nil && err == nil {

		// Creating password hash
		passwordHash, err := utils.GenerateHashPassword(req.Password)
		if err != nil {
			log.Error("Error Generating Hash. Error", err.Error())
			return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server error"}
		}

		if CheckRecaptcha(req.RecaptchaCode) != nil {
			log.Error("ReCaptcha failed")
			return dto.Response{Code: http.StatusBadRequest, Message: "ReCaptcha failed"}
		}

		// Invalid College Name
		if err = impl.collegeRepository.Exists(req.College); err != nil {
			log.Error("Invalid College Name")
			return dto.Response{Code: http.StatusBadRequest, Message: "Invalid College Name"}
		}

		// Fetch College details
		collegeDetails, err := impl.collegeRepository.FindByName(req.College)
		if err != nil {
			log.Error("Error findind college detail. Error : ", err.Error())
			return dto.Response{Code: http.StatusInternalServerError, Message: "Error finding College"}
		}
		// Creating new user
		userReg := models.User{
			Name:         req.Username,
			FullName:     req.Fullname,
			College:      *collegeDetails,
			OtherCollege: req.OtherCollege,
			Email:        req.Email,
			Gender:       models.Gender(req.Sex),
			Country:      req.Country,
			State:        req.State,
			City:         req.City,
			Address:      req.Address,
			Pincode:      req.Pincode,
			Phone:        req.Phone,
			Password:     passwordHash,
			Sponsor:      req.Sponsor,
			VoucherName:  req.VoucherName,
			ReferralCode: req.ReferralCode,
			Degree:       req.Degree,
			Year:         req.Year,
			Nationality:  req.Nationality,
			IsDauth:      false,
		}

		if err = impl.userRepository.CreateUser(&userReg); err != nil {
			log.Error("Failed to create user. Error : ", err.Error())
			return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to create user"}
		}

		userDetails, err := impl.userRepository.FindByEmail(req.Email)
		if userDetails == nil && err == nil {
			log.Error("User not found")
			return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
		}
		// User Created
		jwtToken, err := utils.GenerateToken(userDetails.ID, false, "")

		if err != nil {
			log.Error("Token Not generated. Error : ", err.Error())
			return dto.Response{Code: http.StatusInternalServerError, Message: "Token Not generated"}
		}

		log.Info("User Registered Successfully")

		return dto.Response{Code: http.StatusOK, Message: jwtToken}
	} else if err != nil {
		log.Error("Error Creating User. Error : ", err.Error())
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error Creating User. Try Later"}
	}

	// User already exists
	log.Error("User already exists")
	return dto.Response{Code: http.StatusBadRequest, Message: "User already exists"}
}

func (impl *userServiceImpl) Update(req dto.AuthUserUpdateRequest, userID uint) dto.Response {

	// Checking if user exists in DB
	userDetails, err := impl.userRepository.FindByID(userID)
	if userDetails == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error finding user. Try Later"}
	}

	// Invalid College Name
	if err = impl.collegeRepository.Exists(req.College); err != nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "Invalid College Name"}
	}

	// Fetch College details
	collegeDetails, err := impl.collegeRepository.FindByName(req.College)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Error finding College"}
	}

	// Updating Details
	userDetails.College = *collegeDetails
	userDetails.OtherCollege = req.OtherCollege
	userDetails.Gender = models.Gender(req.Sex)
	userDetails.Country = req.Country
	userDetails.State = req.State
	userDetails.City = req.City
	userDetails.Address = req.Address
	userDetails.Pincode = req.Pincode
	userDetails.Phone = req.Phone
	userDetails.Sponsor = req.Sponsor
	userDetails.VoucherName = req.VoucherName
	userDetails.ReferralCode = req.ReferralCode
	userDetails.Degree = req.Degree
	userDetails.Year = req.Year
	userDetails.Nationality = req.Nationality

	// Error Updating User
	err = impl.userRepository.Update(userDetails)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Failed to update user"}
	}

	//User Updated
	return dto.Response{Code: http.StatusOK, Message: "Account Updated"}
}

func (impl *userServiceImpl) ProfileDetails(userID uint) dto.Response {

	// Checking if user exists in DB
	userDetails, err := impl.userRepository.FindByID(userID)
	if userDetails == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	collegeDetails, _ := impl.userRepository.FindByCollegeID(userDetails.CollegeID)
	userDetails.College = *collegeDetails
	tShirtDetails, errTShirt := impl.userRepository.FindTShirtSize(userID)
	if errTShirt != nil {
		tShirtDetails = &models.TShirts{Size: "Not Selected"}
	}

	res := dto.UserInfoResponse{
		Name:         userDetails.Name,
		FullName:     userDetails.FullName,
		CollegeID:    userDetails.CollegeID,
		OtherCollege: userDetails.OtherCollege,
		Email:        userDetails.Email,
		College:      dto.CollegeResponse{ID: collegeDetails.ID, Name: collegeDetails.Name},
		Gender:       userDetails.Gender,
		Country:      userDetails.Country,
		State:        userDetails.State,
		City:         userDetails.City,
		Address:      userDetails.Address,
		Pincode:      userDetails.Pincode,
		Phone:        userDetails.Phone,
		Sponsor:      userDetails.Sponsor,
		VoucherName:  userDetails.VoucherName,
		ReferralCode: userDetails.ReferralCode,
		Degree:       userDetails.Degree,
		Year:         userDetails.Year,
		Nationality:  userDetails.Nationality,
		IsDauth:      userDetails.IsDauth,
		TShirtSize:   tShirtDetails.Size,
	}

	return dto.Response{Code: http.StatusOK, Message: res}

}

func (impl *userServiceImpl) QRgeneration(userID uint) dto.Response {

	// Checking if user exists in DB
	userDetails, err := impl.userRepository.FindByID(userID)
	if userDetails == nil && err == nil {
		return dto.Response{Code: http.StatusBadRequest, Message: "User not found"}
	} else if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	// Generate token for userEmail
	token, _ := utils.GenerateTokenforQR(userDetails.Email)

	// Generate the QR code for token
	qr, err := qrcode.New(token, qrcode.Medium)
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	encodeImageBase64 := func(img image.Image) (string, error) {
		var base64String string
		buffer := new(bytes.Buffer)
		err := png.Encode(buffer, img)
		if err != nil {
			return base64String, err
		}
		base64String = base64.StdEncoding.EncodeToString(buffer.Bytes())
		return base64String, nil
	}

	base64Image, err := encodeImageBase64(qr.Image(256))
	if err != nil {
		return dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	return dto.Response{Code: http.StatusOK, Message: "data:image/png;base64," + base64Image}
}
