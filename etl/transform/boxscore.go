package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/boxscore"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type Boxscore struct {
	GameID        string
	TeamID        string
	Q1Score       uint8
	Q2Score       uint8
	Q3Score       uint8
	Q4Score       uint8
	OvertimeScore uint8
	TotalScore    uint8
}

// Instantiates Box Score record from ESPN Game Summary
func CreateBoxScoreRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, homeAway string) Boxscore {
	var newRecord Boxscore

	newRecord.GameID = common.ParseGameID(espnGameDetails)
	newRecord.TeamID = common.ParseGameSummaryTeamID(homeAway, espnGameDetails)
	newRecord.Q1Score = boxscore.ParseQuarterScore(espnGameDetails, homeAway, 1)
	newRecord.Q2Score = boxscore.ParseQuarterScore(espnGameDetails, homeAway, 2)
	newRecord.Q3Score = boxscore.ParseQuarterScore(espnGameDetails, homeAway, 3)
	newRecord.Q4Score = boxscore.ParseQuarterScore(espnGameDetails, homeAway, 4)
	newRecord.OvertimeScore = boxscore.ParseOvertimeScore(espnGameDetails, homeAway)
	newRecord.TotalScore = boxscore.ParseTotalScore(espnGameDetails, homeAway)

	return newRecord
}
