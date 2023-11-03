package config

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/joho/godotenv"
)

var ServerPort string
var DBHost string
var DBUser string
var DBPassword string
var DBName string
var DBPort string
var JWTSecret string
var TokenHourLifeSpan string
var DAuthClientID string
var DAuthClientSecret string
var DAuthCallbackURL string
var AdminToken string
var DAuthUserPassword string
var Target string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env"))
	}
	ServerPort = os.Getenv("SERVER_PORT")
	DBHost = os.Getenv("POSTGRES_HOST")
	DBUser = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName = os.Getenv("POSTGRES_DB")
	DBPort = os.Getenv("POSTGRES_PORT")
	JWTSecret = os.Getenv("JWT_SECRET")
	TokenHourLifeSpan = os.Getenv("TOKEN_HOUR_LIFESPAN")
	DAuthClientID = os.Getenv("DAUTH_CLIENT_ID")
	DAuthClientSecret = os.Getenv("DAUTH_CLIENT_SECRET")
	DAuthCallbackURL = os.Getenv("DAUTH_CALLBACK_URL")
	DAuthUserPassword = os.Getenv("DAUTH_USER_PASSWORD")
	Target = os.Getenv("TARGET")
}
