package transform

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Generates "GameID" field from AwayTeamID and HomeTeamID from ESPN Game Summary API
func parseGameID(espnGameDetails types.ESPNGameDetailsResponse) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = FormatStringID(gameID)
	return formattedGameID
}

// Parses "League" field from ESPN Game Summary API
func parseLeague(espnGameDetails types.ESPNGameDetailsResponse) string {
	var league string = espnGameDetails.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func parseWeek(espnGameDetails types.ESPNGameDetailsResponse) int8 {
	var week int8 = espnGameDetails.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func parseYear(espnGameDetails types.ESPNGameDetailsResponse) uint16 {
	var year uint16 = espnGameDetails.Header.Season.Year
	return year

}

// Parses "ID" field from ESPN Game Summary API
func parseEspnGameCode(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameCode string = espnGameDetails.Header.ESPNGameCode
	return gameCode

}

// Extracts CBS team code of 'Home' or 'Away' team from a given CBS Game Code
func extractCbsTeamCode(scorecard *goquery.Selection, homeAway string, cbsGameCode string) string {
	var scorecardProgressTable *goquery.Selection = scorecard.Find("div.team-details-wrapper").Find("div.in-progress-table").Find("table").Find("tbody").Find("tr")
	var teamHREF string

	if strings.ToUpper(homeAway) == "HOME" {
		teamHREF = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")

	} else if strings.ToUpper(homeAway) == "AWAY" {
		teamHREF = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
	} else {
		log.Panicf("Invalid input supplied for 'homeAway': %s", homeAway)
	}

	if teamHREF == "cbsTeamHREF" {
		log.Panicf("Failed to extract CBS Team Code from scorecard: %s", scorecardProgressTable)
		return "cbsTeamCode"
	}

	var teamCBScodeIndex int = strings.Index(teamHREF, "teams/")
	var teamCBScode string = teamHREF[teamCBScodeIndex:]

	return teamCBScode
}

// Extracts CBS game code where AwayTeamID and HomeTeamID match with CBS team codes
func extractCbsGameCode(cbsSchedulePage *goquery.Selection, awayTeam string, homeTeam string) string {
	var scorecards *goquery.Selection = cbsSchedulePage.Find("div.Page-colMain").Find("div.score-card-container").Find("div.score-cards").Find("div.single-score-card")
	var cbsGameCode string = "cbsGameCode"

	scorecards.EachWithBreak(func(i int, scorecard *goquery.Selection) bool {
		cbsGameCode = scorecard.AttrOr("data-abbrev", "cbsGameCode")
		var awayTeamCBScode string = extractCbsTeamCode(scorecard, "AWAY", cbsGameCode)
		var homeTeamCBScode string = extractCbsTeamCode(scorecard, "HOME", cbsGameCode)

		if cbsGameCode == "cbsGameCode" {
			log.Panicf("Failed to extract CBS Game Code from scorecard: %s", scorecard)
		}
		if awayTeamCBScode == "cbsTeamCode" {
			log.Panicf("Failed to extract CBS Team Code for %s from scorecard: %s", scorecard)
		}
		if homeTeamCBScode == "cbsTeamCode" {
			log.Panicf("Failed to extract CBS Team Code for %s from scorecard: %s", scorecard)
		}

		if strings.Contains(awayTeamCBScode, awayTeam) && strings.Contains(homeTeamCBScode, homeTeam) {
			return false
		} else {
			return true
		}
	})

	return cbsGameCode

}

// Generates "TeamID" field for the Home or Away team from ESPN Game Summary API
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

// Parses "Date" field from ESPN Game Summary API
func parseGameZuluTimestamp(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameDate string = espnGameDetails.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func parseBroadcast(espnGameDetails types.ESPNGameDetailsResponse) string {
	var broadcast string = espnGameDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Status" field from ESPN Game Summary API
func parseGameStatus(espnGameDetails types.ESPNGameDetailsResponse) bool {
	var gameStatus bool = espnGameDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

// Transforms and consolidates Game properties from various sources
func Game(espnGameDetails types.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection) types.GameDetails {
	var game types.GameDetails

	game.GameID = parseGameID(espnGameDetails)
	game.League = parseLeague(espnGameDetails)
	game.Week = parseWeek(espnGameDetails)
	game.Year = parseYear(espnGameDetails)
	game.ESPNCode = parseEspnGameCode(espnGameDetails)
	game.CBSCode = extractCbsGameCode(cbsSchedulePage, game.AwayTeamID, game.HomeTeamID)
	game.FoxCode = ""
	game.VegasCode = ""
	game.AwayTeamID = parseTeamID("away", espnGameDetails)
	game.HomeTeamID = parseTeamID("home", espnGameDetails)
	game.ZuluTimestamp = parseGameZuluTimestamp(espnGameDetails)
	game.Broadcast = parseBroadcast(espnGameDetails)
	game.Finished = parseGameStatus(espnGameDetails)

	return game
}
