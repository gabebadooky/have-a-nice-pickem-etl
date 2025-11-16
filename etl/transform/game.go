package transform

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Generates "GameID" field from AwayTeamID and HomeTeamID from ESPN Game Summary API
func parseGameID(espnGameDetails types.ESPNGameDetailsResponse) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = utils.FormatStringID(gameID)
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
func extractCbsTeamCode(scorecard *goquery.Selection, homeAway string) string {
	var scorecardProgressTable *goquery.Selection = scorecard.Find("div.team-details-wrapper").Find("div.in-progress-table").Find("table").Find("tbody").Find("tr")
	var teamHREF string

	switch strings.ToUpper(homeAway) {
	case "HOME":
		teamHREF = scorecardProgressTable.Find("td.team").Eq(1).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
	case "AWAY":
		teamHREF = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
	default:
		log.Panicf("Invalid input supplied for 'homeAway': %s", homeAway)
		teamHREF = "cbsTeamHREF"
	}

	if teamHREF == "cbsTeamHREF" {
		log.Printf("Failed to extract CBS Team Code from Scorecard: %v", scorecardProgressTable)
		return "cbsTeamCode"
	}

	var teamCBScodeIndex int = strings.Index(teamHREF, "teams/")
	var teamCBScode string = teamHREF[teamCBScodeIndex:]

	return teamCBScode
}

// Extracts CBS game code where AwayTeamID and HomeTeamID match with CBS team codes
func extractCbsGameCode(cbsSchedulePage *goquery.Selection, gameID string) string {
	var scorecards *goquery.Selection = cbsSchedulePage.Find("div.Page-colMain").Find("div.score-card-container").Find("div.score-cards").Find("div.single-score-card")
	var cbsGameCode string = "cbsGameCode"

	scorecards.EachWithBreak(func(i int, scorecard *goquery.Selection) bool {
		cbsGameCode = scorecard.AttrOr("data-abbrev", "cbsGameCode")
		var awayTeamCBScode string = extractCbsTeamCode(scorecard, "AWAY")
		var homeTeamCBScode string = extractCbsTeamCode(scorecard, "HOME")

		if cbsGameCode == "cbsGameCode" {
			log.Panicf("Failed to extract CBS Game Code from scorecard: %v", scorecard)
		}
		if awayTeamCBScode == "cbsTeamCode" {
			log.Panicf("Failed to extract CBS Away Team Code from scorecard: %v", scorecard)
		}
		if homeTeamCBScode == "cbsTeamCode" {
			log.Panicf("Failed to extract CBS Away Team Code from scorecard: %v", scorecard)
		}

		// Map CBS Team Code to global Team IDs
		var awayTeamID string = utils.CbsTeamCodeToTeamIDmapping[awayTeamCBScode]
		var homeTeamID string = utils.CbsTeamCodeToTeamIDmapping[homeTeamCBScode]

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return cbsGameCode
}

// Extracts Fox team code of 'Home' or 'Away' team from a given Fox Game Code
func extractFoxTeamCode(foxGameCode string, homeAway string) string {
	var formattedGameCode string = utils.StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode)

	switch strings.ToUpper(homeAway) {
	case "HOME":
		return strings.Split(formattedGameCode, "-vs-")[0]
	case "AWAY":
		return strings.Split(formattedGameCode, "-vs-")[1]
	default:
		log.Printf("\nInvalid HomeAway value supplied to extractFoxTeamCode function: %s\n", homeAway)
		return "foxTeamCode"
	}

}

// Extracts FOX game code where AwayTeamID and HomeTeamID match with FOX team codes
func extractFoxGameCode(foxSchedulePage *goquery.Selection, gameID string) string {
	var gameAnchorTags *goquery.Selection = foxSchedulePage.Find("div.scores-app-root").Find("td.broadcast").Find("div").Find("a")
	var foxGameCode string = "foxGameCode"

	gameAnchorTags.EachWithBreak(func(i int, hyperlink *goquery.Selection) bool {
		// Sample Fox Game HREF:
		// https://www.foxsports.com/college-football/bowling-green-falcons-vs-umass-minutemen-nov-25-2025-game-boxscore-42675
		foxGameCode = strings.SplitAfter(hyperlink.AttrOr("href", "gamehref"), "/")[2]
		var gameHREF string = hyperlink.AttrOr("href", "gamehref")
		var lastSlashIndex int = strings.LastIndex(gameHREF, "/")
		var foxGameCode string = gameHREF[lastSlashIndex+1:]
		var awayTeamFoxCode string = extractFoxTeamCode(foxGameCode, "AWAY")
		var homeTeamFoxCode string = extractFoxTeamCode(foxGameCode, "HOME")

		if foxGameCode == "foxGameCode" {
			log.Panicf("Failed to extract Fox Game Code from Hyperlink: %v", hyperlink)
		}
		if awayTeamFoxCode == "foxTeamCode" {
			log.Panicf("Failed to extract Fox Team Code from Hyperlink: %v", hyperlink)
		}
		if homeTeamFoxCode == "foxTeamCode" {
			log.Panicf("Failed to extract Fox Team Code from Hyperlink: %v", hyperlink)
		}

		// Map Fox Team Code to global Team IDs
		var awayTeamID string = utils.FoxTeamCodeToTeamIDmapping[awayTeamFoxCode]
		var homeTeamID string = utils.FoxTeamCodeToTeamIDmapping[homeTeamFoxCode]

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return foxGameCode
}

// Generates "TeamID" field for the Home or Away team from ESPN Game Summary API
func parseTeamID(homeAway string, espnGameDetails types.ESPNGameDetailsResponse) string {
	var competitorHomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	if homeAway == competitorHomeAway {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[0].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
		return formattedTeamID
	} else {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[1].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
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
func Game(espnGameDetails types.ESPNGameDetailsResponse, cbsSchedulePage *goquery.Selection, foxSchedulePage *goquery.Selection) types.GameDetails {
	var game types.GameDetails

	game.GameID = parseGameID(espnGameDetails)
	game.League = parseLeague(espnGameDetails)
	game.Week = parseWeek(espnGameDetails)
	game.Year = parseYear(espnGameDetails)
	game.ESPNCode = parseEspnGameCode(espnGameDetails)
	game.CBSCode = extractCbsGameCode(cbsSchedulePage, game.GameID)
	game.FoxCode = extractFoxGameCode(foxSchedulePage, game.GameID)
	game.VegasCode = ""
	game.AwayTeamID = parseTeamID("away", espnGameDetails)
	game.HomeTeamID = parseTeamID("home", espnGameDetails)
	game.ZuluTimestamp = parseGameZuluTimestamp(espnGameDetails)
	game.Broadcast = parseBroadcast(espnGameDetails)
	game.Finished = parseGameStatus(espnGameDetails)

	return game
}
