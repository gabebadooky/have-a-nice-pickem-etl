package transform

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type GameDetails struct {
	GameID        string
	League        string
	Week          int8
	Year          uint
	EspnCode      string
	CbsCode       string
	FoxCode       string
	VegasCode     string
	AwayTeamID    string
	HomeTeamID    string
	ZuluTimestamp string
	Broadcast     string
	Location      string
	Finished      bool
}

// Parses "League" field from ESPN Game Summary API
func parseLeague(gameExtract game.Game) string {
	var league string = gameExtract.ESPN.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func parseWeek(gameExtract game.Game) int8 {
	var week int8 = gameExtract.ESPN.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func parseYear(gameExtract game.Game) uint {
	var year uint = gameExtract.ESPN.Header.Season.Year
	return year

}

// Parses "Date" field from ESPN Game Summary API
func parseGameZuluTimestamp(gameExtract game.Game) string {
	var gameDate string = gameExtract.ESPN.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func parseBroadcast(gameExtract game.Game) string {
	var broadcast string = gameExtract.ESPN.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Location" field from ESPN Game Summary API
/*func parseLocation(gameExtract game.Game) string {
	var formattedStadium string = utils.FormatStringID(location.ParseStadium(gameExtract))
	var formattedCity string = utils.FormatStringID(location.ParseCity(gameExtract))
	var formattedState string = utils.FormatStringID(location.ParseState(gameExtract))
	var concatenatedLocation string = fmt.Sprintf("%s-%s-%s", formattedStadium, formattedCity, formattedState)
	return concatenatedLocation
}*/

// Parses "Status" field from ESPN Game Summary API
func parseGameStatus(gameExtract game.Game) bool {
	var gameStatus bool = gameExtract.ESPN.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

func (gameTransformations GameTransformations) InstantiateGameDetails() GameDetails {
	var gameExtract game.Game = gameTransformations.GameExtract

	return GameDetails{
		GameID:        gameExtract.GameID,
		League:        parseLeague(gameExtract),
		Week:          parseWeek(gameExtract),
		Year:          parseYear(gameExtract),
		EspnCode:      common.ParseEspnGameCode(gameExtract),
		CbsCode:       common.ScrapeCbsGameCode(gameExtract),
		FoxCode:       common.ScrapeFoxGameCode(gameExtract),
		VegasCode:     "",
		AwayTeamID:    common.ParseAwayTeamID(gameExtract),
		HomeTeamID:    common.ParseHomeTeamID(gameExtract),
		ZuluTimestamp: parseGameZuluTimestamp(gameExtract),
		Broadcast:     parseBroadcast(gameExtract),
		Location:      "",
		Finished:      parseGameStatus(gameExtract),
	}
}
