package gamedetails

import (
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type GameRow struct {
	GameID        string
	League        string
	Week          int8
	Year          uint16
	ESPNCode      string
	CBSCode       string
	FoxCode       string
	VegasCode     string
	AwayTeamID    string
	HomeTeamID    string
	ZuluTimestamp string
	Broadcast     string
	Location      string
	Finished      bool
}

// Transforms and consolidates Game Details properties from various sources
func (g GameRow) Instantiate(consolidatedGame extract.ConsolidatedGame) GameDetails {
	var gameID string = common.ParseGameID(consolidatedGame)

	return GameRow{
		GameID:        gameID,
		League:        ParseLeague(consolidatedGame),
		Week:          ParseWeek(consolidatedGame),
		Year:          ParseYear(consolidatedGame),
		ESPNCode:      ParseEspnGameCode(consolidatedGame),
		CBSCode:       common.ExtractCbsGameCode(consolidatedGame.CBS, gameID),
		FoxCode:       common.ExtractFoxGameCode(consolidatedGame.FOX, gameID),
		VegasCode:     "",
		AwayTeamID:    common.HomeAwayTeam{HomeAway: "away"}.ParseGameSummaryTeamID(consolidatedGame),
		HomeTeamID:    common.HomeAwayTeam{HomeAway: "home"}.ParseGameSummaryTeamID(consolidatedGame),
		ZuluTimestamp: ParseGameZuluTimestamp(consolidatedGame),
		Broadcast:     ParseBroadcast(consolidatedGame),
		Location:      ParseLocationID(consolidatedGame),
		Finished:      ParseGameStatus(consolidatedGame),
	}
}
