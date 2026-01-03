package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/boxscore"
	"have-a-nice-pickem-etl/etl/transform/common"
)

// Instantiates Box Score record from ESPN Game Summary
func CreateBoxScoreRecord(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties, homeAway string) pickemstructs.Boxscore {
	var newRecord pickemstructs.Boxscore = pickemstructs.Boxscore{
		GameID:        common.ParseGameID(consolidatedGameProperties),
		TeamID:        common.ParseGameSummaryTeamID(homeAway, consolidatedGameProperties),
		Q1Score:       boxscore.ParseQuarterScore(consolidatedGameProperties, homeAway, 1),
		Q2Score:       boxscore.ParseQuarterScore(consolidatedGameProperties, homeAway, 2),
		Q3Score:       boxscore.ParseQuarterScore(consolidatedGameProperties, homeAway, 3),
		Q4Score:       boxscore.ParseQuarterScore(consolidatedGameProperties, homeAway, 4),
		OvertimeScore: boxscore.ParseOvertimeScore(consolidatedGameProperties, homeAway),
		TotalScore:    boxscore.ParseTotalScore(consolidatedGameProperties, homeAway),
	}

	return newRecord
}
