package teamdetails

import "have-a-nice-pickem-etl/etl/pickemstructs"

func ParseTeamID(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var teamID string = teamSummaryDetails.Team.ID
	return teamID
}

func ParseESPNteamCode(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var espnTeamCode string = teamSummaryDetails.Team.Code
	return espnTeamCode
}

func ParseConferenceID(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var conferenceID string = teamSummaryDetails.Team.Groups.ID
	return conferenceID
}

func ParseTeamName(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var teamName string = teamSummaryDetails.Team.Location
	return teamName
}

func ParseTeamMascot(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var teamMascot string = teamSummaryDetails.Team.Name
	return teamMascot
}

func ParsePrimaryColor(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var primaryColor string = teamSummaryDetails.Team.PrimaryColor
	return primaryColor
}

func ParseAlternateColor(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var alternateColor string = teamSummaryDetails.Team.AlternateColor
	return alternateColor
}

func ParseRanking(teamSummaryDetails pickemstructs.TeamSummaryResponse) uint8 {
	var ranking uint8 = teamSummaryDetails.Team.Ranking
	return ranking
}
