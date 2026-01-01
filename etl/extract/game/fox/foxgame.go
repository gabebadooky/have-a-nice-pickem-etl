package foxgame

import (
	foxteam "have-a-nice-pickem-etl/etl/extract/team/fox"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
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

func parseGameCodeFromGameHREF(gameHyperlink string) string {
	lastSlashIndex := strings.LastIndex(gameHyperlink, "/")
	foxGameCode := gameHyperlink[lastSlashIndex+1:]
	return foxGameCode
}

func scrapeFoxGame(foxGameHyperlink string) *goquery.Selection {
	log.Printf("\nRequesting Fox Game page: %s\n", foxGameHyperlink)

	page, err := utils.GetGoQuerySelectionBody(foxGameHyperlink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Extracts FOX game code where AwayTeamID and HomeTeamID match with corresponding FOX team codes
func extractGameHyperlink(gameID string, schedulePage *goquery.Selection) string {
	var foxGameHyperlink string
	gameAnchorTags := schedulePage.Find("div.scores-app-root").Find("td.broadcast").Find("div").Find("a")

	gameAnchorTags.EachWithBreak(func(i int, gameHyperlink *goquery.Selection) bool {
		// Sample Fox Game HREF:
		// https://www.foxsports.com/college-football/bowling-green-falcons-vs-umass-minutemen-nov-25-2025-game-boxscore-42675
		var foxAwayTeamCode string
		var foxHomeTeamCode string

		foxGameHyperlink = gameHyperlink.AttrOr("href", "gamehref")
		foxGameCode := parseGameCodeFromGameHREF(foxGameHyperlink)
		foxAwayTeamCode = foxteam.ExtractFoxTeamCode(foxteam.FoxAwayTeam{FoxGameCode: foxGameCode})
		foxHomeTeamCode = foxteam.ExtractFoxTeamCode(foxteam.FoxAwayTeam{FoxGameCode: foxGameCode})
		awayTeamID := setTeamID(foxAwayTeamCode)
		homeTeamID := setTeamID(foxHomeTeamCode)

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop
			return false
		}
		return true

	})

	return foxGameHyperlink
}

func (g FoxGame) ExtractFoxGameHTML() *goquery.Selection {
	foxGameHyperlink := extractGameHyperlink(g.GameID, g.FoxSchedulePage)
	foxGame := scrapeFoxGame(foxGameHyperlink)
	return foxGame
}
