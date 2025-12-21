package common

import (
	"have-a-nice-pickem-etl/etl/extract/team/cbs"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsGame struct {
	CbsOddsPage *goquery.Selection
	GameId      string
}

func setTeamID(cbsTeamCode string) string {
	teamID, exists := utils.CbsTeamCodeToTeamIDmapping[cbsTeamCode]
	if exists {
		return teamID
	} else {
		return cbsTeamCode
	}
}

func (g CbsGame) ExtractCbsGameHTML() *goquery.Selection {
	var gameOddsTables *goquery.Selection = g.CbsOddsPage.Find("table.OddsBlock-game")
	var cbsGameOddsHTML *goquery.Selection
	// var cbsGameCode string = "cbsGameCode"

	gameOddsTables.EachWithBreak(func(i int, oddsTable *goquery.Selection) bool {
		cbsGameOddsHTML = oddsTable
		// var cbsGameCode string = oddsTable.AttrOr("data-game-abbrev", "cbsGameCode")
		var awayTeamCBScode string = cbs.CbsAwayTeam{OddsPageTable: oddsTable}.ExtractTeamCode()
		var homeTeamCBScode string = cbs.CbsHomeTeam{OddsPageTable: oddsTable}.ExtractTeamCode()
		//var awayTeamCBScode string = ExtractCbsGameCode(oddsTable, "AWAY")
		//var homeTeamCBScode string = ExtractCbsTeamCode(oddsTable, "HOME")

		/*if cbsGameCode == "cbsGameCode" {
			log.Printf("Failed to extract CBS Game Code from scorecard: %v\n", oddsTable)
		}
		if awayTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Away Team Code from scorecard: %v\n", oddsTable)
		}
		if homeTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Home Team Code from scorecard: %v\n", oddsTable)
		}*/

		// Map CBS Team Code to global Team IDs
		var cbsAwayTeamCodeWithoutAbbr string = awayTeamCBScode[strings.Index(awayTeamCBScode, "/")+1:]
		var cbsHomeTeamCodeWithoutAbbr string = homeTeamCBScode[strings.Index(homeTeamCBScode, "/")+1:]

		var awayTeamID string = setTeamID(cbsAwayTeamCodeWithoutAbbr)
		var homeTeamID string = setTeamID(cbsHomeTeamCodeWithoutAbbr)

		if strings.Contains(g.GameId, awayTeamID) && strings.Contains(g.GameId, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return cbsGameOddsHTML
}

/* Extracts CBS team code of 'Home' or 'Away' team from a given CBS Game Code
func ExtractCbsTeamCode(oddsTable *goquery.Selection, homeAway string) string {
	//var scorecardProgressTable *goquery.Selection = oddsTable.Find("div.in-progress-table").Find("table").Find("tbody")
	var teamHREF string
	var teamCBScode string
	var teamCBScodeIndex int

	switch strings.ToUpper(homeAway) {
	case "HOME":
		teamHREF = oddsTable.Find("tbody").Find("tr").Eq(1).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	case "AWAY":
		teamHREF = oddsTable.Find("tbody").Find("tr").Eq(0).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	default:
		log.Panicf("Invalid input supplied for 'homeAway': %s", homeAway)
		teamHREF = "cbsTeamHREF"
	}

	teamCBScodeIndex = strings.Index(teamHREF, "teams/")
	teamCBScode = teamHREF[teamCBScodeIndex+6:]
	teamCBScode = strings.TrimRight(teamCBScode, "/")

	return teamCBScode
}

// Extracts CBS game code where AwayTeamID and HomeTeamID match with CBS team codes
func ExtractCbsGameCode(cbsOddsPage *goquery.Selection, gameID string) string {
	//var scorecards *goquery.Selection = cbsSchedulePage.Find("div.Page-colMain").Find("div.score-card-container").Find("div.score-cards").Find("div.single-score-card")
	var gameOddsTables *goquery.Selection = cbsOddsPage.Find("table.OddsBlock-game")
	var cbsGameCode string = "cbsGameCode"

	gameOddsTables.EachWithBreak(func(i int, oddsTable *goquery.Selection) bool {
		cbsGameCode = oddsTable.AttrOr("data-game-abbrev", "cbsGameCode")
		var awayTeamCBScode string = ExtractCbsGameCode(oddsTable, "AWAY")
		var homeTeamCBScode string = ExtractCbsTeamCode(oddsTable, "HOME")

		if cbsGameCode == "cbsGameCode" {
			log.Printf("Failed to extract CBS Game Code from scorecard: %v\n", oddsTable)
		}
		if awayTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Away Team Code from scorecard: %v\n", oddsTable)
		}
		if homeTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Home Team Code from scorecard: %v\n", oddsTable)
		}

		// Map CBS Team Code to global Team IDs
		var cbsAwayTeamCodeWithoutAbbr string = awayTeamCBScode[strings.Index(awayTeamCBScode, "/")+1:]
		var cbsHomeTeamCodeWithoutAbbr string = homeTeamCBScode[strings.Index(homeTeamCBScode, "/")+1:]

		var awayTeamID string = setTeamID(cbsAwayTeamCodeWithoutAbbr)
		var homeTeamID string = setTeamID(cbsHomeTeamCodeWithoutAbbr)

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return cbsGameCode
}
*/
