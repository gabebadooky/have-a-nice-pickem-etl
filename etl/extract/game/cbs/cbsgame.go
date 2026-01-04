package cbsgame

import (
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
	gameOddsTables := g.CbsOddsPage.Find("table.OddsBlock-game")

	gameOddsTables.EachWithBreak(func(i int, oddsTable *goquery.Selection) bool {
		var awayTeamCBScode string
		var homeTeamCBScode string

		cbsGameOddsHTML = oddsTable
		awayTeamCBScode = cbsteam.ExtractCbsTeamCode(cbsteam.CbsAwayTeam{OddsPageTable: oddsTable})
		homeTeamCBScode = cbsteam.ExtractCbsTeamCode(cbsteam.CbsHomeTeam{OddsPageTable: oddsTable})
		cbsAwayTeamCodeWithoutAbbr := awayTeamCBScode[strings.Index(awayTeamCBScode, "/")+1:]
		cbsHomeTeamCodeWithoutAbbr := homeTeamCBScode[strings.Index(homeTeamCBScode, "/")+1:]
		awayTeamID := setTeamID(cbsAwayTeamCodeWithoutAbbr)
		homeTeamID := setTeamID(cbsHomeTeamCodeWithoutAbbr)

		if strings.Contains(g.GameId, awayTeamID) && strings.Contains(g.GameId, homeTeamID) {
			// Break out of loop
			return false
		} else {
			return true
		}

	})

	return cbsGameOddsHTML
}
