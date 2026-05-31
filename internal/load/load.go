// Package load provides the loading layer for the ETL pipeline.
// This package handles writing transformed data to the target data store.
package load

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/transform"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type New struct {
	transform.Transformation
}

// InstantiateDatabaseConnection opens Postgres via GORM. The pickem schema is not set on
// the DSN; each model's TableName() returns a schema-qualified name (e.g. pickem.games).
func InstantiateDatabaseConnection() *gorm.DB {
	_ = godotenv.Load()

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("SSLMODE")
	channelBinding := os.Getenv("CHANNEL_BINDING")
	timezone := os.Getenv("TIMEZONE")
	schemaName := os.Getenv("SCHEMA")
	if schemaName == "" {
		schemaName = "pickem"
	}

	passwordPart := fmt.Sprintf("password=%s", password)
	if password == "" {
		passwordPart = "password=''"
	}

	dsn := fmt.Sprintf("host=%s user=%s %s dbname=%s port=%s sslmode=%s channel_binding=%s TimeZone=%s search_path=%s,public",
		host, user, passwordPart, dbname, port, sslmode, channelBinding, timezone,
		schemaName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error occurred connecting to Postgres Database %s:\n%s", os.Getenv("DATABASE_NAME"), err)
	}

	// Ensure the session is actually using the expected schema (and log what is visible)
	if err := db.Exec(fmt.Sprintf(`SET search_path TO "%s", public`, schemaName)).Error; err != nil {
		log.Fatalf("Error setting search_path to %s: %v", schemaName, err)
	}

	var searchPath string
	if err := db.Raw(`SHOW search_path`).Scan(&searchPath).Error; err == nil {
		log.Printf("Postgres search_path=%s", searchPath)
	}

	type tableRow struct {
		TableSchema string `gorm:"column:table_schema"`
		TableName   string `gorm:"column:table_name"`
	}
	var tables []tableRow
	if err := db.Raw(`
		SELECT table_schema, table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		  AND table_schema = current_schema()
		ORDER BY table_schema, table_name
	`).Scan(&tables).Error; err == nil {
		log.Printf("Tables visible in current_schema(): %d", len(tables))
		for _, t := range tables {
			log.Printf("table=%s.%s", t.TableSchema, t.TableName)
		}
	}

	return db
}

func (l New) PerformLoad() {
	db := InstantiateDatabaseConnection()

	if len(l.ConferenceTransformations.AllConferences) > 0 {
		loadConferenceDetails(l.ConferenceTransformations.AllConferences, db)
	}

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
		loadGameStats(l.GameTransformations.AllGameStats, db)
	}

	if len(l.ForecastTransformations.AllForecasts) > 0 {
		loadForecastDetails(l.ForecastTransformations.AllForecasts, db)
	}

}

func TestConnection() {
	godotenv.Load()
	db := InstantiateDatabaseConnection()

	type schemaRow struct {
		SchemaName string `gorm:"column:schema_name"`
	}

	var schemas []schemaRow
	if err := db.Raw(`
		SELECT schema_name
		FROM information_schema.schemata
		ORDER BY schema_name
	`).Scan(&schemas).Error; err != nil {
		log.Fatalf("Error querying schemas: %v", err)
	}

	log.Printf("Schemas returned: %d", len(schemas))
	for _, s := range schemas {
		log.Printf("schema_name=%s", s.SchemaName)
	}
}
