/*
Package that:
  - Calls the Scoreboard endpoint (espnHiddenScoreboardBaseURL),
    for a given week, from the ESPN hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b
*/
package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetSchedule(league string, week uint8) any {
	const espnHiddenScoreboardBaseURL string = "https: //site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard"
	var espnScoreboardEndpoint string = fmt.Sprintf("%s?week=%d", espnHiddenScoreboardBaseURL, week)

	log.Printf("\nCalling Scoreboard endpoint for week %d: %s", week, espnScoreboardEndpoint)
	resp, err := http.Get(espnScoreboardEndpoint)
	if err != nil {
		log.Panicf("Error occurred calling ESPN Scoreboard Summary Hidden Endpoint for week %d:\n%s\n", week, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Error occurred parsing ESPN Scoreboard Summary Hidden Endpoint Response for week %d:\n%s\n", week, err)
	}

	var scheduleDetails any
	jsonerr := json.Unmarshal([]byte(body), &scheduleDetails)
	if jsonerr != nil {
		log.Panicf("Error occurred decoding ESPN Scoreboard Summary JSON formatted schedule details for week %d:\n%s\n", week, jsonerr)
	}

	return scheduleDetails
}
