package team

import (
	foxschedule "have-a-nice-pickem-etl/etl/extract/schedule/fox"
	espnteam "have-a-nice-pickem-etl/etl/extract/team/espn"
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
	ESPN espnteam.TeamSummaryEndpoint
	Fox  foxschedule.FoxSchedule
}

func ConsolidateTeamInfo(t AllTeamInfo) Team {
	return t.allTeamInfo()
}

func (t AllCfbTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnCfbTeam{TeamCode: t.TeamCode})

	return Team{
		ESPN: espnTeam,
	}
}

func (t AllNflTeamInfo) allTeamInfo() Team {
	var espnTeam espnteam.TeamSummaryEndpoint = espnteam.GetTeamSummary(espnteam.EspnNflTeam{TeamCode: t.TeamCode})

	return Team{
		ESPN: espnTeam,
	}
}
