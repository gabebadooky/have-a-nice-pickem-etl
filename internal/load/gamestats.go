package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"log"
)

// Stats writes game stats records to data/stats.csv.
func loadGameStats(records []gamestats.GameStats, csvLoadFolderPath string) {
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
}
