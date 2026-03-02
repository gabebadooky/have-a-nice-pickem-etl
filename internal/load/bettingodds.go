package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"log"
)

// BettingOdds writes betting odds records to data/bettingodds.csv.
func loadBettingOdds(records []bettingodds.BettingOdds, csvLoadFolderPath string) {
	bulkDataLoadFilePath := fmt.Sprintf("%s/%s", csvLoadFolderPath, "bettingodds.csv")
	f, w := instantiateCsvWriter(bulkDataLoadFilePath)
	defer f.Close()
	defer w.Flush()

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
	callBulkLoadProcedure(queryString)
}
