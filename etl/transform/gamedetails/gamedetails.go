package gamedetails

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type Instantiator interface {
	instantiate() GameDetails
}

type NewGameDetails struct {
	GameExtract game.Game
}

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

func InstantiateGameDetails(i Instantiator) GameDetails {
	return i.instantiate()
}

// Parses "League" field from ESPN Game Summary API
func (g NewGameDetails) parseLeague() string {
	var league string = g.GameExtract.ESPN.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func (g NewGameDetails) parseWeek() int8 {
	var week int8 = g.GameExtract.ESPN.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func (g NewGameDetails) parseYear() uint {
	var year uint = g.GameExtract.ESPN.Header.Season.Year
	return year

}

// Parses "Date" field from ESPN Game Summary API
func (g NewGameDetails) parseGameZuluTimestamp() string {
	var gameDate string = g.GameExtract.ESPN.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func (g NewGameDetails) parseBroadcast() string {
	var broadcast string = g.GameExtract.ESPN.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Location" field from ESPN Game Summary API
/*func (g NewGameDetails) parseLocation() string {
	var formattedStadium string = utils.FormatStringID(location.ParseStadium(gameExtract))
	var formattedCity string = utils.FormatStringID(location.ParseCity(gameExtract))
	var formattedState string = utils.FormatStringID(location.ParseState(gameExtract))
	var concatenatedLocation string = fmt.Sprintf("%s-%s-%s", formattedStadium, formattedCity, formattedState)
	return concatenatedLocation
}*/

// Parses "Status" field from ESPN Game Summary API
func (g NewGameDetails) parseGameStatus() bool {
	var gameStatus bool = g.GameExtract.ESPN.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

func (g NewGameDetails) instantiate() GameDetails {
	var gameExtract game.Game = g.GameExtract

	return GameDetails{
		GameID:        gameExtract.GameID,
		League:        g.parseLeague(),
		Week:          g.parseWeek(),
		Year:          g.parseYear(),
		EspnCode:      common.ParseEspnGameCode(gameExtract),
		CbsCode:       common.ScrapeCbsGameCode(gameExtract),
		FoxCode:       common.ScrapeFoxGameCode(gameExtract),
		VegasCode:     "",
		AwayTeamID:    common.ParseAwayTeamID(gameExtract),
		HomeTeamID:    common.ParseHomeTeamID(gameExtract),
		ZuluTimestamp: g.parseGameZuluTimestamp(),
		Broadcast:     g.parseBroadcast(),
		Location:      "",
		Finished:      g.parseGameStatus(),
	}
}
