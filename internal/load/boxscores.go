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
			{Name: "game_id"},
			{Name: "team_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"q1_score",
			"q2_score",
			"q3_score",
			"q4_score",
			"overtime",
			"total",
			"updated_at",
		}),
	}).Create(&records)

	/*

		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "game_id"},
				{Name: "team_id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"q1_score":   gorm.Expr("EXCLUDED.q1_score"),
				"q2_score":   gorm.Expr("EXCLUDED.q2_score"),
				"q3_score":   gorm.Expr("EXCLUDED.q3_score"),
				"q4_score":   gorm.Expr("EXCLUDED.q4_score"),
				"overtime":   gorm.Expr("EXCLUDED.overtime"),
				"total":      gorm.Expr("EXCLUDED.total"),
				"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
			}),
		}).Create(&records)

			bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "boxscores.csv")
			f, w := instantiateCsvWriter(bulkDataLoadFilePath)
			defer f.Close()
			defer w.Flush()

			log.Printf("Writing Boxscores records to %s", bulkDataLoadFilePath)

			for _, record := range records {
				w.Write([]string{
					record.GameID,
					record.TeamID,
					fmt.Sprintf("%d", record.Q1Score),
					fmt.Sprintf("%d", record.Q2Score),
					fmt.Sprintf("%d", record.Q3Score),
					fmt.Sprintf("%d", record.Q4Score),
					fmt.Sprintf("%d", record.OvertimeScore),
					fmt.Sprintf("%d", record.TotalScore),
				})
			}

			if err := w.Error(); err != nil {
				log.Fatal(err)
			}

			queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_box_scores", bulkDataLoadFilePath)
			callBulkLoadProcedure(queryString)
	*/
}
