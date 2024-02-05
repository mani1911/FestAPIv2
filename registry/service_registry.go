package registry

import (
	"github.com/delta/FestAPI/service"
	"github.com/delta/FestAPI/service/impl"
)

func (r *registry) NewUserService() service.UserService {
	return impl.NewUserServiceImpl(
		r.NewUserRepository(),
		r.NewCollegeRepository(),
	)
}

func (r *registry) NewAdminService() service.AdminService {
	return impl.NewAdminServiceImpl(
		r.NewAdminRepository(),
		r.NewUserRepository(),
	)
}

func (r *registry) NewEventService() service.EventService {
	return impl.NewEventServiceImpl(
		r.NewEventRepository(),
		r.NewUserRepository(),
	)
}

func (r *registry) NewHospiService() service.HospiService {
	return impl.NewHospiServiceImpl(
		r.NewHospiRepository(),
		r.NewAdminRepository(),
	)
}

func (r *registry) NewCMSService() service.CMSService {
	return impl.NewCMSServiceImpl(
		r.NewEventRepository(),
	)
}

func (r *registry) NewPublicService() service.PublicService {
	return impl.NewPublicServiceImpl(
		r.NewCollegeRepository(),
	)
}
