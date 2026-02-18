// Package espngame provides ESPN game summary API client functionality.
// It calls the ESPN Game Summary API endpoint to retrieve detailed game information
// including scores, betting odds, and game status for both college football (CFB) and NFL.
package espngame

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
)

type EspnNflGame struct {
	GameCode string
}

type EspnCfbGame struct {
	GameCode string
}

type instantiator interface {
	getGameSummary() GameSummaryEndpoint
}

// GetGameSummary runs the given game instantiator and returns the ESPN game summary response.
func GetGameSummary(g instantiator) GameSummaryEndpoint {
	return g.getGameSummary()
}

// fetchGameSummary calls the ESPN game summary endpoint and decodes the JSON response.
func fetchGameSummary(gameCode string, espnGameEndpoint string) GameSummaryEndpoint {
	log.Printf("\nCalling ESPN endpoint for game code %s: %s\n", gameCode, espnGameEndpoint)

	body, err := utils.CallEndpoint(espnGameEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	gameDetails, err := utils.DecodeJSON[GameSummaryEndpoint](body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	//log.Printf("gameDetails:\n%v\n", gameDetails)
	return gameDetails
}

// getGameSummary fetches the ESPN college football game summary for the configured game code.
func (g EspnCfbGame) getGameSummary() GameSummaryEndpoint {
	espnGameEndpoint := fmt.Sprintf("%s%s", utils.ESPN_CFB_GAME_ENDPOINT_URL, g.GameCode)
	return fetchGameSummary(g.GameCode, espnGameEndpoint)

}

// getGameSummary fetches the ESPN NFL game summary for the configured game code.
func (g EspnNflGame) getGameSummary() GameSummaryEndpoint {
	espnGameEndpoint := fmt.Sprintf("%s%s", utils.ESPN_NFL_GAME_ENDPOINT_URL, g.GameCode)
	return fetchGameSummary(g.GameCode, espnGameEndpoint)

}
