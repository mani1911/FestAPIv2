package repository

type TShirtsRepository interface {
	UpdateSize(userID uint, size string, rollNo string, code string, screenshotLink string) error
}
