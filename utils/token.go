package utils

import (
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint             `json:"name"`
	Admin  bool             `json:"admin"`
	Role   models.AdminRole `json:"role"`
	jwt.RegisteredClaims
}

type JWTCustomClaimsforQR struct {
	UserEmail string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, Admin bool, AdminRole models.AdminRole) (string, error) {
	claims := &JWTCustomClaims{
		userID,
		Admin,
		AdminRole,
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

func GenerateTokenforQR(userEmail string) (string, error) {
	claims := &JWTCustomClaimsforQR{
		userEmail,
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
