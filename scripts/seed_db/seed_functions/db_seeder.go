package seeder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/utils"
	"github.com/fatih/color"
)

type EventDetails struct {
	ID          int    `json:"id"`
	EventName   string `json:"event_name"`
	MaxTeamSize int    `json:"max_team_size"`
}

type Events struct {
	Events []EventDetails `json:"events"`
}

type EventAbstract struct {
	ID              int    `json:"id"`
	EventID         int    `json:"event_id"`
	ForwardEmail    string `json:"forward_email"`
	MaxParticipants int    `json:"max_participants"`
}

type EventAbstracts struct {
	EventAbstracts []EventAbstract `json:"event_abstract_details"`
}

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

func fetchFromCMS(url string) (Events, EventAbstracts) {

	var events Events
	var eventAbstracts EventAbstracts
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Write the body to file
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result interface{}
	_ = json.Unmarshal(out, &result)

	for _, v := range result.(map[string]interface{})["data"].([]interface{}) {
		id := v.(map[string]interface{})["id"].(float64)

		eventURL := config.CmsURL + "/api/clusters/" + fmt.Sprint(id) + "?populate[Cluster_Details][populate][Events][populate]=*"

		resp, err := http.Get(eventURL)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		eventOut, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var eventResult interface{}
		_ = json.Unmarshal(eventOut, &eventResult)

		for _, v := range eventResult.(map[string]interface{})["data"].(map[string]interface{})["attributes"].(map[string]interface{})["Cluster_Details"].(map[string]interface{})["Events"].([]interface{}) {
			var event EventDetails
			event.ID = int(v.(map[string]interface{})["id"].(float64))
			event.EventName = v.(map[string]interface{})["name"].(string)
			event.MaxTeamSize = int(v.(map[string]interface{})["Max_Team_Size"].(float64))
			events.Events = append(events.Events, event)

			if v.(map[string]interface{})["Abstract_Needed"].(bool) {
				var eventAbstract EventAbstract
				eventAbstract.ID = int(v.(map[string]interface{})["id"].(float64))
				eventAbstract.EventID = int(v.(map[string]interface{})["id"].(float64))
				if v.(map[string]interface{})["Forward_Email"] != nil {
					eventAbstract.ForwardEmail = v.(map[string]interface{})["Forward_Email"].(string)
				} else {
					eventAbstract.ForwardEmail = ""
				}
				eventAbstract.MaxParticipants = int(v.(map[string]interface{})["Max_Participants"].(float64))
				eventAbstracts.EventAbstracts = append(eventAbstracts.EventAbstracts, eventAbstract)
			}
		}

	}

	return events, eventAbstracts

}

func DBSeeder() {

	fmt.Println(color.HiYellowString("Started seeding the database"))

	var seeds = []string{
		"colleges",
		"informals_details",
		"hostels",
		"rooms",
	}

	events, eventAbstracts := fetchFromCMS(config.CmsURL + "/api/clusters?populate=*")
	eventsJSON, err := json.Marshal(events)
	if err != nil {
		fmt.Println(err)
	}
	var seedableEvents map[string][]map[string]interface{}
	_ = json.Unmarshal(eventsJSON, &seedableEvents)
	seedTable("events", seedableEvents)

	eventAbstractsJSON, err := json.Marshal(eventAbstracts)
	if err != nil {
		fmt.Println(err)
	}
	var seedableEventAbstracts map[string][]map[string]interface{}
	_ = json.Unmarshal(eventAbstractsJSON, &seedableEventAbstracts)
	seedTable("event_abstract_details", seedableEventAbstracts)

	SeedData(seeds)
}
