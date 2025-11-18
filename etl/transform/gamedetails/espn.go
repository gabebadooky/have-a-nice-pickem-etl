package gamedetails

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
)

// Parses "League" field from ESPN Game Summary API
func ParseLeague(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var league string = espnGameDetails.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func ParseWeek(espnGameDetails pickemstructs.ESPNGameDetailsResponse) int8 {
	var week int8 = espnGameDetails.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func ParseYear(espnGameDetails pickemstructs.ESPNGameDetailsResponse) uint16 {
	var year uint16 = espnGameDetails.Header.Season.Year
	return year

}

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var gameCode string = espnGameDetails.Header.ESPNGameCode
	return gameCode

}

// Parses "Date" field from ESPN Game Summary API
func ParseGameZuluTimestamp(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var gameDate string = espnGameDetails.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func ParseBroadcast(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var broadcast string = espnGameDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Status" field from ESPN Game Summary API
func ParseGameStatus(espnGameDetails pickemstructs.ESPNGameDetailsResponse) bool {
	var gameStatus bool = espnGameDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}
