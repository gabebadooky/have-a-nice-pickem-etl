package gamedetails

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/location"
	"have-a-nice-pickem-etl/etl/utils"
)

// Parses "League" field from ESPN Game Summary API
func ParseLeague(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) string {
	var league string = consolidatedGameProps.EspnDetails.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func ParseWeek(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) int8 {
	var week int8 = consolidatedGameProps.EspnDetails.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func ParseYear(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) uint16 {
	var year uint16 = consolidatedGameProps.EspnDetails.Header.Season.Year
	return year

}

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) string {
	var gameCode string = consolidatedGameProps.EspnDetails.Header.ESPNGameCode
	return gameCode

}

// Parses "Date" field from ESPN Game Summary API
func ParseGameZuluTimestamp(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) string {
	var gameDate string = consolidatedGameProps.EspnDetails.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func ParseBroadcast(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) string {
	var broadcast string = consolidatedGameProps.EspnDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Location" field from ESPN Game Summary API
func ParseLocationID(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) string {
	var formattedStadium string = utils.FormatStringID(location.ParseStadium(consolidatedGameProps))
	var formattedCity string = utils.FormatStringID(location.ParseCity(consolidatedGameProps))
	var formattedState string = utils.FormatStringID(location.ParseState(consolidatedGameProps))
	var concatenatedLocation string = fmt.Sprintf("%s-%s-%s", formattedStadium, formattedCity, formattedState)
	return concatenatedLocation
}

// Parses "Status" field from ESPN Game Summary API
func ParseGameStatus(consolidatedGameProps pickemstructs.ConsolidatedGameProperties) bool {
	var gameStatus bool = consolidatedGameProps.EspnDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}
