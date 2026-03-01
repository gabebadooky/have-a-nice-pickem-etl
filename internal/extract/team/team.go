// Package team provides team extraction functionality that consolidates team data
// from multiple sources (ESPN API, CBS web scraping) into a unified Team structure
// for both college football (CFB) and NFL.
package team

import (
	"fmt"
	cbsteam "have-a-nice-pickem-etl/internal/extract/team/cbs"
	espnteam "have-a-nice-pickem-etl/internal/extract/team/espn"
	"have-a-nice-pickem-etl/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

type CfbTeam struct {
	EspnCode string
}

type NflTeam struct {
	EspnCode string
}

type Team struct {
	TeamID string
	League string
	ESPN   espnteam.TeamSummaryEndpoint
	CBS    *goquery.Selection
}

// extractTeam fetches and consolidates college football team data from ESPN and CBS.
func (t CfbTeam) extractTeam() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.EspnCfbTeam{TeamCode: t.EspnCode}.GetTeamSummary()
	teamFullName := fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	teamID := utils.FormatStringID(teamFullName)
	var cbsTeamStats *goquery.Selection = cbsteam.CbsCfbTeam{TeamID: teamID}.GetTeamPage()

	return Team{
		TeamID: teamID,
		League: "CFB",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}

// extractTeam fetches and consolidates NFL team data from ESPN and CBS.
func (t NflTeam) extractTeam() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.EspnNflTeam{TeamCode: t.EspnCode}.GetTeamSummary()
	teamLocationName := fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	teamID := utils.FormatStringID(teamLocationName)
	var cbsTeamStats *goquery.Selection = cbsteam.CbsNflTeam{TeamID: teamID}.GetTeamPage()

	return Team{
		TeamID: teamID,
		League: "NFL",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}
