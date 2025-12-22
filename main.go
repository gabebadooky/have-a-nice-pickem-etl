package main

import "have-a-nice-pickem-etl/etl/extract"

func main() {
	var week uint8 = 7
	games := extract.CfbGamesExtract{Week: week}.ExtractGames()
	teams := extract.CfbTeamsExtract{Week: week}.ExtractTeams()
}
