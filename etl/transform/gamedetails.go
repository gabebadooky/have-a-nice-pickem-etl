package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"

	"github.com/PuerkitoBio/goquery"
)

// Transforms and consolidates Game properties from various sources
func CreateGameDetailsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection, foxSchedulePage *goquery.Selection) pickemstructs.GameDetails {
	var newRecord pickemstructs.GameDetails

	newRecord.GameID = common.ParseGameID(espnGameDetails)
	newRecord.League = gamedetails.ParseLeague(espnGameDetails)
	newRecord.Week = gamedetails.ParseWeek(espnGameDetails)
	newRecord.Year = gamedetails.ParseYear(espnGameDetails)
	newRecord.ESPNCode = gamedetails.ParseEspnGameCode(espnGameDetails)
	newRecord.CBSCode = gamedetails.ExtractCbsGameCode(cbsSchedulePage, newRecord.GameID)
	newRecord.FoxCode = gamedetails.ExtractFoxGameCode(foxSchedulePage, newRecord.GameID)
	newRecord.VegasCode = ""
	newRecord.AwayTeamID = common.ParseTeamID("away", espnGameDetails)
	newRecord.HomeTeamID = common.ParseTeamID("home", espnGameDetails)
	newRecord.ZuluTimestamp = gamedetails.ParseGameZuluTimestamp(espnGameDetails)
	newRecord.Broadcast = gamedetails.ParseBroadcast(espnGameDetails)
	newRecord.Finished = gamedetails.ParseGameStatus(espnGameDetails)

	return newRecord
}
