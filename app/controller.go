package app

type Controller struct {
	User     UserController
	Event    EventController
	Admin    AdminController
	Hospi    HospiController
	CMS      CMSController
	Public   PublicController
	TShirts  TShirtsController
	Treasury TreasuryController
	PR       PRController
}
