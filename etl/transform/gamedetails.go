package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"

	"github.com/PuerkitoBio/goquery"
)

type GameDetails struct {
	GameID        string
	League        string
	Week          int8
	Year          uint16
	ESPNCode      string
	CBSCode       string
	FoxCode       string
	VegasCode     string
	AwayTeamID    string
	HomeTeamID    string
	ZuluTimestamp string
	Broadcast     string
	Finished      bool
}

// Transforms and consolidates Game Details properties from various sources
func CreateGameDetailsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection, foxSchedulePage *goquery.Selection) GameDetails {
	var newRecord GameDetails

	newRecord.GameID = common.ParseGameID(espnGameDetails)
	newRecord.League = gamedetails.ParseLeague(espnGameDetails)
	newRecord.Week = gamedetails.ParseWeek(espnGameDetails)
	newRecord.Year = gamedetails.ParseYear(espnGameDetails)
	newRecord.ESPNCode = gamedetails.ParseEspnGameCode(espnGameDetails)
	newRecord.CBSCode = gamedetails.ExtractCbsGameCode(cbsSchedulePage, newRecord.GameID)
	newRecord.FoxCode = gamedetails.ExtractFoxGameCode(foxSchedulePage, newRecord.GameID)
	newRecord.VegasCode = ""
	newRecord.AwayTeamID = common.ParseGameSummaryTeamID("away", espnGameDetails)
	newRecord.HomeTeamID = common.ParseGameSummaryTeamID("home", espnGameDetails)
	newRecord.ZuluTimestamp = gamedetails.ParseGameZuluTimestamp(espnGameDetails)
	newRecord.Broadcast = gamedetails.ParseBroadcast(espnGameDetails)
	newRecord.Finished = gamedetails.ParseGameStatus(espnGameDetails)

	return newRecord
}
