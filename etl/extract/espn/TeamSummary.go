package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func callEspnTeamDataEndpoint(teamID string) any {
	const espnHiddenTeamSummaryBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/"
	var espnTeamEndpoint string = fmt.Sprintf("%s?event=%s", espnHiddenTeamSummaryBaseURL, teamID)
	var logmessage string

	logmessage = fmt.Sprintf("\nCalling Team %s endpoint: %s", teamID, espnTeamEndpoint)
	fmt.Println(logmessage)
	log.Println(logmessage)

	resp, err := http.Get(espnTeamEndpoint)

	if err != nil {
		logmessage = fmt.Sprintf("Error occurred calling ESPN Team Summary Hidden Endpoint for TeamID %s: %s", teamID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logmessage = fmt.Sprintf("Error occurred parsing ESPN Team Summary Hidden Endpoint Response for TeamID %s: %s", teamID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	var teamDetails any
	log.Println(string(body))

	jsonerr := json.Unmarshal(body, teamDetails)
	if jsonerr != nil {
		logmessage = fmt.Sprintf("Error occurred decoding ESPN Team Summary JSON formatted game details for TeamID %s: %s", teamID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	return teamDetails

}

// "401754528"
