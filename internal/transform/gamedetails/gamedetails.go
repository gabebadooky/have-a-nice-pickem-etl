// Package gamedetails provides game details transformation functionality that extracts
// and structures comprehensive game information including league, week, year, game codes
// from multiple sources (ESPN, CBS, Fox), team IDs, timestamps, broadcast info, and game status.
package gamedetails

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	espngame "have-a-nice-pickem-etl/internal/extract/game/espn"
	"have-a-nice-pickem-etl/internal/extract/location"
	"have-a-nice-pickem-etl/internal/transform/common"
)

type New struct {
	game.Game
	Locations []location.Location
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
	LocationID    string
	Finished      bool
}

func (g New) setCbsGameCode() string {
	cbsGameCode, err := common.ScrapeCbsGameCode(g.Game)
	if err != nil {
		return "CbsGameCode"
	}
	return cbsGameCode
}

func (g New) setFoxGameCode() string {
	foxGameCode, err := common.ScrapeFoxGameCode(g.Game)
	if err != nil {
		foxGameCode = "FoxGameCode"
	}
	return foxGameCode
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
	var broadcastSlice []espngame.BroadcastsProperty = g.ESPN.Header.Competitions[0].Broadcasts
	if len(broadcastSlice) == 0 {
		return ""
	}
	var broadcast string = broadcastSlice[0].Media.ShortName
	return broadcast
}

// Returns the Maidenhead property, as LocationID, for a given game
func (g New) parseLocationID() string {
	var gameStadium string = g.ESPN.GameInfo.Venue.FullName
	var gameCity string = g.ESPN.GameInfo.Venue.Address.City
	var gameState string = g.ESPN.GameInfo.Venue.Address.State

	for i := range g.Locations {
		opencageLocationResults := g.Locations[i].Opencage.Results[0]

		stadium := opencageLocationResults.Components.Stadium
		city := opencageLocationResults.Components.City
		state := opencageLocationResults.Components.State

		if (gameStadium == stadium) && (gameCity == city) && (gameState == state) {
			return opencageLocationResults.Annotations.Maidenhead
		}
	}

	return ""
}

// parseGameStatus returns whether the game is completed from the ESPN competition status.
func (g New) parseGameStatus() bool {
	var gameStatus bool = g.ESPN.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

// InstantiateGameDetails builds a GameDetails value from the extracted game data.
func (g New) Instantiate() GameDetails {
	return GameDetails{
		GameID:        g.GameID,
		League:        g.parseLeague(),
		Week:          g.parseWeek(),
		Year:          g.parseYear(),
		EspnCode:      common.ParseEspnGameCode(g.Game),
		CbsCode:       g.setCbsGameCode(),
		FoxCode:       g.setFoxGameCode(),
		VegasCode:     "",
		AwayTeamID:    common.ParseAwayTeamID(g.Game),
		HomeTeamID:    common.ParseHomeTeamID(g.Game),
		ZuluTimestamp: g.parseGameZuluTimestamp(),
		Broadcast:     g.parseBroadcast(),
		LocationID:    g.parseLocationID(),
		Finished:      g.parseGameStatus(),
	}
}
