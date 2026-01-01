package main

import (
	"have-a-nice-pickem-etl/etl/extract"
)

func main() {
	var week uint = 7
	var league string = "CFB"

	if league == "CFB" {
		games := extract.ExtractGames(extract.CfbGamesExtract{Week: week})
		teams := extract.ExtractTeams(extract.CfbTeamsExtract{Week: week})
		locations := extract.ExtractLocations(extract.CfbLocationsExtract{Week: week})
	}
}
