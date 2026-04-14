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

	if len(l.LocationTransformations.AllLocations) > 0 {
		loadLocationDetails(l.LocationTransformations.AllLocations, db)
	}

	if len(l.TeamTransformations.AllTeams) > 0 {
		loadTeamDetails(l.TeamTransformations.AllTeams, db)
	}

	if len(l.TeamTransformations.AllTeamRecords) > 0 {
		loadTeamRecord(l.TeamTransformations.AllTeamRecords, db)
	}

	if len(l.GameTransformations.AllGameDetails) > 0 {
		loadGameDetails(l.GameTransformations.AllGameDetails, db)
	}

	if len(l.GameTransformations.AllBettingOdds) > 0 {
		loadBettingOdds(l.GameTransformations.AllBettingOdds, db)
	}

	if len(l.GameTransformations.AllBoxscores) > 0 {
		loadBoxscores(l.GameTransformations.AllBoxscores, db)
	}

	if len(l.GameTransformations.AllGameStats) > 0 {
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("GameStats Length: %d", len(l.GameTransformations.AllGameStats))
		fmt.Println("")
		fmt.Printf("GameStats: %v", l.GameTransformations.AllGameStats)
		fmt.Println("")
		fmt.Println("")
		loadGameStats(l.GameTransformations.AllGameStats, db)
	}

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
