package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/teamdetails"
)

func CreateTeamDetailsRecord(espnTeamDetails pickemstructs.TeamSummaryResponse, league string) pickemstructs.TeamDetails {
	var newRecord pickemstructs.TeamDetails

	newRecord.TeamID = teamdetails.ParseTeamID(espnTeamDetails)
	newRecord.League = league
	newRecord.ESPNCode = teamdetails.ParseESPNteamCode(espnTeamDetails)
	newRecord.CBSCode = ""
	newRecord.FoxCode = ""
	newRecord.VegasCode = ""
	newRecord.ConferenceID = teamdetails.ParseConferenceID(espnTeamDetails)
	newRecord.Name = teamdetails.ParseTeamName(espnTeamDetails)
	newRecord.Mascot = teamdetails.ParseTeamMascot(espnTeamDetails)
	newRecord.PrimaryColor = teamdetails.ParsePrimaryColor(espnTeamDetails)
	newRecord.AlternateColor = teamdetails.ParseAlternateColor(espnTeamDetails)
	newRecord.Ranking = teamdetails.ParseRanking(espnTeamDetails)

	return newRecord
}
