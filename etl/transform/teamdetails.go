package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/teamdetails"
	"have-a-nice-pickem-etl/etl/utils"
)

func setCbsCode(teamID string) string {
	cbsCode, cbsMappingExists := utils.TeamIDtoCbsTeamCode[teamID]
	if cbsMappingExists {
		return cbsCode
	} else {
		return teamID
	}
}

func setFoxCode(teamID string) string {
	foxCode, foxMappingExists := utils.TeamIDtoFoxTeamCode[teamID]
	if foxMappingExists {
		return foxCode
	} else {
		return teamID
	}
}

// Instantiates Team Details Record from ESPN Team Summary
func CreateTeamDetailsRecord(espnTeamDetails pickemstructs.TeamSummaryResponse, league string) pickemstructs.TeamDetails {
	var teamID string = common.ParseTeamSummaryTeamID(espnTeamDetails)

	var newRecord pickemstructs.TeamDetails = pickemstructs.TeamDetails{
		TeamID:         teamID,
		League:         league,
		ESPNCode:       teamdetails.ParseESPNteamCode(espnTeamDetails),
		CBSCode:        setCbsCode(teamID),
		FoxCode:        setFoxCode(teamID),
		VegasCode:      "",
		ConferenceID:   teamdetails.ParseConferenceID(espnTeamDetails),
		Name:           teamdetails.ParseTeamName(espnTeamDetails),
		Mascot:         teamdetails.ParseTeamMascot(espnTeamDetails),
		PrimaryColor:   teamdetails.ParsePrimaryColor(espnTeamDetails),
		AlternateColor: teamdetails.ParseAlternateColor(espnTeamDetails),
		Ranking:        teamdetails.ParseRanking(espnTeamDetails),
	}

	return newRecord
}
