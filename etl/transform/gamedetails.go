package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"

	"github.com/PuerkitoBio/goquery"
)

// Transforms and consolidates Game Details properties from various sources
func CreateGameDetailsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection, foxSchedulePage *goquery.Selection) pickemstructs.GameDetails {
	var gameID string = common.ParseGameID(espnGameDetails)

	var newRecord pickemstructs.GameDetails = pickemstructs.GameDetails{
		GameID:        gameID,
		League:        gamedetails.ParseLeague(espnGameDetails),
		Week:          gamedetails.ParseWeek(espnGameDetails),
		Year:          gamedetails.ParseYear(espnGameDetails),
		ESPNCode:      gamedetails.ParseEspnGameCode(espnGameDetails),
		CBSCode:       gamedetails.ExtractCbsGameCode(cbsSchedulePage, gameID),
		FoxCode:       gamedetails.ExtractFoxGameCode(foxSchedulePage, gameID),
		VegasCode:     "",
		AwayTeamID:    common.ParseGameSummaryTeamID("away", espnGameDetails),
		HomeTeamID:    common.ParseGameSummaryTeamID("home", espnGameDetails),
		ZuluTimestamp: gamedetails.ParseGameZuluTimestamp(espnGameDetails),
		Broadcast:     gamedetails.ParseBroadcast(espnGameDetails),
		Finished:      gamedetails.ParseGameStatus(espnGameDetails),
	}

	return newRecord
}
