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
		DoUpdates: clause.Assignments(map[string]interface{}{
			"league":         gorm.Expr("EXCLUDED.league"),
			"weeknum":        gorm.Expr("EXCLUDED.weeknum"),
			"season":         gorm.Expr("EXCLUDED.season"),
			"espn_code":      gorm.Expr("EXCLUDED.espn_code"),
			"cbs_code":       gorm.Expr("EXCLUDED.cbs_code"),
			"fox_code":       gorm.Expr("EXCLUDED.fox_code"),
			"vegas_code":     gorm.Expr("EXCLUDED.vegas_code"),
			"away_team_id":   gorm.Expr("EXCLUDED.away_team_id"),
			"home_team_id":   gorm.Expr("EXCLUDED.home_team_id"),
			"zulu_game_time": gorm.Expr("EXCLUDED.zulu_game_time"),
			"broadcast":      gorm.Expr("EXCLUDED.broadcast"),
			"location_id":    gorm.Expr("EXCLUDED.location_id"),
			"finished":       gorm.Expr("EXCLUDED.finished"),
			"updated_at":     gorm.Expr("CURRENT_TIMESTAMP"),
		}),
	}).Create(&records)

	/*
		bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "gamedetails.csv")
		f, w := instantiateCsvWriter(bulkDataLoadFilePath)
		defer f.Close()
		defer w.Flush()

		log.Printf("Writing Game Details records to %s", bulkDataLoadFilePath)

		for _, record := range records {
			w.Write([]string{
				record.GameID,
				record.League,
				fmt.Sprintf("%d", record.Week),
				fmt.Sprintf("%d", record.Year),
				record.EspnCode,
				record.CbsCode,
				record.FoxCode,
				record.VegasCode,
				record.AwayTeamID,
				record.HomeTeamID,
				record.ZuluTimestamp,
				record.Broadcast,
				record.LocationID,
				fmt.Sprintf("%t", record.Finished),
			})
		}

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_games", bulkDataLoadFilePath)
		callBulkLoadProcedure(queryString)
	*/
}
