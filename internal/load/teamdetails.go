package load

import (
	"have-a-nice-pickem-etl/internal/transform/teamdetails"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TeamDetails writes team detail records to data/teams.csv.
func loadTeamDetails(records []teamdetails.TeamDetails, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"team_id":         gorm.Expr("EXCLUDEDteam_id"),
			"league":          gorm.Expr("EXCLUDED.league"),
			"espn_code":       gorm.Expr("EXCLUDED.espn_code"),
			"cbs_code":        gorm.Expr("EXCLUDED.cbs_code"),
			"fox_code":        gorm.Expr("EXCLUDED.fox_code"),
			"vegas_code":      gorm.Expr("EXCLUDED.vegas_code"),
			"conference_id":   gorm.Expr("EXCLUDED.conference_id"),
			"division_name":   gorm.Expr("EXCLUDED.division_name"),
			"team_name":       gorm.Expr("EXCLUDED.team_name"),
			"team_mascot":     gorm.Expr("EXCLUDED.team_mascot"),
			"primary_color":   gorm.Expr("EXCLUDED.primary_color"),
			"alternate_color": gorm.Expr("EXCLUDED.alternate_color"),
			"ranking":         gorm.Expr("EXCLUDED.ranking"),
			"updated_at":      gorm.Expr("CURRENT_TIMESTAMP"),
		}),
	}).Create(&records)

	/*
		bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "teams.csv")
		f, w := instantiateCsvWriter(bulkDataLoadFilePath)
		defer f.Close()
		defer w.Flush()

		log.Printf("Writing Team Details records to %s", bulkDataLoadFilePath)

		for _, record := range records {
			w.Write([]string{
				record.TeamID,
				record.League,
				record.ESPNCode,
				record.CBSCode,
				record.FoxCode,
				record.VegasCode,
				record.ConferenceID,
				record.Name,
				record.Mascot,
				record.PrimaryColor,
				record.AlternateColor,
				fmt.Sprintf("%d", record.Ranking),
			})
		}

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_teams", bulkDataLoadFilePath)
		callBulkLoadProcedure(queryString)
	*/
}
