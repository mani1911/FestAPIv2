package repository

type TShirtsRepository interface {
	UpdateSize(userID uint, size string, rollNo string) error
}
