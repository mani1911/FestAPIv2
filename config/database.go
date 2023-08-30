package config

import (
	"fmt"

	"github.com/delta/FestAPI/models"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database
var db *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var er error
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	db, er = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	createEnums()
	if er != nil {
		fmt.Println(color.RedString("Error connecting to database"))
	} else {
		fmt.Println(color.GreenString("Database connected"))
	}
}

// Create ENUMS
func createEnums() {
	createGender := db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender') THEN CREATE TYPE gender AS ENUM('MALE','FEMALE','OTHER'); END IF; END$$;")
	if createGender.Error != nil {
		fmt.Println(color.RedString("Error creating Gender ENUM"))
	}
}

// GetDB returns the database
func GetDB() *gorm.DB {
	return db
}

// MigrateDB migrates schemas
func MigrateDB() {
	db := GetDB()

	for _, schema := range []interface{}{&models.College{}, &models.User{}} {
		if err := db.AutoMigrate(&schema); err != nil {
			panic(err)
		}
	}
	fmt.Println(color.BlueString("Migration Success"))
}
