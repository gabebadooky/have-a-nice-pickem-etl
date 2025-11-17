package transform

import (
	"have-a-nice-pickem-etl/etl/transform/gamedetails"
	"have-a-nice-pickem-etl/etl/types"

	"github.com/PuerkitoBio/goquery"
)

// Transforms and consolidates Game properties from various sources
func CreateRecord(espnGameDetails types.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection, foxSchedulePage *goquery.Selection) types.GameDetails {
	var game types.GameDetails

	game.GameID = gamedetails.ParseGameID(espnGameDetails)
	game.League = gamedetails.ParseLeague(espnGameDetails)
	game.Week = gamedetails.ParseWeek(espnGameDetails)
	game.Year = gamedetails.ParseYear(espnGameDetails)
	game.ESPNCode = gamedetails.ParseEspnGameCode(espnGameDetails)
	game.CBSCode = gamedetails.ExtractCbsGameCode(cbsSchedulePage, game.GameID)
	game.FoxCode = gamedetails.ExtractFoxGameCode(foxSchedulePage, game.GameID)
	game.VegasCode = ""
	game.AwayTeamID = gamedetails.ParseTeamID("away", espnGameDetails)
	game.HomeTeamID = gamedetails.ParseTeamID("home", espnGameDetails)
	game.ZuluTimestamp = gamedetails.ParseGameZuluTimestamp(espnGameDetails)
	game.Broadcast = gamedetails.ParseBroadcast(espnGameDetails)
	game.Finished = gamedetails.ParseGameStatus(espnGameDetails)

	return game
}
