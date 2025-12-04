/*
Package that:

  - Calls the Scoreboard endpoint (espnHiddenScoreboardBaseURL),
    for a given week, from the ESPN hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"encoding/json"
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnCFBSchedule struct {
	week uint8
}

type EspnNFLSchedule struct {
	week uint8
}

type ESPNScheduleResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID string `json:"id"`
}

func decodeEspnScoreboardResponse(body []byte) (ESPNScheduleResponse, error) {
	var scheduleDetails ESPNScheduleResponse

	err := json.Unmarshal([]byte(body), &scheduleDetails)
	if err != nil {
		return ESPNScheduleResponse{}, fmt.Errorf("error occurred decoding espn scoreboard summary api endpoint response:\n%s", err)
	}

	return scheduleDetails, nil
}

// Call CFB ESPN Scoreboard Summary API Endpoint
func (e EspnCFBSchedule) GetScheduleForWeek() ESPNScheduleResponse {
	var espnScoreboardEndpoint string

	if e.week > utils.CFB_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL, e.week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.week, espnScoreboardEndpoint)
	} else {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_POST_SEASON_SCHEDULE_URL, e.week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.week, espnScoreboardEndpoint)
	}

	body, err := extract.CallEndpoint(espnScoreboardEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	scheduleDetails, err := decodeEspnScoreboardResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return scheduleDetails
}

// Call NFL ESPN Scoreboard Summary API Endpoint
func (e EspnNFLSchedule) GetScheduleForWeek() ESPNScheduleResponse {
	var espnScoreboardEndpoint string

	if e.week > utils.NFL_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL, e.week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.week, espnScoreboardEndpoint)
	} else {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_POST_SEASON_SCHEDULE_URL, e.week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.week, espnScoreboardEndpoint)
	}

	body, err := extract.CallEndpoint(espnScoreboardEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	scheduleDetails, err := decodeEspnScoreboardResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return scheduleDetails
}
