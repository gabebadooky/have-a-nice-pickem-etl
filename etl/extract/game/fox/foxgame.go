package foxgame

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
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

// Map Fox Team Code to global Team IDs
func getTeamID(foxTeamCode string) string {
	teamID, exists := utils.FoxTeamCodeToTeamIDmapping[foxTeamCode]
	if exists {
		return teamID
	} else {
		return foxTeamCode
	}
}

// Handle web scrape attempt for given Fox hyperlink
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

// Extracts FOX game code where AwayTeamID and HomeTeamID match with corresponding FOX team codes
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
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_GAME_BASE_URL, g.FoxSchedulePage)

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

func (g FoxGame) scrapeGame() FoxGamePages {
	foxGameHyperlink := scrapeGameHyperlink(g.GameID, utils.FOX_GAME_BASE_URL, g.FoxSchedulePage)

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
