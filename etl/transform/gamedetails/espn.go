package gamedetails

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/location"
	"have-a-nice-pickem-etl/etl/utils"
)

// Parses "League" field from ESPN Game Summary API
func ParseLeague(consolidatedGameProps extract.ConsolidatedGame) string {
	var league string = consolidatedGameProps.ESPN.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func ParseWeek(consolidatedGameProps extract.ConsolidatedGame) int8 {
	var week int8 = consolidatedGameProps.ESPN.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func ParseYear(consolidatedGameProps extract.ConsolidatedGame) uint16 {
	var year uint16 = consolidatedGameProps.ESPN.Header.Season.Year
	return year

}

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(consolidatedGameProps extract.ConsolidatedGame) string {
	var gameCode string = consolidatedGameProps.ESPN.Header.ESPNGameCode
	return gameCode

}

// Parses "Date" field from ESPN Game Summary API
func ParseGameZuluTimestamp(consolidatedGameProps extract.ConsolidatedGame) string {
	var gameDate string = consolidatedGameProps.ESPN.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func ParseBroadcast(consolidatedGameProps extract.ConsolidatedGame) string {
	var broadcast string = consolidatedGameProps.ESPN.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Location" field from ESPN Game Summary API
func ParseLocationID(consolidatedGameProps extract.ConsolidatedGame) string {
	var formattedStadium string = utils.FormatStringID(location.ParseStadium(consolidatedGameProps))
	var formattedCity string = utils.FormatStringID(location.ParseCity(consolidatedGameProps))
	var formattedState string = utils.FormatStringID(location.ParseState(consolidatedGameProps))
	var concatenatedLocation string = fmt.Sprintf("%s-%s-%s", formattedStadium, formattedCity, formattedState)
	return concatenatedLocation
}

// Parses "Status" field from ESPN Game Summary API
func ParseGameStatus(consolidatedGameProps extract.ConsolidatedGame) bool {
	var gameStatus bool = consolidatedGameProps.ESPN.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}
