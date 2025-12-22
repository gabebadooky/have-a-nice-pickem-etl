package team

import (
	"have-a-nice-pickem-etl/etl/extract/team/espn"
)

type AllTeamInfo interface {
}

type AllCfbTeamInfo struct {
	TeamCode string
}

type AllNflTeamInfo struct {
	TeamCode string
}

type Team struct {
	ESPN espn.TeamSummaryEndpoint
}

func (t AllCfbTeamInfo) ConsolidateTeamInfo() Team {
	var EspnTeam espn.TeamSummaryEndpoint = espn.EspnCfbTeam{TeamCode: t.TeamCode}.GetTeamSummary()

	return Team{
		ESPN: EspnTeam,
	}
}

func (t AllNflTeamInfo) ConsolidateTeamInfo() Team {
	var EspnTeam espn.TeamSummaryEndpoint = espn.EspnNflTeam{TeamCode: t.TeamCode}.GetTeamSummary()

	return Team{
		ESPN: EspnTeam,
	}
}
