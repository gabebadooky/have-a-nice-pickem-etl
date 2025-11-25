package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/teamdetails"
	"have-a-nice-pickem-etl/etl/utils"
)

type TeamDetails struct {
	TeamID         string
	League         string
	ESPNCode       string
	CBSCode        string
	FoxCode        string
	VegasCode      string
	ConferenceID   string
	Name           string
	Mascot         string
	PrimaryColor   string
	AlternateColor string
	Ranking        uint8
}

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
func CreateTeamDetailsRecord(espnTeamDetails pickemstructs.TeamSummaryResponse, league string) TeamDetails {
	var newRecord TeamDetails
	var teamID string = common.ParseTeamSummaryTeamID(espnTeamDetails)

	newRecord.TeamID = teamID
	newRecord.League = league
	newRecord.ESPNCode = teamdetails.ParseESPNteamCode(espnTeamDetails)
	newRecord.CBSCode = setCbsCode(teamID)
	newRecord.FoxCode = setFoxCode(teamID)
	newRecord.VegasCode = ""
	newRecord.ConferenceID = teamdetails.ParseConferenceID(espnTeamDetails)
	newRecord.Name = teamdetails.ParseTeamName(espnTeamDetails)
	newRecord.Mascot = teamdetails.ParseTeamMascot(espnTeamDetails)
	newRecord.PrimaryColor = teamdetails.ParsePrimaryColor(espnTeamDetails)
	newRecord.AlternateColor = teamdetails.ParseAlternateColor(espnTeamDetails)
	newRecord.Ranking = teamdetails.ParseRanking(espnTeamDetails)

	return newRecord
}
