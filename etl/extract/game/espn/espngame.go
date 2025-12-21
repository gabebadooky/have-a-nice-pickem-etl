/*
Package that:

  - Calls the Game Summary endpoint (espnHiddenGameSummaryBaseURL),
    for a given ESPN GameID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response
*/
package espn

import (
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnGame interface {
	GetGameSummary() GameSummaryEndpoint
}

type EspnNflGame struct {
	GameCode string
}

type EspnCfbGame struct {
	GameCode string
}

func makeAndHandleGameEndpointCall(gameCode string, espnGameEndpoint string) GameSummaryEndpoint {
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

// Call ESPN CFB Game Summary API Endpoint for a given ESPN Game code
func (e EspnCfbGame) GetGameSummary() GameSummaryEndpoint {
	const espnGameEndpoint string = utils.ESPN_CFB_GAME_ENDPOINT_URL
	return makeAndHandleGameEndpointCall(e.GameCode, espnGameEndpoint)
}

// Call ESPN NFL Game Summary API Endpoint for a given ESPN Game code
func (e EspnNflGame) GetGameSummary() GameSummaryEndpoint {
	const espnGameEndpoint string = utils.ESPN_NFL_GAME_ENDPOINT_URL
	return makeAndHandleGameEndpointCall(e.GameCode, espnGameEndpoint)
}
