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
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/utils"
	"io"
	"log"
	"net/http"
)

// Call ESPN Game Summary API Endpoint for a given ESPN Game code
func GetGame(gameID string) pickemstructs.ESPNGameDetailsResponse {
	const espnHiddenGameSummaryBaseURL string = utils.ESPN_CFB_GAME_ENDPOINT_URL
	var espnGameEndpoint string = fmt.Sprintf("%s?event=%s", espnHiddenGameSummaryBaseURL, gameID)

	log.Printf("\nCalling Game %s endpoint: %s\n", gameID, espnGameEndpoint)
	resp, err := http.Get(espnGameEndpoint)
	if err != nil {
		log.Printf("Error occurred calling ESPN Game Summary Hidden Endpoint for GameID %s:\n%s\n", gameID, err)
		return pickemstructs.ESPNGameDetailsResponse{}

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred parsing ESPN Game Summary Hidden Endpoint Response for GameID %s: %s\n", gameID, err)
		return pickemstructs.ESPNGameDetailsResponse{}

	}

	var gameDetails pickemstructs.ESPNGameDetailsResponse
	jsonerr := json.Unmarshal([]byte(body), &gameDetails)
	if jsonerr != nil {
		log.Printf("Error occurred decoding ESPN Game Summary JSON formatted game details for GameID %s: %s\n", gameID, jsonerr)
		return pickemstructs.ESPNGameDetailsResponse{}

	}

	//log.Printf("gameDetails:\n%v\n", gameDetails)
	return gameDetails
}
