package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/record"
)

// Instantiates Team Overall Record from ESPN Team Summary
func CreateOverallRecordRow(espnTeamDetails pickemstructs.TeamSummaryResponse) pickemstructs.Record {
	var newRecord pickemstructs.Record = pickemstructs.Record{
		TeamID:     common.ParseTeamSummaryTeamID(espnTeamDetails),
		RecordType: "Overall",
		Wins:       record.ParseOverallWins(espnTeamDetails),
		Losses:     record.ParseOverallLosses(espnTeamDetails),
		Ties:       record.ParseOverallTies(espnTeamDetails),
	}

	return newRecord
}
