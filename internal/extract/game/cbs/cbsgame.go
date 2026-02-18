// Package cbsgame provides CBS Sports game page web scraping functionality.
// It extracts game-specific data from CBS Sports HTML pages by matching team codes
// and locating the corresponding game odds table.
package cbsgame

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsGame struct {
	CbsOddsPage *goquery.Selection
	GameId      string
}

type instantiator interface {
	scrapeGame() *goquery.Selection
}

// GetGamePage runs the given game instantiator and returns the scraped CBS game odds selection.
func GetGamePage(g instantiator) *goquery.Selection {
	return g.scrapeGame()
}

// getTeamID maps a CBS team code to the global team ID using the package mapping.
func getTeamID(teamCode string) string {
	cbsTeamCodeWithoutAbbr := teamCode[strings.Index(teamCode, "/")+1:]
	teamID, exists := utils.CbsTeamCodeToTeamIDmapping[cbsTeamCodeWithoutAbbr]
	if exists {
		return teamID
	} else {
		return cbsTeamCodeWithoutAbbr
	}
}

// scrapeGame finds the CBS odds block whose away and home team IDs match the game ID.
func (g CbsGame) scrapeGame() *goquery.Selection {
	var cbsGameOddsHTML *goquery.Selection
	gameOddsTables := g.CbsOddsPage.Find(`div.OddsBlock`)

	gameOddsTables.EachWithBreak(func(i int, gameOddsTable *goquery.Selection) bool {
		// Use the table element directly instead of finding it again
		cbsGameOddsHTML = gameOddsTable

		awayTeamID := getTeamID(awayTeam{oddsPageTable: gameOddsTable}.scrapeTeamCode())
		homeTeamID := getTeamID(homeTeam{oddsPageTable: gameOddsTable}.scrapeTeamCode())

		if strings.Contains(g.GameId, awayTeamID) && strings.Contains(g.GameId, homeTeamID) {
			// Break out of loop if GameId string containts awayTeamID and homeTeamID
			return false
		} else {
			// Continue to next iteration
			return true
		}

	})

	return cbsGameOddsHTML
}
