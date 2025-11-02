/*
Package that:

  - Calls the Team Summary endpoint (espnHiddenTeamSummaryBaseURL),
    for a given ESPN TeamID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b
  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Team(teamID string) map[string]any {
	const espnHiddenTeamSummaryBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/"
	var espnTeamEndpoint string = fmt.Sprintf("%s%s", espnHiddenTeamSummaryBaseURL, teamID)

	log.Printf("\nCalling Team %s endpoint: %s", teamID, espnTeamEndpoint)
	resp, err := http.Get(espnTeamEndpoint)
	if err != nil {
		log.Printf("Error occurred calling ESPN Team Summary Hidden Endpoint for TeamID %s: %s\n", teamID, err)
		return nil

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred parsing ESPN Team Summary Hidden Endpoint Response for TeamID %s: %s\n", teamID, err)
		return nil

	}

	var teamDetails map[string]any
	jsonerr := json.Unmarshal(body, &teamDetails)
	if jsonerr != nil {
		log.Printf("Error occurred decoding ESPN Team Summary JSON formatted team details for TeamID %s: %s\n", teamID, jsonerr)
		return nil

	}

	log.Println()
	return teamDetails

}
