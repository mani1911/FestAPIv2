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

func (r *registry) NewHospiRepository() repository.HospiRepository {
	return impl.NewHospiRepositoryImpl(r.db)
}

func (r *registry) NewTShirtsRepository() repository.TShirtsRepository {
	return impl.NewTShirtsRepositoryImpl(r.db)
}

func (r *registry) NewTreasuryRepository() repository.TreasuryRepository {
	return impl.NewTreasuryRepositoryImpl(r.db)
}
