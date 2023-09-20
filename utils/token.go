package utils

import (
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint `json:"name"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, Admin bool) (string, error) {
	claims := &JWTCustomClaims{
		userID,
		Admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
