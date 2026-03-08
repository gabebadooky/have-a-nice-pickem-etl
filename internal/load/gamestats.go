package load

import (
	"have-a-nice-pickem-etl/internal/transform/gamestats"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Stats writes game stats records to data/stats.csv.
func loadGameStats(records []gamestats.GameStats, db *gorm.DB) {
	for i := range records {
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "game_id"},
				{Name: "team_id"},
				{Name: "stat_type"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"stat_value": gorm.Expr("EXCLUDED.stat_value"),
				"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
			}),
		}).Create(&records[i].Stats)
	}

	/*
		bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "stats.csv")
		f, w := instantiateCsvWriter(bulkDataLoadFilePath)
		defer f.Close()
		defer w.Flush()

		log.Printf("Writing Stats records to %s", bulkDataLoadFilePath)

		for _, record := range records {
			var gameID string = record.GameID
			var teamID string = record.TeamID

			for _, stat := range record.Stats {
				w.Write([]string{
					gameID,
					teamID,
					stat.StatType,
					fmt.Sprintf("%f", stat.Value),
				})
			}

		}

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_locations", bulkDataLoadFilePath)
		callBulkLoadProcedure(queryString)
	*/
}
