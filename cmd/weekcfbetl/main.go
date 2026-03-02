package weekcfbetl

import (
	"have-a-nice-pickem-etl/internal/extract"
	"have-a-nice-pickem-etl/internal/load"
	"have-a-nice-pickem-etl/internal/transform"
)

// Runs the ETL pipeline for a given weeks games, teams, and locations
func main(week uint) {
	weekExtract := extract.CfbExtract{Week: week}.PerformExtract()
	weekTransformation := transform.New{Extract: weekExtract}.PerformTransformation()
	load.New{Transformation: weekTransformation}.PerformLoad()
}
