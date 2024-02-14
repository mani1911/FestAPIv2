package main

import (
	"fmt"
	"os"

	"github.com/delta/FestAPI/config"
	seeder "github.com/delta/FestAPI/scripts/seed_db/seed_functions"
	"github.com/fatih/color"
)

func main() {

	// Initing the database
	config.InitConfig()
	config.ConnectDB()
	config.MigrateDB()

	args := os.Args

	if len(args) != 2 {
		fmt.Println("\n", color.RedString("Error:"), "Enter a valid argument")
		return
	}

	// seeding admin table
	if args[1] == "admin" {
		seeder.AdminSeeder()
	}

	// seeding non admin tables
	if args[1] == "db" {
		seeder.DBSeeder()
	}

	// seeding townscript
	if args[1] == "townscript" {
		seeder.TownScriptSeeder()
	}

}
