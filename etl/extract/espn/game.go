/*
Package that:

  - Calls the Game Summary endpoint (espnHiddenGameSummaryBaseURL),
    for a given ESPN GameID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnCFBGame struct {
	gameCode string
}

type EspnNFLGame struct {
	gameCode string
}

type ESPNGameDetailsResponse struct {
	GameInfo   GameInfo     `json:"gameInfo"`
	Header     Header       `json:"header"`
	Pickcenter []Pickcenter `json:"pickcenter"`
	Predictor  Predictor    `json:"predictor"`
}

type Header struct {
	Week         int8           `json:"week"`
	Season       Season         `json:"season"`
	ESPNGameCode string         `json:"id"`
	Competitions []Competitions `json:"competitions"`
	League       League         `json:"league"`
}

type Season struct {
	Year uint16 `json:"year"`
}

type Competitions struct {
	Competitors []Competitors `json:"competitors"`
	Date        string        `json:"date"`
	Broadcasts  []Broadcasts  `json:"broadcasts"`
	Status      Status        `json:"status"`
}

type Competitors struct {
	HomeAway   string         `json:"homeAway"`
	Linescores []Linescore    `json:"linescores"`
	Score      string         `json:"score"`
	Team       CompetitorTeam `json:"team"`
}

type Linescore struct {
	DisplayValue string `json:"displayValue"`
}

type CompetitorTeam struct {
	DisplayName string `json:"displayName"`
}

type Broadcasts struct {
	Media Media `json:"media"`
}

type Media struct {
	ShortName string `json:"shortName"`
}

type Status struct {
	Type Type `json:"type"`
}

type Type struct {
	Completed bool `json:"completed"`
}

type League struct {
	Abbreviation string `json:"abbreviation"`
}

type GameInfo struct {
	Venue Venue `json:"venue"`
}

type Venue struct {
	FullName string  `json:"fullName"`
	Address  Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipCode"`
	Country string `json:"country"`
}

type Pickcenter struct {
	OverUnder    float32  `json:"overUnder"`
	Spread       float32  `json:"spread"`
	AwayTeamOdds TeamOdds `json:"awayTeamOdds"`
	HomeTeamOdds TeamOdds `json:"homeTeamOdds"`
}

type TeamOdds struct {
	Moneyline int16 `json:"moneyline"`
}

type Predictor struct {
	AwayTeam PredictorTeam `json:"awayTeam"`
	HomeTeam PredictorTeam `json:"homeTeam"`
}

type PredictorTeam struct {
	GameProjection string `json:"gameProjection"`
}

func decodeEspnGameResponse(body []byte) (ESPNGameDetailsResponse, error) {
	return utils.DecodeJSON[ESPNGameDetailsResponse](body)
}

// Call ESPN CFB Game Summary API Endpoint for a given ESPN Game code
func (e EspnCFBGame) GetGameSummary() ESPNGameDetailsResponse {
	const espnGameEndpoint string = utils.ESPN_CFB_GAME_ENDPOINT_URL
	log.Printf("\nCalling ESPN endpoint for CFB game code %s: %s\n", e.gameCode, espnGameEndpoint)

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
func (e EspnNFLGame) GetGameSummary() ESPNGameDetailsResponse {
	const espnGameEndpoint string = utils.ESPN_NFL_GAME_ENDPOINT_URL
	log.Printf("\nCalling ESPN endpoint for NFL game code %s: %s\n", e.gameCode, espnGameEndpoint)

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
