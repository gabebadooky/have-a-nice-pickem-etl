package load

import (
	"have-a-nice-pickem-etl/internal/transform/bettingodds"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BettingOdds writes betting odds records to data/bettingodds.csv.
func loadBettingOdds(records []bettingodds.BettingOdds, db *gorm.DB) {
	if len(records) > 0 {
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "game_id"},
				{Name: "team_id"},
				{Name: "source"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"over_under",
				"moneyline",
				"spread",
				"win_probability",
				"updated_at",
			}),
		}).Create(&records)
	}
}
