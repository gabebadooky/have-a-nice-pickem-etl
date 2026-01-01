package cbsteam

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsTeam interface {
	teamCode() string
}

type CbsAwayTeam struct {
	OddsPageTable *goquery.Selection
}

type CbsHomeTeam struct {
	OddsPageTable *goquery.Selection
}

func ExtractCbsTeamCode(t CbsTeam) string {
	return t.teamCode()
}

func extractCbsTeamCodeFromTeamHREF(teamHREF string) string {
	_, after, _ := strings.Cut(teamHREF, "teams/")
	teamCBScode := strings.TrimRight(after, "/")
	return teamCBScode
}

// Extracts team hyperlink in first "tr" tag in a given Odds Page Table goquery selection
func (t CbsAwayTeam) teamCode() string {
	const awayTrIndex int = 1
	teamHREF := t.OddsPageTable.Find("tbody").Find("tr").Eq(awayTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	cbsTeamCode := extractCbsTeamCodeFromTeamHREF(teamHREF)
	return cbsTeamCode
}

// Extracts team hyperlink in second "tr" tag in a given Odds Page Table goquery selection
func (t CbsHomeTeam) teamCode() string {
	const homeTrIndex int = 0
	teamHREF := t.OddsPageTable.Find("tbody").Find("tr").Eq(homeTrIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	cbsTeamCode := extractCbsTeamCodeFromTeamHREF(teamHREF)
	return cbsTeamCode
}
