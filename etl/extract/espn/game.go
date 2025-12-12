/*
Package that:

  - Calls the Game Summary endpoint (espnHiddenGameSummaryBaseURL),
    for a given ESPN GameID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"have-a-nice-pickem-etl/etl/sharedtypes"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnCFBGame struct {
	GameCode string
}

type EspnNFLGame struct {
	GameCode string
}

func decodeEspnGameResponse(body []byte) (sharedtypes.ESPNGameDetailsResponse, error) {
	return utils.DecodeJSON[sharedtypes.ESPNGameDetailsResponse](body)
}

// Call ESPN CFB Game Summary API Endpoint for a given ESPN Game code
func (e EspnCFBGame) GetGameSummary() sharedtypes.ESPNGameDetailsResponse {
	const espnGameEndpoint string = utils.ESPN_CFB_GAME_ENDPOINT_URL
	log.Printf("\nCalling ESPN endpoint for CFB game code %s: %s\n", e.GameCode, espnGameEndpoint)

	body, err := utils.CallEndpoint(espnGameEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	gameDetails, err := decodeEspnGameResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	//log.Printf("gameDetails:\n%v\n", gameDetails)
	return gameDetails
}

// Call ESPN NFL Game Summary API Endpoint for a given ESPN Game code
func (e EspnNFLGame) GetGameSummary() sharedtypes.ESPNGameDetailsResponse {
	const espnGameEndpoint string = utils.ESPN_NFL_GAME_ENDPOINT_URL
	log.Printf("\nCalling ESPN endpoint for NFL game code %s: %s\n", e.GameCode, espnGameEndpoint)

	body, err := utils.CallEndpoint(espnGameEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	gameDetails, err := decodeEspnGameResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	//log.Printf("gameDetails:\n%v\n", gameDetails)
	return gameDetails
}
