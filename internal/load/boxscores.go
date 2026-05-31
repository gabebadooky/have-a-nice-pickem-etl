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
			{Name: "game"},
			{Name: "team"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"quarter1",
			"quarter2",
			"quarter3",
			"quarter4",
			"overtime",
			"total",
			"updated_at",
		}),
	}).Create(&records)
}
