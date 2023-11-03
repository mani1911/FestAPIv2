package registry

import (
	"github.com/delta/FestAPI/repository"
	"github.com/delta/FestAPI/repository/impl"
)

func (r *registry) NewUserRepository() repository.UserRepository {
	return impl.NewUserRepositoryImpl(r.db)
}

func (r *registry) NewAdminRepository() repository.AdminRepository {
	return impl.NewAdminRepositoryImpl(r.db)
}

func (r *registry) NewEventRepository() repository.EventRepository {
	return impl.NewEventRepositoryImpl(r.db)
}

func (r *registry) NewCollegeRepository() repository.CollegeRepository {
	return impl.NewCollegeRepositoryImpl(r.db)
}
