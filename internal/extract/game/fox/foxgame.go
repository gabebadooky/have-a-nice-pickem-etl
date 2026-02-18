// Package foxgame provides Fox Sports game page web scraping functionality.
// It extracts game-specific data from Fox Sports HTML pages including boxscores,
// statistics, and betting odds by matching team codes and locating game hyperlinks.
package foxgame

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type FoxGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}

/*type FoxCFBGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}

type FoxNFLGame struct {
	FoxSchedulePage *goquery.Selection
	GameID          string
}*/

// scrapeGame function return type
type FoxGamePages struct {
	BoxscorePage *goquery.Selection
	StatsPage    *goquery.Selection
	OddsPage     *goquery.Selection
}

type instantiator interface {
	scrapeGame() FoxGamePages
}

// GetGamePages runs the given game instantiator and returns boxscore, stats, and odds page selections.
func GetGamePages(g instantiator) FoxGamePages {
	return g.scrapeGame()
}

/* // Retrieve string after first occurrence of "/" in gameHyperlink
func parseGameCodeFromGameHREF(gameHyperlink string) string {
	lastSlashIndex := strings.LastIndex(gameHyperlink, "/")
	foxGameCode := gameHyperlink[lastSlashIndex+1:]
	foxGameCode = utils.StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode)
	return foxGameCode
}*/

// getTeamID maps a Fox team code to the global team ID using the package mapping.
func getTeamID(foxTeamCode string) string {
	teamID, exists := utils.FoxTeamCodeToTeamIDmapping[foxTeamCode]
	if exists {
		return teamID
	} else {
		return foxTeamCode
	}
}

// scrapeFoxGame fetches the Fox Sports page at the given URL and returns its body as a goquery selection.
func scrapeFoxGame(foxGameHyperlink string) *goquery.Selection {
	var page *goquery.Selection
	log.Printf("\nRequesting Fox Game page: %s\n", foxGameHyperlink)

	page, err := utils.GetGoQuerySelectionBody(foxGameHyperlink)
	if err != nil {
		log.Printf("Error occuring scraping %s\n\n", foxGameHyperlink)
		return nil
	}

	return page
}

// scrapeGameHyperlink finds the Fox game URL whose away and home team IDs match the given gameID.
func scrapeGameHyperlink(gameID string, urlPrefix string, schedulePage *goquery.Selection) string {
	var foxGameHyperlink string
	gameAnchorTags := schedulePage.Find("div.scores-scorechips-container").Find("table.data-table").Find(`td[data-index="3"]`).Find("a")

	gameAnchorTags.EachWithBreak(func(i int, anchorTag *goquery.Selection) bool {
		foxGameHyperlink = fmt.Sprintf("%s%s", urlPrefix, anchorTag.AttrOr("href", "gamehref"))
		awayTeamID := getTeamID(awayTeam{hyperlink: foxGameHyperlink}.scrapeTeamCode())
		homeTeamID := getTeamID(homeTeam{hyperlink: foxGameHyperlink}.scrapeTeamCode())

		if strings.Contains(gameID, awayTeamID) && strings.Contains(gameID, homeTeamID) {
			// Break out of loop `gameID` string contains `awayTeamID` and `homeTeamID`
			return false
		}
		return true

	})

	return foxGameHyperlink
}

/*func (g FoxCFBGame) scrapeGame() FoxGamePages {
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_BASE_URL, g.FoxSchedulePage)

	foxGameBoxscoreHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_BOXSCORE_URL_SUFFIX)
	foxGameStatsHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_STATS_URL_SUFFIX)
	foxGameOddsHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_ODDS_URL_SUFFIX)

	boxscorePage := scrapeFoxGame(foxGameBoxscoreHyperlink)
	statsPage := scrapeFoxGame(foxGameStatsHyperlink)
	oddsPage := scrapeFoxGame(foxGameOddsHyperlink)

	return FoxGamePages{
		BoxscorePage: boxscorePage,
		StatsPage:    statsPage,
		OddsPage:     oddsPage,
	}
}*/

// scrapeGame fetches the Fox boxscore, stats, and odds pages for the configured game.
func (g FoxGame) scrapeGame() FoxGamePages {
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_BASE_URL, g.FoxSchedulePage)

	foxGameBoxscoreHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_BOXSCORE_URL_SUFFIX)
	foxGameStatsHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_STATS_URL_SUFFIX)
	foxGameOddsHyperlink := fmt.Sprint(foxGameHyperlink, utils.FOX_GAME_ODDS_URL_SUFFIX)

	boxscorePage := scrapeFoxGame(foxGameBoxscoreHyperlink)
	statsPage := scrapeFoxGame(foxGameStatsHyperlink)
	oddsPage := scrapeFoxGame(foxGameOddsHyperlink)

	return FoxGamePages{
		BoxscorePage: boxscorePage,
		StatsPage:    statsPage,
		OddsPage:     oddsPage,
	}
}
