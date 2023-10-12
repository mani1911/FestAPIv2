package seeder

import (
	"fmt"
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/utils"
	"github.com/fatih/color"
)

// printing the seeded rows
func PrintSeededRow(tableName string, seedStatus string, element map[string]interface{}) {
	delete(element, "created_at")
	delete(element, "updated_at")
	count := 0
	last := len(element)

	fmt.Print(color.HiCyanString(tableName + " : " + seedStatus + " seed : "))
	for k, v := range element {
		count++
		fmt.Print(color.HiMagentaString(k) + ":")
		fmt.Print(v)
		if count == last {
			fmt.Print("")
		} else {
			fmt.Print(", ")
		}
	}
}

// seeding the table
func seedTable(name string, result map[string][]map[string]interface{}) {
	fmt.Println(color.BlueString("Started seeding the table : " + name))

	seedContent := result[name]
	db := config.GetDB()

	for _, element := range seedContent {
		id := element["id"]
		var count int64

		db.Table(name).Where("id = ?", id).Count(&count)

		if count == 0 {
			element["created_at"] = time.Now()
			element["updated_at"] = time.Now()
			if err := db.Table(name).Create(element).Error; err != nil {
				PrintSeededRow(name, "error in creating", element)
				fmt.Println("\n", color.RedString("Error:"), err)
			} else {
				PrintSeededRow(name, "created", element)
			}
		} else {
			element["updated_at"] = time.Now()
			if err := db.Table(name).Where("id = ?", id).Updates(element).Error; err != nil {
				PrintSeededRow(name, "error in updating", element)
				fmt.Println("\n", color.RedString("Error:"), err)
			} else {
				PrintSeededRow(name, "updated", element)
			}
		}
		fmt.Println("")
	}
}

// seeding data
func SeedData(seeds []string) {
	for _, v := range seeds {
		result := utils.ReadJSON("scripts/seed_db/seed_functions/content/" + v + ".json")
		seedTable(v, result)
	}
}

func DBSeeder() {

	var seeds = []string{
		"colleges",
		"events",
	}

	SeedData(seeds)
}
