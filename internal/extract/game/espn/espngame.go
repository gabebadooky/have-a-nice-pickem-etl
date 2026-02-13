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

func GetGameSummary(g instantiator) GameSummaryEndpoint {
	return g.getGameSummary()
}

// Make and handle API call to given ESPN game summary endpoint
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

// Fetch ESPN CFB Game Summary for a given ESPN Game code
func (g EspnCfbGame) getGameSummary() GameSummaryEndpoint {
	espnGameEndpoint := fmt.Sprintf("%s%s", utils.ESPN_CFB_GAME_ENDPOINT_URL, g.GameCode)
	return fetchGameSummary(g.GameCode, espnGameEndpoint)

}

// Fetch ESPN NFL Game Summary a given ESPN Game code
func (g EspnNflGame) getGameSummary() GameSummaryEndpoint {
	espnGameEndpoint := fmt.Sprintf("%s%s", utils.ESPN_NFL_GAME_ENDPOINT_URL, g.GameCode)
	return fetchGameSummary(g.GameCode, espnGameEndpoint)

}
