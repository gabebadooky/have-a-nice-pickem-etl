// Package load provides the loading layer for the ETL pipeline.
// This package handles writing transformed data to the target data store.
package load

import (
	"context"
	"have-a-nice-pickem-etl/internal/transform"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type New struct {
	transform.Transformation
}

func callBulkLoadProcedure(queryString string) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Error occurred instantiating Postgres Database connection: %s", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "%s", queryString)
	if err != nil {
		log.Printf("Error occurred with the SQL command %s: %s", queryString, err)
	}
}

func (l New) PerformLoad() {
	csvLoadFolderPath := instantiateLoadDirectory()

	loadBettingOdds(l.GameTransformations.AllBettingOdds, csvLoadFolderPath)
	loadBoxscores(l.GameTransformations.AllBoxscores, csvLoadFolderPath)
	loadGameDetails(l.GameTransformations.AllGameDetails, csvLoadFolderPath)
	loadGameStats(l.GameTransformations.AllGameStats, csvLoadFolderPath)

	loadTeamDetails(l.TeamTransformations.AllTeams, csvLoadFolderPath)
	loadTeamRecord(l.TeamTransformations.AllTeamRecords, csvLoadFolderPath)

	loadLocationDetails(l.LocationTransformations.AllLocations, csvLoadFolderPath)
}
