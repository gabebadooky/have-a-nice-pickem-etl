package load

import (
	"have-a-nice-pickem-etl/internal/transform/boxscore"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Boxscores writes boxscore records to data/boxscores.csv.
func loadBoxscores(records []boxscore.Boxscore, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "game_id"},
			{Name: "team_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"q1_score",
			"q2_score",
			"q3_score",
			"q4_score",
			"overtime",
			"total",
			"updated_at",
		}),
	}).Create(&records)
}
