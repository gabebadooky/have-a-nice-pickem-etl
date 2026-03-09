// Package load provides the loading layer for the ETL pipeline.
// This package handles writing transformed data to the target data store.
package load

import (
	"context"
	"fmt"
	"have-a-nice-pickem-etl/internal/transform"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type New struct {
	transform.Transformation
}

func instantiateDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("SSLMODE"),
		os.Getenv("TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error occurred connecting to Postgres Database %s:\n%s", os.Getenv("DATABASE_NAME"), err)
	}

	return db
}

func callBulkLoadProcedure(queryString string) {
	godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Error occurred instantiating Postgres Database connection: %s", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), queryString)
	if err != nil {
		log.Printf("Error occurred with the SQL command %s: %s", queryString, err)
	}
}

func (l New) PerformLoad() {
	db := instantiateDatabaseConnection()

	loadBettingOdds(l.GameTransformations.AllBettingOdds, db)
	loadBoxscores(l.GameTransformations.AllBoxscores, db)
	loadGameDetails(l.GameTransformations.AllGameDetails, db)
	loadGameStats(l.GameTransformations.AllGameStats, db)

	loadTeamDetails(l.TeamTransformations.AllTeams, db)
	loadTeamRecord(l.TeamTransformations.AllTeamRecords, db)

	loadLocationDetails(l.LocationTransformations.AllLocations, db)
}

func TestConnection() {
	godotenv.Load()
	ctx := context.Background()
	databaseURL := os.Getenv("DATABASE_URL")
	println(databaseURL)
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Printf("Error occurred instantiating Postgres Database connection: %s", err)
	}
	defer conn.Close(context.Background())
}
