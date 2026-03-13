package load

import (
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type stats struct {
	GameID    string    `gorm:"column:game_id"`
	TeamID    string    `gorm:"column:team_id"`
	StatType  string    `gorm:"column:stat_type"`
	StatValue float32   `gorm:"column:stat_value"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (stats) TableName() string {
	return "pickem.stats"
}

func flattenStats(records []gamestats.GameStats) []stats {
	var flattenedStats []stats

	for i := range records {
		var gameID string = records[i].GameID
		var teamID string = records[i].TeamID
		var currentStatRow []gamestats.Stat = records[i].Stats

		for j := range currentStatRow {
			s := stats{
				GameID:    gameID,
				TeamID:    teamID,
				StatType:  currentStatRow[j].StatType,
				StatValue: currentStatRow[j].Value,
			}

			flattenedStats = append(flattenedStats, s)
		}
	}

	return flattenedStats
}

// Stats writes game stats records to data/stats.csv.
func loadGameStats(records []gamestats.GameStats, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "game_id"},
			{Name: "team_id"},
			{Name: "stat_type"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"stat_value",
			"updated_at",
		}),
	}).Create(flattenStats(records))

	/* for i := range records {
		db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "game_id"},
				{Name: "team_id"},
				{Name: "stat_type"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"stat_value",
				"updated_at",
			}),
		}).Create(&records[i].Stats)
	} */

	/*
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
