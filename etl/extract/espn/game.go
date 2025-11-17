/*
Package that:

  - Calls the Game Summary endpoint (espnHiddenGameSummaryBaseURL),
    for a given ESPN GameID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"encoding/json"
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
	"io"
	"log"
	"net/http"
)

func GetGame(gameID string) types.ESPNGameDetailsResponse {
	const espnHiddenGameSummaryBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary"
	var espnGameEndpoint string = fmt.Sprintf("%s?event=%s", espnHiddenGameSummaryBaseURL, gameID)

	log.Printf("\nCalling Game %s endpoint: %s", gameID, espnGameEndpoint)
	resp, err := http.Get(espnGameEndpoint)
	if err != nil {
		log.Printf("Error occurred calling ESPN Game Summary Hidden Endpoint for GameID %s:\n%s\n", gameID, err)
		return types.ESPNGameDetailsResponse{}

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred parsing ESPN Game Summary Hidden Endpoint Response for GameID %s: %s\n", gameID, err)
		return types.ESPNGameDetailsResponse{}

	}

	var gameDetails types.ESPNGameDetailsResponse
	jsonerr := json.Unmarshal([]byte(body), &gameDetails)
	if jsonerr != nil {
		log.Printf("Error occurred decoding ESPN Game Summary JSON formatted game details for GameID %s: %s\n", gameID, jsonerr)
		return types.ESPNGameDetailsResponse{}

	}

	log.Printf("gameDetails:\n%v\n", gameDetails)
	return gameDetails
}
