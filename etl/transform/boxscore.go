package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/boxscore"
	"have-a-nice-pickem-etl/etl/transform/common"
)

// Instantiates Box Score record from ESPN Game Summary
func CreateBoxScoreRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, homeAway string) pickemstructs.Boxscore {
	var newRecord pickemstructs.Boxscore = pickemstructs.Boxscore{
		GameID:        common.ParseGameID(espnGameDetails),
		TeamID:        common.ParseGameSummaryTeamID(homeAway, espnGameDetails),
		Q1Score:       boxscore.ParseQuarterScore(espnGameDetails, homeAway, 1),
		Q2Score:       boxscore.ParseQuarterScore(espnGameDetails, homeAway, 2),
		Q3Score:       boxscore.ParseQuarterScore(espnGameDetails, homeAway, 3),
		Q4Score:       boxscore.ParseQuarterScore(espnGameDetails, homeAway, 4),
		OvertimeScore: boxscore.ParseOvertimeScore(espnGameDetails, homeAway),
		TotalScore:    boxscore.ParseTotalScore(espnGameDetails, homeAway),
	}

	return newRecord
}
