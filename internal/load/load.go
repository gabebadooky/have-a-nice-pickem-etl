// Package load provides the loading layer for the ETL pipeline.
// This package handles writing transformed data to the target data store.
package load

import (
	"encoding/csv"
	"fmt"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"have-a-nice-pickem-etl/internal/transform/locationdetails"
	"have-a-nice-pickem-etl/internal/transform/record"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
	"log"
	"os"
)

// instantiateCsvWriter creates and returns a CSV writer for the given file path.
func instantiateCsvWriter(filepath string) *csv.Writer {
	f, err := os.OpenFile(filepath, os.O_CREATE, 0o644)
	if err != nil {
		log.Fatalf("Error occurred instantiating CSV Writer to %s", filepath)
	}
	defer f.Close()

	csvwriter := csv.NewWriter(f)
	return csvwriter
}

// BettingOdds writes betting odds records to data/bettingodds.csv.
func BettingOdds(records []bettingodds.BettingOdds) {
	bulkDataLoadFilePath := "./data/bettingodds.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

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
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// Boxscores writes boxscore records to data/boxscores.csv.
func Boxscores(records []boxscore.Boxscore) {
	bulkDataLoadFilePath := "./data/boxscores.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

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
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// GameDetails writes game detail records to data/gamedetails.csv.
func GameDetails(records []gamedetails.GameDetails) {
	bulkDataLoadFilePath := "./data/gamedetails.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

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
			record.Location,
			fmt.Sprintf("%t", record.Finished),
		})
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// Stats writes game stats records to data/stats.csv.
func Stats(records []gamestats.GameStats) {
	bulkDataLoadFilePath := "./data/stats.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

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
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// LocationDetails writes location detail records to data/locations.csv.
func LocationDetails(records []locationdetails.LocationDetails) {
	bulkDataLoadFilePath := "./data/locations.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

	for _, record := range records {
		w.Write([]string{
			record.LocationID,
			record.Stadium,
			record.City,
			record.State,
			fmt.Sprintf("%f", record.Latitude),
			fmt.Sprintf("%f", record.Longitude),
		})
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// TeamRecord writes team record rows to data/teamrecords.csv.
func TeamRecord(records []record.Record) {
	bulkDataLoadFilePath := "./data/teamrecords.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

	for _, record := range records {
		w.Write([]string{
			record.TeamID,
			record.RecordType,
			fmt.Sprintf("%d", record.Wins),
			fmt.Sprintf("%d", record.Losses),
			fmt.Sprintf("%d", record.Ties),
		})
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

// TeamDetails writes team detail records to data/teams.csv.
func TeamDetails(records []teamdetails.TeamDetails) {
	bulkDataLoadFilePath := "./data/teams.csv"
	w := instantiateCsvWriter(bulkDataLoadFilePath)

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
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
