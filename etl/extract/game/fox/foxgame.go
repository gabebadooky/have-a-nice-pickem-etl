package foxgame

import (
	"fmt"
	foxteam "have-a-nice-pickem-etl/etl/extract/team/fox"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type FoxGame interface {
	scrapeGame() *goquery.Selection
}

type FoxCFBGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}

type FoxNFLGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}

func GetGamePage(g FoxGame) *goquery.Selection {
	return g.scrapeGame()
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
	foxGameCode = utils.StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode)
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
func scrapeGameHyperlink(gameID string, urlPrefix string, schedulePage *goquery.Selection) string {
	var foxGameHyperlink string
	// gameAnchorTags := schedulePage.Find("div.scores-app-root").Find("td.broadcast").Find("div").Find("a")
	gameAnchorTags := schedulePage.Find("div.scores-scorechips-container").Find("table.data-table").Find(`td[data-index="3"]`).Find("a")

	gameAnchorTags.EachWithBreak(func(i int, anchorTag *goquery.Selection) bool {
		// Sample Fox Game HREF:
		// https://www.foxsports.com/college-football/bowling-green-falcons-vs-umass-minutemen-nov-25-2025-game-boxscore-42675
		foxGameHyperlink = fmt.Sprintf("%s%s", urlPrefix, anchorTag.AttrOr("href", "gamehref"))

		var foxAwayTeamCode string = foxteam.ExtractFoxTeamCode(foxteam.FoxAwayTeam{FoxGameHyperlink: foxGameHyperlink})
		var foxHomeTeamCode string = foxteam.ExtractFoxTeamCode(foxteam.FoxHomeTeam{FoxGameHyperlink: foxGameHyperlink})

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

func (g FoxCFBGame) scrapeGame() *goquery.Selection {
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_GAME_BASE_URL, g.FoxSchedulePage)
	foxGame := scrapeFoxGame(foxGameHyperlink)
	return foxGame
}

func (g FoxNFLGame) scrapeGame() *goquery.Selection {
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_GAME_BASE_URL, g.FoxSchedulePage)
	foxGame := scrapeFoxGame(foxGameHyperlink)
	return foxGame
}
