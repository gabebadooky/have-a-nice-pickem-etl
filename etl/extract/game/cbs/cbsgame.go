package cbsgame

import (
	"fmt"
	cbsteam "have-a-nice-pickem-etl/etl/extract/team/cbs"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type New interface {
	scrapeGame() *goquery.Selection
}

type CbsGame struct {
	CbsOddsPage *goquery.Selection
	GameId      string
}

func GetGamePage(g New) *goquery.Selection {
	return g.scrapeGame()
}

func setTeamID(cbsTeamCode string) string {
	teamID, exists := utils.CbsTeamCodeToTeamIDmapping[cbsTeamCode]
	if exists {
		return teamID
	} else {
		return cbsTeamCode
	}
}

// Extracts CBS game code where AwayTeamID and HomeTeamID match with corresponding CBS team codes
func (g CbsGame) scrapeGame() *goquery.Selection {
	var cbsGameOddsHTML *goquery.Selection
	gameOddsTables := g.CbsOddsPage.Find(`div.OddsBlock`)

	gameOddsTables.EachWithBreak(func(i int, gameOddsTable *goquery.Selection) bool {
		// Use the table element directly instead of finding it again
		cbsGameOddsHTML = gameOddsTable

		if abbrev, ok := gameOddsTable.Attr("data-game-abbrev"); ok {
			fmt.Printf("\ndata-game-abbrev: %s\n", abbrev)
		}

		var awayTeamCBScode string = cbsteam.ExtractCbsTeamCode(cbsteam.CbsAwayTeam{OddsPageTable: cbsGameOddsHTML})
		var homeTeamCBScode string = cbsteam.ExtractCbsTeamCode(cbsteam.CbsHomeTeam{OddsPageTable: cbsGameOddsHTML})

		cbsAwayTeamCodeWithoutAbbr := awayTeamCBScode[strings.Index(awayTeamCBScode, "/")+1:]
		cbsHomeTeamCodeWithoutAbbr := homeTeamCBScode[strings.Index(homeTeamCBScode, "/")+1:]

		awayTeamID := setTeamID(cbsAwayTeamCodeWithoutAbbr)
		homeTeamID := setTeamID(cbsHomeTeamCodeWithoutAbbr)

		if strings.Contains(g.GameId, awayTeamID) && strings.Contains(g.GameId, homeTeamID) {
			// Break out of loop - match found, cbsGameOddsHTML is set
			return false
		} else {
			// Continue to next iteration
			return true
		}

	})

	return cbsGameOddsHTML
}
