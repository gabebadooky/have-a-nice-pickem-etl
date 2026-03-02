package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
	"log"
)

// TeamDetails writes team detail records to data/teams.csv.
func loadTeamDetails(records []teamdetails.TeamDetails, csvLoadFolderPath string) {
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
}
