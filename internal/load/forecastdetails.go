package load

import (
	"have-a-nice-pickem-etl/internal/transform/forecastdetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func loadForecastDetails(records []forecastdetails.ForecastDetails, db *gorm.DB) {
	if len(records) > 0 {
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "location_id"},
				{Name: "zulu_timestamp"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"temperature",
				"feels_like",
				"humidity",
				"visibility",
				"wind_speed",
				"short_description",
				"long_description",
				"updated_at",
			}),
		}).Create(&records)
	}
}
