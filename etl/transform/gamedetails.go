package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"
)

// Transforms and consolidates Game Details properties from various sources
func CreateGameDetailsRecord(consolidatedGameSources pickemstructs.ConsolidatedGameProperties) pickemstructs.GameDetails {
	var gameID string = common.ParseGameID(consolidatedGameSources)

	var newRecord pickemstructs.GameDetails = pickemstructs.GameDetails{
		GameID:        gameID,
		League:        gamedetails.ParseLeague(consolidatedGameSources),
		Week:          gamedetails.ParseWeek(consolidatedGameSources),
		Year:          gamedetails.ParseYear(consolidatedGameSources),
		ESPNCode:      gamedetails.ParseEspnGameCode(consolidatedGameSources),
		CBSCode:       common.ExtractCbsGameCode(consolidatedGameSources.CbsPage, gameID),
		FoxCode:       gamedetails.ExtractFoxGameCode(consolidatedGameSources.FoxPage, gameID),
		VegasCode:     "",
		AwayTeamID:    common.ParseGameSummaryTeamID("away", consolidatedGameSources),
		HomeTeamID:    common.ParseGameSummaryTeamID("home", consolidatedGameSources),
		ZuluTimestamp: gamedetails.ParseGameZuluTimestamp(consolidatedGameSources),
		Broadcast:     gamedetails.ParseBroadcast(consolidatedGameSources),
		Location:      gamedetails.ParseLocationID(consolidatedGameSources),
		Finished:      gamedetails.ParseGameStatus(consolidatedGameSources),
	}

	return newRecord
}
