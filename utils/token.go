package utils

import (
	"errors"
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint   `json:"name"`
	Admin  bool   `json:"admin"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTCustomClaimsforQR struct {
	UserEmail string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, Admin bool, AdminRole string) (string, error) {
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

func ParseToken(tokenString string) (*JWTCustomClaimsforQR, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaimsforQR{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*JWTCustomClaimsforQR); ok {
		return claims, nil
	}
	return nil, errors.New("could not extract claims from token")
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
