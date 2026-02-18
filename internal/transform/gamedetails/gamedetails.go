// Package gamedetails provides game details transformation functionality that extracts
// and structures comprehensive game information including league, week, year, game codes
// from multiple sources (ESPN, CBS, Fox), team IDs, timestamps, broadcast info, and game status.
package gamedetails

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
)

type New struct {
	game.Game
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

// parseLeague returns the league abbreviation (CFB or NFL) from the ESPN game header.
func (g New) parseLeague() string {
	var league string = g.ESPN.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// parseWeek returns the week number from the ESPN game header.
func (g New) parseWeek() int8 {
	var week int8 = g.ESPN.Header.Week
	return week

}

// parseYear returns the season year from the ESPN game header.
func (g New) parseYear() uint {
	var year uint = g.ESPN.Header.Season.Year
	return year

}

// parseGameZuluTimestamp returns the game date from the ESPN competition data.
func (g New) parseGameZuluTimestamp() string {
	var gameDate string = g.ESPN.Header.Competitions[0].Date
	return gameDate
}

// parseBroadcast returns the broadcast network from the ESPN competition data.
func (g New) parseBroadcast() string {
	var broadcast string = g.ESPN.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Location" field from ESPN Game Summary API
/*func (g New) parseLocation() string {
	var formattedStadium string = utils.FormatStringID(location.ParseStadium(g))
	var formattedCity string = utils.FormatStringID(location.ParseCity(g))
	var formattedState string = utils.FormatStringID(location.ParseState(g))
	var concatenatedLocation string = fmt.Sprintf("%s-%s-%s", formattedStadium, formattedCity, formattedState)
	return concatenatedLocation
}*/

// parseGameStatus returns whether the game is completed from the ESPN competition status.
func (g New) parseGameStatus() bool {
	var gameStatus bool = g.ESPN.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

// InstantiateGameDetails builds a GameDetails value from the extracted game data.
func (g New) InstantiateGameDetails() GameDetails {
	return GameDetails{
		GameID:        g.GameID,
		League:        g.parseLeague(),
		Week:          g.parseWeek(),
		Year:          g.parseYear(),
		EspnCode:      common.ParseEspnGameCode(g.Game),
		CbsCode:       common.ScrapeCbsGameCode(g.Game),
		FoxCode:       common.ScrapeFoxGameCode(g.Game),
		VegasCode:     "",
		AwayTeamID:    common.ParseAwayTeamID(g.Game),
		HomeTeamID:    common.ParseHomeTeamID(g.Game),
		ZuluTimestamp: g.parseGameZuluTimestamp(),
		Broadcast:     g.parseBroadcast(),
		Location:      "",
		Finished:      g.parseGameStatus(),
	}
}
