package seeder

import (
	"fmt"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/models"
	"github.com/delta/FestAPI/utils"
	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSeeder() {

	db := config.GetDB()

	name := "admin"
	result := utils.ReadJSON("scripts/seed_db/seed_functions/content/" + name + ".json")
	seedContent := result[name]

	for _, element := range seedContent {
		var adminDetails models.Admin
		username := fmt.Sprint(element["username"])
		password := fmt.Sprint(element["password"])
		role := fmt.Sprint(element["role"])
		if err := db.Where("username = ? ", username).First(&adminDetails).Error; err != nil {
			// If admin doesn't exist
			if err == gorm.ErrRecordNotFound {
				passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
				if err != nil {
					PrintSeededRow(name, "error in creating", element)
					fmt.Println("\n", color.RedString("Error:"), err)
					return
				}
				adminRole := models.AdminRole(role)
				if len(adminRole) == 0 {
					PrintSeededRow(name, "error in creating", element)
					fmt.Println("\n", color.RedString("Error:"), "no Admin Role provided")
					return
				}
				adminReg := models.Admin{
					Username: username,
					Password: passwordHash,
					Role:     adminRole,
				}
				if err := db.Create(&adminReg).Error; err != nil {
					PrintSeededRow(name, "error in creating", element)
					fmt.Println("\n", color.RedString("Error:"), err)
					return
				}
				PrintSeededRow(name, "created", element)
				fmt.Println('\n')
				continue
			}
		}

		// if admin exists: Update details
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			PrintSeededRow(name, "error in updating", element)
			fmt.Println("\n", color.RedString("Error:"), err)
			return
		}
		adminRole := models.AdminRole(role)
		if len(adminRole) == 0 {
			PrintSeededRow(name, "error in updating", element)
			fmt.Println("\n", color.RedString("Error:"), "no Admin Role provided")
			continue
		}
		adminDetails.Password = passwordHash
		adminDetails.Role = adminRole
		if err := db.Save(&adminDetails).Error; err != nil {
			PrintSeededRow(name, "error in updating", element)
			fmt.Println("\n", color.RedString("Error:"), err)
			return
		}
		PrintSeededRow(name, "updated", element)
		fmt.Println('\n')

	}

}
