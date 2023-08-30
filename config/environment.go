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
}
