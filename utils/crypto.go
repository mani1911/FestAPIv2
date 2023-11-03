package utils

import "golang.org/x/crypto/bcrypt"

func ComapareHashPassword(userPassword []byte, curPassword string) error {
	return bcrypt.CompareHashAndPassword(userPassword, []byte(curPassword))
}

func GenerateHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}
