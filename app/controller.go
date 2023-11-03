package app

type Controller struct {
	User  UserController
	Event EventController
	Admin AdminController
}
