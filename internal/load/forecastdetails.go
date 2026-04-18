package load

import (
	"have-a-nice-pickem-etl/internal/transform/forecastdetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// writes betting odds records to data/bettingodds.csv.
func loadForecastDetails(records []forecastdetails.ForecastDetails, db *gorm.DB) {
	if len(records) > 0 {
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "location_id"},
				{Name: "zulu_game_time"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"temperature",
				"feels_like",
				"humidity",
				"visibility",
				"wind_speed",
				"short_description",
				"long_description",
			}),
		}).Create(&records)
	}
}
