package load

import (
	"have-a-nice-pickem-etl/internal/transform/bettingodds"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BettingOdds writes betting odds records to data/bettingodds.csv.
func loadBettingOdds(records []bettingodds.BettingOdds, db *gorm.DB) {
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

	/*db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "game_id"},
			{Name: "team_id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"source":          gorm.Expr("EXCLUDED.source"),
			"over_under":      gorm.Expr("EXCLUDED.over_under"),
			"moneyline":       gorm.Expr("EXCLUDED.moneyline"),
			"spread":          gorm.Expr("EXCLUDED.spread"),
			"win_probability": gorm.Expr("EXCLUDED.win_probability"),
			"updated_at":      gorm.Expr("CURRENT_TIMESTAMP"),
		}),
	}).Create(&records)*/

	/*var insertRows string
	for _, record := range records {
		newRow := fmt.Sprintf("('%s', '%s', '%s', '%f', '%d', '%f', '%f')", record.GameID, record.TeamID, record.Source, record.OverUnder, record.Moneyline, record.Spread, record.WinProbability)
		insertRows = fmt.Sprintf(insertRows, ", ", newRow)
	}

	bulkLoadSqlStatement := fmt.Sprintf("INSERT INTO betting_odds VALUES %s", insertRows)
	callBulkLoadProcedure(bulkLoadSqlStatement)*/

	/*
		bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "bettingodds.csv")
		f, w := instantiateCsvWriter(bulkDataLoadFilePath)
		defer f.Close()
		defer w.Flush()

		log.Printf("Writing Betting Odds header to %s", bulkDataLoadFilePath)
		w.Write([]string{"GameID", "TeamID", "Source", "OverUnder", "Moneyline", "Spread", "WinProbability"})

		log.Printf("Writing Betting Odds records to %s", bulkDataLoadFilePath)

		for _, record := range records {
			w.Write([]string{
				record.GameID,
				record.TeamID,
				record.Source,
				fmt.Sprintf("%f", record.OverUnder),
				fmt.Sprintf("%d", record.Moneyline),
				fmt.Sprintf("%f", record.Spread),
				fmt.Sprintf("%f", record.WinProbability),
			})
		}

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		queryString := fmt.Sprintf("CALL %s('%s')", "bulk_load_betting_odds", bulkDataLoadFilePath)
		callBulkLoadProcedure(queryString) */
}
