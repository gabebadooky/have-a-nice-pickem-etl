/*
Package that:

  - Calls the Scoreboard endpoint (espnHiddenScoreboardBaseURL),
    for a given week, from the ESPN hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/sharedtypes"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnCFBSchedule struct {
	Week uint8
}

type EspnNFLSchedule struct {
	Week uint8
}

func decodeEspnScoreboardResponse(body []byte) (sharedtypes.ESPNScheduleResponse, error) {
	return utils.DecodeJSON[sharedtypes.ESPNScheduleResponse](body)
}

// Call CFB ESPN Scoreboard Summary API Endpoint
func (e EspnCFBSchedule) GetScheduleForWeek() sharedtypes.ESPNScheduleResponse {
	var espnScoreboardEndpoint string

	if e.Week > utils.CFB_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL, e.Week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.Week, espnScoreboardEndpoint)
	} else {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_POST_SEASON_SCHEDULE_URL, e.Week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.Week, espnScoreboardEndpoint)
	}

	body, err := utils.CallEndpoint(espnScoreboardEndpoint)
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
func (e EspnNFLSchedule) GetScheduleForWeek() sharedtypes.ESPNScheduleResponse {
	var espnScoreboardEndpoint string

	if e.Week > utils.NFL_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL, e.Week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.Week, espnScoreboardEndpoint)
	} else {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_POST_SEASON_SCHEDULE_URL, e.Week)
		log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", e.Week, espnScoreboardEndpoint)
	}

	body, err := utils.CallEndpoint(espnScoreboardEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	scheduleDetails, err := decodeEspnScoreboardResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return scheduleDetails
}
