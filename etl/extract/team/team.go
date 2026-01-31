package team

import (
	"fmt"
	cbsteam "have-a-nice-pickem-etl/etl/extract/team/cbs"
	espnteam "have-a-nice-pickem-etl/etl/extract/team/espn"
	"have-a-nice-pickem-etl/etl/utils"

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

type teamInstantiator interface {
	extractTeam() Team
}

func ConsolidateTeamInfo(t teamInstantiator) Team {
	return t.extractTeam()
}

func (t CfbTeam) extractTeam() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnCfbTeam{TeamCode: t.EspnCode})
	teamLocationName := fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	teamID := utils.FormatStringID(teamLocationName)
	var cbsTeamStats *goquery.Selection = cbsteam.GetTeamStatsPage(cbsteam.CbsCfbTeam{TeamID: teamID})

	return Team{
		TeamID: teamID,
		League: "CFB",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}

func (t NflTeam) extractTeam() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnNflTeam{TeamCode: t.EspnCode})
	teamLocationName := fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	teamID := utils.FormatStringID(teamLocationName)
	var cbsTeamStats *goquery.Selection = cbsteam.GetTeamStatsPage(cbsteam.CbsNflTeam{TeamID: teamID})

	return Team{
		TeamID: teamID,
		League: "NFL",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}
