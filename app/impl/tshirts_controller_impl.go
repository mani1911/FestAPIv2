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

type tshirtsControllerImpl struct {
	tShirtsService service.TShirtsService
}

func NewTShirtsControllerImpl(tShirtsService service.TShirtsService) app.TShirtsController {
	return &tshirtsControllerImpl{tShirtsService: tShirtsService}
}

// @Summary		TShirt Size Update
// @Description	Update the TShirt size for the user.
// @ID				TShirts
// @Tags			TShirts
// @Param			request	body	dto.TShirtsUpdateRequest	true	"Add/update tshirt size"
// @Produce		json
// @Success		200	{object}	string	"Updated TShirt Size"
// @Failure		500	{object}	string	"Error Updating TShirt Size"
// @Router			/tshirt/updateSize [post]
func (impl *tshirtsControllerImpl) UpdateSize(c echo.Context) error {

	var req dto.TShirtsUpdateRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JWTCustomClaims)
	userID := claims.UserID

	res := impl.tShirtsService.UpdateSize(dto.TShirtsUpdateDTO{
		UserID:         userID,
		Size:           req.Size,
		Code:           req.Code,
		ScreenshotLink: req.ScreenshotLink,
		RecaptchaCode:  req.RecaptchaCode,
	})

	return utils.SendResponse(c, res.Code, res.Message)
}
