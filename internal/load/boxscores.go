package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"log"
)

// Boxscores writes boxscore records to data/boxscores.csv.
func loadBoxscores(records []boxscore.Boxscore, csvLoadFolderPath string) {
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
}
