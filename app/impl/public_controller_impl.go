package impl

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/utils"
	"github.com/labstack/echo/v4"
)

type publicControllerImpl struct {
	PublicService service.PublicService
}

func NewPublicControllerImpl(publicService service.PublicService) app.PublicController {
	return &publicControllerImpl{PublicService: publicService}
}

// @Summary		Get details of all colleges
// @Description	Fetches colleges Id and name of all colleges.
// @ID				Colleges
// @Tags			Public
// @Produce		json
// @Success		200	{object}	[]dto.CollegeResponse
// @Failure		500	{object}	string	"Error fetching colleges"
// @Router			/api/colleges [get]
func (impl *publicControllerImpl) Colleges(c echo.Context) error {

	res := impl.PublicService.AllColleges()

	return utils.SendResponse(c, res.Code, res.Message)
}
