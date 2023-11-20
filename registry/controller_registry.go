package registry

import (
	"github.com/delta/FestAPI/app"
	"github.com/delta/FestAPI/app/impl"
)

func (r *registry) NewUserController() app.UserController {
	return impl.NewUserControllerImpl(r.NewUserService())
}

func (r *registry) NewAdminController() app.AdminController {
	return impl.NewAdminControllerImpl(r.NewAdminService())
}

func (r *registry) NewEventController() app.EventController {
	return impl.NewEventControllerImpl(r.NewEventService())
}

func (r *registry) NewHospiController() app.HospiController {
	return impl.NewHospiControllerImpl(r.NewHospiService())
}
