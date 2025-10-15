package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func callEspnGameDataEndpoint(gameID string) any {
	const espnHiddenGameSummaryBaseURL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary"
	var espnGameEndpoint string = fmt.Sprintf("%s?event=%s", espnHiddenGameSummaryBaseURL, gameID)
	var logmessage string

	logmessage = fmt.Sprintf("\nCalling Game %s endpoint: %s", gameID, espnGameEndpoint)
	fmt.Println(logmessage)
	log.Println(logmessage)

	resp, err := http.Get(espnGameEndpoint)

	if err != nil {
		logmessage = fmt.Sprintf("Error occurred calling ESPN Game Summary Hidden Endpoint for GameID %s: %s", gameID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logmessage = fmt.Sprintf("Error occurred parsing ESPN Game Summary Hidden Endpoint Response for GameID %s: %s", gameID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	var gameDetails any
	log.Println(string(body))

	jsonerr := json.Unmarshal(body, gameDetails)
	if jsonerr != nil {
		logmessage = fmt.Sprintf("Error occurred decoding ESPN Game Summary JSON formatted game details for GameID %s: %s", gameID, err)
		fmt.Println(logmessage)
		log.Println(logmessage)
		return nil

	}

	return gameDetails

}

// "401754528"
