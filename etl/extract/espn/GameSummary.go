package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GameSummary(gameID string) map[string]any {
	const espnHiddenGameSummaryBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary"
	var espnGameEndpoint string = fmt.Sprintf("%s?event=%s", espnHiddenGameSummaryBaseURL, gameID)

	log.Printf("\nCalling Game %s endpoint: %s", gameID, espnGameEndpoint)
	resp, err := http.Get(espnGameEndpoint)
	if err != nil {
		log.Printf("Error occurred calling ESPN Game Summary Hidden Endpoint for GameID %s: %s\n", gameID, err)
		return nil

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred parsing ESPN Game Summary Hidden Endpoint Response for GameID %s: %s\n", gameID, err)
		return nil

	}

	var gameDetails map[string]any
	jsonerr := json.Unmarshal(body, &gameDetails)
	if jsonerr != nil {
		log.Printf("Error occurred decoding ESPN Game Summary JSON formatted game details for GameID %s: %s\n", gameID, err)
		return nil

	}

	log.Println()
	return gameDetails

}
