package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/record"
)

type Record struct {
	TeamID     string
	RecordType string
	Wins       uint8
	Losses     uint8
	Ties       uint8
}

// Instantiates Team Overall Record from ESPN Team Summary
func CreateOverallRecordRow(espnTeamDetails pickemstructs.TeamSummaryResponse) Record {
	var newRecord Record

	newRecord.TeamID = common.ParseTeamSummaryTeamID(espnTeamDetails)
	newRecord.RecordType = "Overall"
	newRecord.Wins = record.ParseOverallWins(espnTeamDetails)
	newRecord.Losses = record.ParseOverallLosses(espnTeamDetails)
	newRecord.Ties = record.ParseOverallTies(espnTeamDetails)

	return newRecord
}

/* Instantiates Team Overall Record from ESPN Team Summary
func CreateConferenceRecordRow(espnTeamDetails pickemstructs.TeamSummaryResponse) Record {
	var newRecord Record

	newRecord.TeamID = common.ParseTeamSummaryTeamID(espnTeamDetails)
	newRecord.RecordType = "Conference"
	newRecord.Wins = record.ParseOverallWins(espnTeamDetails)
	newRecord.Losses = record.ParseOverallLosses(espnTeamDetails)
	newRecord.Ties = record.ParseOverallTies(espnTeamDetails)

	return newRecord
}
*/
