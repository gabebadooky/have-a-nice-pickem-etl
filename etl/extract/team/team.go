package team

import (
	"fmt"
	cbsteam "have-a-nice-pickem-etl/etl/extract/team/cbs"
	espnteam "have-a-nice-pickem-etl/etl/extract/team/espn"
	"have-a-nice-pickem-etl/etl/utils"

	"github.com/PuerkitoBio/goquery"
)

type AllTeamInfo interface {
	allTeamInfo() Team
}

type AllCfbTeamInfo struct {
	EspnCode string
}

type AllNflTeamInfo struct {
	EspnCode string
}

type Team struct {
	TeamID string
	League string
	ESPN   espnteam.TeamSummaryEndpoint
	CBS    *goquery.Selection
}

func ConsolidateTeamInfo(t AllTeamInfo) Team {
	return t.allTeamInfo()
}

func (t AllCfbTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnCfbTeam{TeamCode: t.EspnCode})
	var teamLocationName string = fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	var teamID string = utils.FormatStringID(teamLocationName)
	var cbsTeamStats *goquery.Selection = cbsteam.GetTeamStatsPage(cbsteam.CbsCfbTeam{TeamID: teamID})

	return Team{
		TeamID: teamID,
		League: "CFB",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}

func (t AllNflTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnNflTeam{TeamCode: t.EspnCode})
	var teamLocationName string = fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	var teamID string = utils.FormatStringID(teamLocationName)
	var cbsTeamStats *goquery.Selection = cbsteam.GetTeamStatsPage(cbsteam.CbsNflTeam{TeamID: teamID})

	return Team{
		TeamID: teamID,
		League: "NFL",
		ESPN:   espnTeam,
		CBS:    cbsTeamStats,
	}
}
