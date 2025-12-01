package gamedetails

import (
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Extracts Fox team code of 'Home' or 'Away' team from a given Fox Game Code
func ExtractFoxTeamCode(foxGameCode string, homeAway string) string {
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
func ExtractFoxGameCode(foxSchedulePage *goquery.Selection, gameID string) string {
	var gameAnchorTags *goquery.Selection = foxSchedulePage.Find("div.scores-app-root").Find("td.broadcast").Find("div").Find("a")
	var foxGameCode string = "foxGameCode"

	gameAnchorTags.EachWithBreak(func(i int, hyperlink *goquery.Selection) bool {
		// Sample Fox Game HREF:
		// https://www.foxsports.com/college-football/bowling-green-falcons-vs-umass-minutemen-nov-25-2025-game-boxscore-42675
		foxGameCode = strings.SplitAfter(hyperlink.AttrOr("href", "gamehref"), "/")[2]
		var gameHREF string = hyperlink.AttrOr("href", "gamehref")
		var lastSlashIndex int = strings.LastIndex(gameHREF, "/")
		var foxGameCode string = gameHREF[lastSlashIndex+1:]
		var awayTeamFoxCode string = ExtractFoxTeamCode(foxGameCode, "AWAY")
		var homeTeamFoxCode string = ExtractFoxTeamCode(foxGameCode, "HOME")
		var awayTeamID string
		var homeTeamID string

		if foxGameCode == "foxGameCode" {
			log.Printf("Failed to extract Fox Game Code from Hyperlink: %v\n", hyperlink)
		}
		if awayTeamFoxCode == "foxTeamCode" {
			log.Printf("Failed to extract Fox Team Code from Hyperlink: %v\n", hyperlink)
		}
		if homeTeamFoxCode == "foxTeamCode" {
			log.Printf("Failed to extract Fox Team Code from Hyperlink: %v\n", hyperlink)
		}

		// Map Fox Team Code to global Team IDs
		mappedAwayTeamCodeValue, exists := utils.FoxTeamCodeToTeamIDmapping[awayTeamFoxCode]
		if exists {
			awayTeamID = mappedAwayTeamCodeValue
		} else {
			awayTeamID = awayTeamFoxCode
		}

		mappedHomeTeamCodeValue, exists := utils.FoxTeamCodeToTeamIDmapping[homeTeamFoxCode]
		if exists {
			homeTeamID = mappedHomeTeamCodeValue
		} else {
			homeTeamID = homeTeamFoxCode
		}

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		}
		return true

	})

	return foxGameCode
}
