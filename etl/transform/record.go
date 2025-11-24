package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/record"
)

func CreateOverallRecordRow(espnTeamDetails pickemstructs.TeamSummaryResponse) pickemstructs.Record {
	var newRecord pickemstructs.Record

	newRecord.TeamID = common.ParseTeamSummaryTeamID(espnTeamDetails)
	newRecord.RecordType = "Overall"
	newRecord.Wins = record.ParseOverallWins(espnTeamDetails)
	newRecord.Losses = record.ParseOverallLosses(espnTeamDetails)
	newRecord.Ties = record.ParseOverallTies(espnTeamDetails)

	return newRecord
}
