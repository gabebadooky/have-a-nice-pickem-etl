// Package cbsgame provides helper functions for extracting team codes from CBS game pages.
// These functions parse team identifiers from CBS Sports HTML structures for matching
// games with team IDs.
package cbsgame

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type awayTeam struct {
	oddsPageTable *goquery.Selection
}

type homeTeam struct {
	oddsPageTable *goquery.Selection
}

// extractCbsTeamCodeFromTeamHREF parses the CBS team code from a team profile URL path.
func extractCbsTeamCodeFromTeamHREF(teamHREF string) string {
	_, after, _ := strings.Cut(teamHREF, "teams/")
	teamCBScode := strings.TrimRight(after, "/")
	return teamCBScode
}

// scrapeTeamCode returns the CBS team code for the away team from the odds block row.
func (a awayTeam) scrapeTeamCode() string {
	const awayTrIndex int = 1
	teamHREF := a.oddsPageTable.Find("tbody").Find("tr").Eq(awayTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return extractCbsTeamCodeFromTeamHREF(teamHREF)

}

// scrapeTeamCode returns the CBS team code for the home team from the odds block row.
func (h homeTeam) scrapeTeamCode() string {
	const homeTrIndex int = 0
	teamHREF := h.oddsPageTable.Find("tbody").Find("tr").Eq(homeTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return extractCbsTeamCodeFromTeamHREF(teamHREF)

}
