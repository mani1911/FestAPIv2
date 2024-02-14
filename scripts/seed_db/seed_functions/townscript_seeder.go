package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/delta/FestAPI/config"
	"github.com/delta/FestAPI/dto"
	repositoryImpl "github.com/delta/FestAPI/repository/impl"
)

type TownScriptAttendeeResponse struct {
	Result string `json:"result"`
	Data   string `json:"data"`
}

func TownScriptSeeder() {
	eventCodes := strings.Split(config.TownScriptEventCodes, ",")
	for _, eventCode := range eventCodes {
		request := &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:   "https",
				Host:     "www.townscript.com",
				Path:     "/api/registration/getRegisteredUsers",
				RawQuery: "eventCode=" + eventCode,
			},
			Proto: "HTTP/1.1",
			Header: http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {config.TownScriptToken},
			},
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		var townScriptAttendeeResponse TownScriptAttendeeResponse
		err = json.Unmarshal(body, &townScriptAttendeeResponse)
		if err != nil {
			fmt.Println(err)
		}
		var townScriptAttendees []dto.TownScriptRequest
		err = json.Unmarshal([]byte(townScriptAttendeeResponse.Data), &townScriptAttendees)
		if err != nil {
			fmt.Println(err)
		}
		impl := repositoryImpl.NewTreasuryRepositoryImpl(config.GetDB())
		for _, attendee := range townScriptAttendees {
			if err := impl.Townscript(&attendee); err != nil {
				fmt.Println(err)
			}
		}
	}
}
