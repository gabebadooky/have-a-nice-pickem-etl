package gamedetails

import (
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func setTeamID(cbsTeamCode string) string {
	teamID, exists := utils.CbsTeamCodeToTeamIDmapping[cbsTeamCode]
	if exists {
		return teamID
	} else {
		return cbsTeamCode
	}
}

// Extracts CBS team code of 'Home' or 'Away' team from a given CBS Game Code
func ExtractCbsTeamCode(scorecard *goquery.Selection, homeAway string) string {
	var scorecardProgressTable *goquery.Selection = scorecard.Find("div.in-progress-table").Find("table").Find("tbody")
	var teamHREF string
	var teamCBScodeIndex int
	var teamCBScode string

	switch strings.ToUpper(homeAway) {
	case "HOME":
		teamHREF = scorecardProgressTable.Find("td.team").Eq(1).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
	case "AWAY":
		teamHREF = scorecardProgressTable.Find("td.team").Eq(0).Find("div.team-details-wrapper").Find("a").First().AttrOr("href", "cbsTeamHREF")
	default:
		log.Panicf("Invalid input supplied for 'homeAway': %s", homeAway)
		teamHREF = "cbsTeamHREF"
	}

	teamCBScodeIndex = strings.Index(teamHREF, "teams/")
	teamCBScode = teamHREF[teamCBScodeIndex+6:]
	teamCBScode = strings.TrimRight(teamCBScode, "/")

	log.Printf("teamCBScode: %s", teamCBScode)
	return teamCBScode
}

// Extracts CBS game code where AwayTeamID and HomeTeamID match with CBS team codes
func ExtractCbsGameCode(cbsSchedulePage *goquery.Selection, gameID string) string {
	var scorecards *goquery.Selection = cbsSchedulePage.Find("div.Page-colMain").Find("div.score-card-container").Find("div.score-cards").Find("div.single-score-card")
	var cbsGameCode string = "cbsGameCode"

	scorecards.EachWithBreak(func(i int, scorecard *goquery.Selection) bool {
		cbsGameCode = scorecard.AttrOr("data-abbrev", "cbsGameCode")
		var awayTeamCBScode string = ExtractCbsTeamCode(scorecard, "AWAY")
		var homeTeamCBScode string = ExtractCbsTeamCode(scorecard, "HOME")
		var cbsAwayTeamCodeWithoutAbbr string
		var cbsHomeTeamCodeWithoutAbbr string
		var awayTeamID string
		var homeTeamID string

		if cbsGameCode == "cbsGameCode" {
			log.Printf("Failed to extract CBS Game Code from scorecard: %v", scorecard)
		}
		if awayTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Away Team Code from scorecard: %v", scorecard)
		}
		if homeTeamCBScode == "cbsTeamCode" {
			log.Printf("Failed to extract CBS Home Team Code from scorecard: %v", scorecard)
		}

		// Map CBS Team Code to global Team IDs
		cbsAwayTeamCodeWithoutAbbr = awayTeamCBScode[strings.Index(awayTeamCBScode, "/")+1:]
		cbsHomeTeamCodeWithoutAbbr = homeTeamCBScode[strings.Index(homeTeamCBScode, "/")+1:]

		awayTeamID = setTeamID(cbsAwayTeamCodeWithoutAbbr)
		homeTeamID = setTeamID(cbsHomeTeamCodeWithoutAbbr)

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return cbsGameCode
}
