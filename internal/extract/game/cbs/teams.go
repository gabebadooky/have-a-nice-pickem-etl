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

func extractCbsTeamCodeFromTeamHREF(teamHREF string) string {
	_, after, _ := strings.Cut(teamHREF, "teams/")
	teamCBScode := strings.TrimRight(after, "/")
	return teamCBScode
}

// Extracts team hyperlink in first "tr" tag in a given Odds Page Table goquery selection
func (a awayTeam) scrapeTeamCode() string {
	const awayTrIndex int = 1
	teamHREF := a.oddsPageTable.Find("tbody").Find("tr").Eq(awayTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return extractCbsTeamCodeFromTeamHREF(teamHREF)

}

// Extracts team hyperlink in second "tr" tag in a given Odds Page Table goquery selection
func (h homeTeam) scrapeTeamCode() string {
	const homeTrIndex int = 0
	teamHREF := h.oddsPageTable.Find("tbody").Find("tr").Eq(homeTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return extractCbsTeamCodeFromTeamHREF(teamHREF)

}
