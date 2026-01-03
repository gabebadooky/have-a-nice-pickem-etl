package team

import (
	"fmt"
	espnteam "have-a-nice-pickem-etl/etl/extract/team/espn"
	"have-a-nice-pickem-etl/etl/utils"
)

type AllTeamInfo interface {
	allTeamInfo() Team
}

type AllCfbTeamInfo struct {
	TeamCode string
}

type AllNflTeamInfo struct {
	TeamCode string
}

type Team struct {
	TeamID string
	League string
	ESPN   espnteam.TeamSummaryEndpoint
}

func ConsolidateTeamInfo(t AllTeamInfo) Team {
	return t.allTeamInfo()
}

func (t AllCfbTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnCfbTeam{TeamCode: t.TeamCode})
	var teamLocationName string = fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	var teamID string = utils.FormatStringID(teamLocationName)

	return Team{
		TeamID: teamID,
		League: "CFB",
		ESPN:   espnTeam,
	}
}

func (t AllNflTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnNflTeam{TeamCode: t.TeamCode})
	var teamLocationName string = fmt.Sprintf("%s %s", espnTeam.Team.Location, espnTeam.Team.Name)
	var teamID string = utils.FormatStringID(teamLocationName)

	return Team{
		TeamID: teamID,
		League: "NFL",
		ESPN:   espnTeam,
	}
}
