package load

import (
	"have-a-nice-pickem-etl/internal/transform/gamedetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GameDetails writes game detail records to data/gamedetails.csv.
func loadGameDetails(records []gamedetails.GameDetails, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"league",
			"week",
			"season",
			"espn_code",
			"cbs_code",
			"fox_code",
			"vegas_code",
			"away_team",
			"home_team",
			"kickoff_utc",
			"broadcast",
			"location",
			"finished",
			"updated_at",
		}),
	}).Create(&records)
}
