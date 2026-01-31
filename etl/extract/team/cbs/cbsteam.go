package cbsteam

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsCfbTeam struct {
	TeamID string
}

type CbsNflTeam struct {
	TeamID string
}
type cbsTeamInstantiator interface {
	scrapeTeamPage() *goquery.Selection
}

func GetTeamStatsPage(t cbsTeamInstantiator) *goquery.Selection {
	return t.scrapeTeamPage()
}

// Make and handle CBS Team page web scrape attempt
func scrapePage(teamsPageHyperlink string) *goquery.Selection {
	page, err := utils.GetGoQuerySelectionBody(teamsPageHyperlink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Retrieve all team page hyperlinks in current selection
func compileAllTeamHyperlinks(teamsPageSelection goquery.Selection) []string {
	var allHyperlinks []string

	teamsPageSelection.Find("div.TeamLogoNameLockup-logo").Each(func(i int, logoDiv *goquery.Selection) {
		teamStatsHyperlink := logoDiv.Find("a").AttrOr("href", "teamStatsHyperlink")
		allHyperlinks = append(allHyperlinks, teamStatsHyperlink)
	})

	return allHyperlinks
}

// Return hyperlink from allHyperlinks that contains `teamID` string
func locateTeamPageHyperLink(allHyperlinks []string, teamID string) string {
	var teamStatsHyperlink string

	for i := range allHyperlinks {
		currentHyperlink := allHyperlinks[i]
		mappedCbsCode := utils.GetCbsTeamCode(teamID)
		if strings.Contains(currentHyperlink, mappedCbsCode) {
			teamStatsHyperlink = currentHyperlink
			break
		}
	}

	return teamStatsHyperlink
}

func setTeamPageHyperlink(teamsPageHyperlink string, teamID string) string {
	log.Printf("\nRequesting All CBS Teams page: %s\n", teamsPageHyperlink)

	allTeamsPage := scrapePage(teamsPageHyperlink)
	teamHyperlinks := compileAllTeamHyperlinks(*allTeamsPage)
	teamHyperlink := locateTeamPageHyperLink(teamHyperlinks, teamID)
	return teamHyperlink
}

func (c CbsCfbTeam) scrapeTeamPage() *goquery.Selection {
	teamPageHyperlink := setTeamPageHyperlink(utils.CBS_CFB_ALL_TEAMS_PAGE_URL, c.TeamID)
	teamStatsHyperlink := fmt.Sprintf("%s%s%s", utils.CBS_BASE_URL, teamPageHyperlink, utils.CBS_TEAM_STATS_URL_SUFFIX)
	teamStatsPage := scrapePage(teamStatsHyperlink)
	return teamStatsPage
}

func (n CbsNflTeam) scrapeTeamPage() *goquery.Selection {
	teamPageHyperlink := setTeamPageHyperlink(utils.CBS_NFL_ALL_TEAMS_PAGE_URL, n.TeamID)
	teamStatsHyperlink := fmt.Sprintf("%s%s%s", utils.CBS_BASE_URL, teamPageHyperlink, utils.CBS_TEAM_STATS_URL_SUFFIX)
	teamStatsPage := scrapePage(teamStatsHyperlink)
	return teamStatsPage
}
