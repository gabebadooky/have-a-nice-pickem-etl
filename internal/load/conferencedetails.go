package load

import (
	"have-a-nice-pickem-etl/internal/transform/conferencedetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TeamDetails writes team detail records to data/teams.csv.
func loadConferenceDetails(records []conferencedetails.ConferenceDetails, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"abbreviation",
			"power_conference",
			"updated_at",
		}),
	}).Create(&records)
}
