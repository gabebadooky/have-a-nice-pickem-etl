// Package cbsteam provides CBS Sports team stats page web scraping functionality.
// It extracts team statistics by locating team pages from the CBS Sports teams directory
// and scraping team stats pages for both college football (CFB) and NFL.
package cbsteam

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
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

// setTeamPageHyperlink finds the CBS team stats page URL for the given team ID from the teams directory page.
func setTeamPageHyperlink(teamsPageHyperlink string, teamID string) string {
	allTeamsPage := scrapePage(teamsPageHyperlink)
	teamHyperlinks := compileAllTeamHyperlinks(*allTeamsPage)
	teamHyperlink := locateTeamPageHyperLink(teamHyperlinks, teamID)
	return teamHyperlink
}

// scrapeTeamPage scrapes the CBS college football team stats page for the configured team.
func (c CbsCfbTeam) GetTeamPage() *goquery.Selection {
	teamPageHyperlink := setTeamPageHyperlink(utils.CBS_CFB_ALL_TEAMS_PAGE_URL, c.TeamID)
	if teamPageHyperlink == "" {
		return nil
	}

	teamStatsHyperlink := fmt.Sprintf("%s%s/%s", utils.CBS_BASE_URL, teamPageHyperlink, utils.CBS_TEAM_STATS_URL_SUFFIX)
	teamStatsPage := scrapePage(teamStatsHyperlink)
	return teamStatsPage
}

// scrapeTeamPage scrapes the CBS NFL team stats page for the configured team.
func (n CbsNflTeam) GetTeamPage() *goquery.Selection {
	teamPageHyperlink := setTeamPageHyperlink(utils.CBS_NFL_ALL_TEAMS_PAGE_URL, n.TeamID)
	if teamPageHyperlink == "" {
		return nil
	}

	teamStatsHyperlink := fmt.Sprintf("%s%s/%s", utils.CBS_BASE_URL, teamPageHyperlink, utils.CBS_TEAM_STATS_URL_SUFFIX)
	teamStatsPage := scrapePage(teamStatsHyperlink)
	return teamStatsPage
}
