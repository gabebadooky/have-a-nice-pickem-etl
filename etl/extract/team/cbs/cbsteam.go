package cbsteam

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsTeam interface {
	teamPage() *goquery.Selection
}

type CbsCfbTeam struct {
	TeamID string
}

type CbsNflTeam struct {
	TeamID string
}

func GetTeamStatsPage(t CbsTeam) *goquery.Selection {
	return t.teamPage()
}

func scrapePage(teamsPageHyperlink string) *goquery.Selection {
	page, err := utils.GetGoQuerySelectionBody(teamsPageHyperlink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

func compileAllTeamHyperlinks(teamsPageSelection goquery.Selection) []string {
	var allHyperlinks []string

	teamsPageSelection.Find("div.TeamLogoNameLockup-logo").Each(func(i int, logoDiv *goquery.Selection) {
		teamStatsHyperlink := logoDiv.Find("a").AttrOr("href", "teamStatsHyperlink")
		allHyperlinks = append(allHyperlinks, teamStatsHyperlink)
	})

	return allHyperlinks
}

func getTeamID(teamCode string) string {
	cbsCode, cbsMappingExists := utils.CbsTeamCodeToTeamIDmapping[teamCode]
	if cbsMappingExists {
		return cbsCode
	}
	return teamCode
}

func locateTeamStatsLink(allHyperlinks []string, teamID string) string {
	var teamStatsHyperlink string

	for i := range allHyperlinks {
		currentHyperlink := allHyperlinks[i]
		//cbsTeamCodeWithoutAbbr := strings.Split(currentHyperlink, "/")[4]
		mappedCbsCode := utils.GetCbsTeamCode(teamID)
		if strings.Contains(currentHyperlink, mappedCbsCode) {
			teamStatsHyperlink = currentHyperlink
			break
		}
	}

	return teamStatsHyperlink
}

func locateAndScrapeTeamStatsPage(teamsPageHyperlink string, teamID string) *goquery.Selection {
	log.Printf("\nRequesting All CBS Teams page: %s\n", teamsPageHyperlink)

	allTeamsPage := scrapePage(teamsPageHyperlink)
	teamHyperlinks := compileAllTeamHyperlinks(*allTeamsPage)
	teamStatsHyperlink := locateTeamStatsLink(teamHyperlinks, teamID)
	fullTeamStatsHyperlink := fmt.Sprintf("%s%s%s", utils.CBS_BASE_URL, teamStatsHyperlink, "stats/")

	log.Printf("\nRequesting CBS Teams Stats page: %s\n", fullTeamStatsHyperlink)
	teamStatsPage := scrapePage(fullTeamStatsHyperlink)
	return teamStatsPage
}

func (c CbsCfbTeam) teamPage() *goquery.Selection {
	var teamsPageHyperlink string = utils.CBS_CFB_ALL_TEAMS_PAGE_URL
	teamStatsPage := locateAndScrapeTeamStatsPage(teamsPageHyperlink, c.TeamID)
	return teamStatsPage
}

func (n CbsNflTeam) teamPage() *goquery.Selection {
	var teamsPageHyperlink string = utils.CBS_NFL_ALL_TEAMS_PAGE_URL
	teamStatsPage := locateAndScrapeTeamStatsPage(teamsPageHyperlink, n.TeamID)
	return teamStatsPage
}
