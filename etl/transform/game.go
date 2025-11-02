package transform

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseGameID(espnGameDetails types.ESPNGameDetailsResponse) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = FormatStringID(gameID)
	return formattedGameID
}

func parseLeague(espnGameDetails types.ESPNGameDetailsResponse) string {
	var league string = espnGameDetails.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

func parseWeek(espnGameDetails types.ESPNGameDetailsResponse) int8 {
	var week int8 = espnGameDetails.Header.Week
	return week

}

func parseYear(espnGameDetails types.ESPNGameDetailsResponse) uint16 {
	var year uint16 = espnGameDetails.Header.Season.Year
	return year

}

func parseEspnGameCode(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameCode string = espnGameDetails.Header.ESPNGameCode
	return gameCode

}

func extractCbsGameCode(cbsSchedulePage *goquery.Selection) string {
	var scorecards *goquery.Selection = cbsSchedulePage.Find("div.Page-colMain").Find("div.score-card-container").Find("div.score-cards").Find("div.single-score-card")
	scorecards.Each(func(i int, s *goquery.Selection) {
		var cbsGameCode string = s.AttrOr("data-abbrev", "cbsGameCode")
		var scorecardProgressTable *goquery.Selection = s.Find("div.team-details-wrapper").Find("div.in-progress-table").Find("table").Find("tbody").Find("tr")
		var awayTeamHREF string = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
		var awayTeamCBScodeIndex int = strings.Index(awayTeamHREF, "teams/")
		var awayTeamCBScode string = awayTeamHREF[awayTeamCBScodeIndex:]

		var homeTeamHREF string = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
		var homeTeamCBScodeIndex int = strings.Index(homeTeamHREF, "teams/")
		var homeTeamCBScode string = homeTeamHREF[homeTeamCBScodeIndex:]
	})

}

func parseTeamID(homeAway string, espnGameDetails types.ESPNGameDetailsResponse) string {
	var competitorHomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	if homeAway == competitorHomeAway {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[0].Team.DisplayName
		var formattedTeamID string = FormatStringID(teamID)
		return formattedTeamID
	} else {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[1].Team.DisplayName
		var formattedTeamID string = FormatStringID(teamID)
		return formattedTeamID
	}

}

func parseGameZuluTimestamp(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameDate string = espnGameDetails.Header.Competitions[0].Date
	return gameDate
}

func parseBroadcast(espnGameDetails types.ESPNGameDetailsResponse) string {
	var broadcast string = espnGameDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

func parseGameStatus(espnGameDetails types.ESPNGameDetailsResponse) bool {
	var gameStatus bool = espnGameDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

func Game(espnGameDetails types.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection) types.GameDetails {
	var game types.GameDetails

	game.GameID = parseGameID(espnGameDetails)
	game.League = parseLeague(espnGameDetails)
	game.Week = parseWeek(espnGameDetails)
	game.Year = parseYear(espnGameDetails)
	game.ESPNCode = parseEspnGameCode(espnGameDetails)
	game.CBSCode = extractCbsGameCode(cbsSchedulePage)
	game.FoxCode = ""
	game.VegasCode = ""
	game.AwayTeamID = parseTeamID("away", espnGameDetails)
	game.HomeTeamID = parseTeamID("home", espnGameDetails)
	game.ZuluTimestamp = parseGameZuluTimestamp(espnGameDetails)
	game.Broadcast = parseBroadcast(espnGameDetails)
	game.Finished = parseGameStatus(espnGameDetails)

	return game
}
