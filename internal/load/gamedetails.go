package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"log"
)

// GameDetails writes game detail records to data/gamedetails.csv.
func loadGameDetails(records []gamedetails.GameDetails, csvLoadFolderPath string) {
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
}
