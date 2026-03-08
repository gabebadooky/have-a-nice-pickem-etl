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
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
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
	csvLoadFolderPath := instantiateLoadDirectory()

	loadBettingOdds(l.GameTransformations.AllBettingOdds, db)
	loadBoxscores(l.GameTransformations.AllBoxscores, csvLoadFolderPath)
	loadGameDetails(l.GameTransformations.AllGameDetails, csvLoadFolderPath)
	loadGameStats(l.GameTransformations.AllGameStats, csvLoadFolderPath)

	loadTeamDetails(l.TeamTransformations.AllTeams, csvLoadFolderPath)
	loadTeamRecord(l.TeamTransformations.AllTeamRecords, csvLoadFolderPath)

	loadLocationDetails(l.LocationTransformations.AllLocations, csvLoadFolderPath)
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
