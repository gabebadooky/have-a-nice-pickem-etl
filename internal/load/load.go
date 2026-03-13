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

func InstantiateDatabaseConnection() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("SSLMODE")
	timezone := os.Getenv("TIMEZONE")

	passwordPart := fmt.Sprintf("password=%s", password)
	if password == "" {
		passwordPart = "password=''"
	}

	dsn := fmt.Sprintf("host=%s user=%s %s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, passwordPart, dbname, port, sslmode, timezone,
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
	db := InstantiateDatabaseConnection()

	loadLocationDetails(l.LocationTransformations.AllLocations, db)

	loadTeamDetails(l.TeamTransformations.AllTeams, db)
	loadTeamRecord(l.TeamTransformations.AllTeamRecords, db)

	loadGameDetails(l.GameTransformations.AllGameDetails, db)

	loadBettingOdds(l.GameTransformations.AllBettingOdds, db)
	loadBoxscores(l.GameTransformations.AllBoxscores, db)
	loadGameStats(l.GameTransformations.AllGameStats, db)
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
