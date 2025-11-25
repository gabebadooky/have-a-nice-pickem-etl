package common

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/utils"
)

// Generates "TeamID" field from ESPN Team Summary API
func ParseTeamSummaryTeamID(teamSummaryDetails pickemstructs.TeamSummaryResponse) string {
	var teamID string = teamSummaryDetails.Team.ID
	return teamID
}

// Generates "TeamID" field for the Home or Away team from ESPN Game Summary API
func ParseGameSummaryTeamID(homeAway string, espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
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

// Generates "GameID" field from AwayTeamID and HomeTeamID from ESPN Game Summary API
func ParseGameID(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var awayTeamID string = ParseGameSummaryTeamID("away", espnGameDetails)
	var homeTeamID string = ParseGameSummaryTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = utils.FormatStringID(gameID)
	return formattedGameID
}
