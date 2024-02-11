package registry

import (
	"github.com/delta/FestAPI/app"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() app.Controller
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() app.Controller {
	return app.Controller{
		User:     r.NewUserController(),
		Admin:    r.NewAdminController(),
		Event:    r.NewEventController(),
		Hospi:    r.NewHospiController(),
		CMS:      r.NewCMSController(),
		Public:   r.NewPublicController(),
		TShirts:  r.NewTShirtsController(),
		Treasury: r.NewTreasuryController(),
	}
}
