package gamedetails

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
	"have-a-nice-pickem-etl/etl/utils"
)

// Generates "GameID" field from AwayTeamID and HomeTeamID from ESPN Game Summary API
func ParseGameID(espnGameDetails types.ESPNGameDetailsResponse) string {
	var awayTeamID string = ParseTeamID("away", espnGameDetails)
	var homeTeamID string = ParseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = utils.FormatStringID(gameID)
	return formattedGameID
}

// Parses "League" field from ESPN Game Summary API
func ParseLeague(espnGameDetails types.ESPNGameDetailsResponse) string {
	var league string = espnGameDetails.Header.League.Abbreviation
	if league == "NCAAF" {
		return "CFB"
	} else {
		return "NFL"
	}
}

// Parses "Week" field from ESPN Game Summary API
func ParseWeek(espnGameDetails types.ESPNGameDetailsResponse) int8 {
	var week int8 = espnGameDetails.Header.Week
	return week

}

// Parses "Year" field from ESPN Game Summary API
func ParseYear(espnGameDetails types.ESPNGameDetailsResponse) uint16 {
	var year uint16 = espnGameDetails.Header.Season.Year
	return year

}

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameCode string = espnGameDetails.Header.ESPNGameCode
	return gameCode

}

// Generates "TeamID" field for the Home or Away team from ESPN Game Summary API
func ParseTeamID(homeAway string, espnGameDetails types.ESPNGameDetailsResponse) string {
	var competitorHomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	if homeAway == competitorHomeAway {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[0].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
		return formattedTeamID
	} else {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[1].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
		return formattedTeamID
	}

}

// Parses "Date" field from ESPN Game Summary API
func ParseGameZuluTimestamp(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameDate string = espnGameDetails.Header.Competitions[0].Date
	return gameDate
}

// Parses "Broadcast" field from ESPN Game Summary API
func ParseBroadcast(espnGameDetails types.ESPNGameDetailsResponse) string {
	var broadcast string = espnGameDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

// Parses "Status" field from ESPN Game Summary API
func ParseGameStatus(espnGameDetails types.ESPNGameDetailsResponse) bool {
	var gameStatus bool = espnGameDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}
