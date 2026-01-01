package espngame

import (
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnGame interface {
	gameSummary() GameSummaryEndpoint
}

type EspnNflGame struct {
	GameCode string
}

type EspnCfbGame struct {
	GameCode string
}

func GetGameSummary(g EspnGame) GameSummaryEndpoint {
	return g.gameSummary()
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
func (g EspnCfbGame) gameSummary() GameSummaryEndpoint {
	const espnGameEndpoint string = utils.ESPN_CFB_GAME_ENDPOINT_URL
	gameSummary := makeAndHandleGameEndpointCall(g.GameCode, espnGameEndpoint)
	return gameSummary
}

// Call ESPN NFL Game Summary API Endpoint for a given ESPN Game code
func (g EspnNflGame) gameSummary() GameSummaryEndpoint {
	const espnGameEndpoint string = utils.ESPN_NFL_GAME_ENDPOINT_URL
	gameSummary := makeAndHandleGameEndpointCall(g.GameCode, espnGameEndpoint)
	return gameSummary
}
