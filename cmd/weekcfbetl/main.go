package main

import (
	"flag"
	"have-a-nice-pickem-etl/internal/extract"
	"have-a-nice-pickem-etl/internal/load"
	"have-a-nice-pickem-etl/internal/transform"

	"github.com/joho/godotenv"
)

// Runs the ETL pipeline for a given weeks games, teams, and locations
func main() {
	godotenv.Load()

	week := flag.Uint("week", 1, "The week to extract")
	flag.Parse()
	weekExtract := extract.CfbExtract{Week: *week}.PerformExtract()
	weekTransformation := transform.New{Extract: weekExtract}.PerformTransformation()
	load.New{Transformation: weekTransformation}.PerformLoad()
}
