package impl

import (
	"net/http"

	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/dto"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserControllerImpl(userService service.UserService) app.UserController {
	return &userControllerImpl{userService: userService}
}

// @Summary		Authenticate user with DAuth
// @Description	Callback url for DAuth, returns JWT token if successful
// @ID				DAuthUserLogin
// @Tags			User
// @Produce		json
// @Param			code	query		string	true	"DAuth code"
// @Param			site	query		string	true	"type of site"
// @Success		200		{object}	string	"Success"
// @Failure		400		{object}	string	"Invalid Request"
// @Failure		500		{object}	string	"Internal Server Error"
// @Router			/api/user/dauth/callback [get]
func (impl *userControllerImpl) DAuthLogin(c echo.Context) error {
	var req dto.AuthUserRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Code")
	}

	res := impl.userService.DAuthLogin(req)

	if res.Code == http.StatusOK {
		cookie := utils.GenerateCookie(res.Message.(string))
		c.SetCookie(cookie)
	}

	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Authenticate and log in a user.
// @Description	Authenticates a user using email and password.
// @ID				AuthUserLogin
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			request	body		dto.AuthUserLoginRequest	true	"User authentication request"
// @Success		200		{object}	string						"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		500		{object}	string						"Internal Server Error"
// @Router			/api/user/login [post]
func (impl *userControllerImpl) Login(c echo.Context) error {
	var req dto.AuthUserLoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.userService.Login(req)

	if res.Code == http.StatusOK {
		cookie := utils.GenerateCookie(res.Message.(string))
		c.SetCookie(cookie)
	}

	return utils.SendResponse(c, res.Code, res.Message.(string))
}

// @Summary		Register a new user.
// @Description	Register a new user with the provided details.
// @ID				AuthUserRegister
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			request	body		dto.AuthUserRegisterRequest	true	"User registration request"
// @Success		200		{object}	string						"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		500		{object}	string						"Internal Server Error"
// @Router			/api/user/register [post]
func (impl *userControllerImpl) Register(c echo.Context) error {

	log := utils.GetControllerLogger("UserController Register")
	var req dto.AuthUserRegisterRequest
	if err := c.Bind(&req); err != nil {
		// remove after debugging
		log.Error("Error Binding Request. Error : ", err.Error())
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}
	res := impl.userService.Register(req)

	if res.Code == http.StatusOK {
		cookie := utils.GenerateCookie(res.Message.(string))
		c.SetCookie(cookie)
	}

	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Update user information.
// @Description	Update user information with the provided details.
// @ID				AuthUserUpdate
// @Tags			User
// @Accept			json
// @Produce		json
// @Security		middleware.UserAuth
// @Param			request	body		dto.AuthUserUpdateRequest	true	"User update request"
// @Success		200		{object}	string						"Success"
// @Failure		400		{object}	string						"Invalid Request"
// @Failure		401		{object}	string						"Unauthorized"
// @Failure		500		{object}	string						"Internal Server Error"
// @Security		ApiKeyAuth
// @Router			/api/user/update [patch]
func (impl *userControllerImpl) Update(c echo.Context) error {

	// obtaining user id from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	var req dto.AuthUserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	res := impl.userService.Update(req, userID)
	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		Profile information.
// @Description	profile information to be displayed.
// @ID				ProfileDetails
// @Tags			Profile
// @Produce		json
// @Security		middleware.UserAuth
// @Failure		400	{object}	string	"User not found"
// @Failure		500	{object}	string	"Internal Server Error"
// @Router			/api/user/details [get]
func (impl *userControllerImpl) ProfileDetails(c echo.Context) error {

	// obtaining user id from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	res := impl.userService.ProfileDetails(userID)

	return utils.SendResponse(c, res.Code, res.Message)
}

// @Summary		QR Generation.
// @Description	QR for the profile page.
// @ID				ProfileQR
// @Tags			Profile
// @Produce		json
// @Security		middleware.UserAuth
// @Success		200	{object}	string
// @Failure		400	{object}	string	"User not found"
// @Failure		500	{object}	string	"Internal Server Error"
// @Router			/api/user/qr [get]
func (impl *userControllerImpl) QRgeneration(c echo.Context) error {

	// obtaining user id from jwt
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	res := impl.userService.QRgeneration(userID)

	return utils.SendResponse(c, res.Code, res.Message)
}
