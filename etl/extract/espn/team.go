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
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/utils"
	"io"
	"log"
	"net/http"
)

// Call ESPN Team Summary API Endpoint for a given ESPN team code
func Team(espnTeamCode string) pickemstructs.TeamSummaryResponse {
	const espnHiddenTeamSummaryBaseURL string = utils.ESPN_CFB_TEAM_ENDPOINT_URL
	var espnTeamEndpoint string = fmt.Sprintf("%s/%s", espnHiddenTeamSummaryBaseURL, espnTeamCode)

	log.Printf("\nCalling Team %s endpoint: %s\n", espnTeamCode, espnTeamEndpoint)
	resp, err := http.Get(espnTeamEndpoint)
	if err != nil {
		log.Printf("Error occurred calling ESPN Team Summary Hidden Endpoint for TeamID %s: %s\n", espnTeamCode, err)
		return pickemstructs.TeamSummaryResponse{}

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred parsing ESPN Team Summary Hidden Endpoint Response for TeamID %s: %s\n", espnTeamCode, err)
		return pickemstructs.TeamSummaryResponse{}

	}

	var teamDetails pickemstructs.TeamSummaryResponse
	jsonerr := json.Unmarshal(body, &teamDetails)
	if jsonerr != nil {
		log.Printf("Error occurred decoding ESPN Team Summary JSON formatted team details for TeamID %s: %s\n", espnTeamCode, jsonerr)
		return pickemstructs.TeamSummaryResponse{}

	}

	//log.Println(teamDetails)
	return teamDetails

}
