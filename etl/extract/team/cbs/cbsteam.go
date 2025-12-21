package cbs

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsTeam interface {
	ExtractCbsTeamCode() string
}

type CbsAwayTeam struct {
	OddsPageTable *goquery.Selection
}

type CbsHomeTeam struct {
	OddsPageTable *goquery.Selection
}

func ExtractCbsTeamCodeFromTeamHREF(teamHREF string) string {
	var teamCBScode string
	var teamCBScodeIndex int = strings.Index(teamHREF, "teams/")
	teamCBScode = teamHREF[teamCBScodeIndex+6:]
	teamCBScode = strings.TrimRight(teamCBScode, "/")
	return teamCBScode
}

func (a CbsAwayTeam) ExtractTeamCode() string {
	const trIndex int = 1
	var teamHREF string = a.OddsPageTable.Find("tbody").Find("tr").Eq(trIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return ExtractCbsTeamCodeFromTeamHREF(teamHREF)
}

func (a CbsHomeTeam) ExtractTeamCode() string {
	const trIndex int = 0
	var teamHREF string = a.OddsPageTable.Find("tbody").Find("tr").Eq(trIndex).Find("span.OddsBlock-teamText").Find("a").AttrOr("href", "cbsTeamHREF")
	return ExtractCbsTeamCodeFromTeamHREF(teamHREF)
}
