package load

import (
	"have-a-nice-pickem-etl/internal/transform/record"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TeamRecord writes team record rows to data/teamrecords.csv.
func loadTeamRecord(records []record.Record, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "team_id"},
			{Name: "record_type"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"wins",
			"losses",
			"ties",
			"updated_at",
		}),
	}).Create(&records)

	/*
		bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "teamrecords.csv")
		f, w := instantiateCsvWriter(bulkDataLoadFilePath)
		defer f.Close()
		defer w.Flush()

		log.Printf("Writing Team Record rows to %s", bulkDataLoadFilePath)

		for _, record := range records {
			w.Write([]string{
				record.TeamID,
				record.RecordType,
				fmt.Sprintf("%d", record.Wins),
				fmt.Sprintf("%d", record.Losses),
				fmt.Sprintf("%d", record.Ties),
			})
		}

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_team_records", bulkDataLoadFilePath)
		callBulkLoadProcedure(queryString)
	*/
}
