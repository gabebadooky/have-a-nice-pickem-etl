package fox

import (
	"have-a-nice-pickem-etl/etl/extract/team/fox"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type FoxGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}

func setTeamID(foxTeamCode string) string {
	// Map Fox Team Code to global Team IDs
	teamID, exists := utils.FoxTeamCodeToTeamIDmapping[foxTeamCode]
	if exists {
		return teamID
	} else {
		return foxTeamCode
	}
}

// Extracts FOX game code where AwayTeamID and HomeTeamID match with FOX team codes
func (g FoxGame) ExtractFoxGameHTML() *goquery.Selection {
	var gameAnchorTags *goquery.Selection = g.FoxSchedulePage.Find("div.scores-app-root").Find("td.broadcast").Find("div").Find("a")
	var foxGameHTML *goquery.Selection

	gameAnchorTags.EachWithBreak(func(i int, hyperlink *goquery.Selection) bool {
		// Sample Fox Game HREF:
		// https://www.foxsports.com/college-football/bowling-green-falcons-vs-umass-minutemen-nov-25-2025-game-boxscore-42675
		// foxGameCode = strings.SplitAfter(hyperlink.AttrOr("href", "gamehref"), "/")[2]
		var gameHREF string = hyperlink.AttrOr("href", "gamehref")
		var lastSlashIndex int = strings.LastIndex(gameHREF, "/")
		var foxGameCode string = gameHREF[lastSlashIndex+1:]
		//var awayTeamFoxCode string = ExtractFoxTeamCode(foxGameCode, "AWAY")
		//var homeTeamFoxCode string = ExtractFoxTeamCode(foxGameCode, "HOME")
		var foxAwayTeamCode string = fox.FoxAwayTeam{FoxGameCode: foxGameCode}.ExtractFoxTeamCode()
		var foxHomeTeamCode string = fox.FoxHomeTeam{FoxGameCode: foxGameCode}.ExtractFoxTeamCode()
		var awayTeamID string = setTeamID(foxAwayTeamCode)
		var homeTeamID string = setTeamID(foxHomeTeamCode)

		if strings.Contains(g.GameID, awayTeamID) && strings.Contains(g.GameID, homeTeamID) {
			// Break out of loop
			return false
		}
		return true

	})

	return foxGameHTML
}
