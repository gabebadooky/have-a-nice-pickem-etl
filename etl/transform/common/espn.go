package common

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/sharedtypes"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"
)

type HomeAwayTeam struct {
	HomeAway string
}

// Generates "TeamID" field from ESPN Team Summary API
func ParseTeamSummaryTeamID(teamSummaryDetails sharedtypes.ESPNTeamSummaryResponse) string {
	var teamID string = teamSummaryDetails.Team.ID
	return teamID
}

// Generates "TeamID" field for the Home or Away team from ESPN Game Summary API
func (t HomeAwayTeam) ParseGameSummaryTeamID(consolidatedGame extract.ConsolidatedGame) string {
	var competitorHomeAway string = consolidatedGame.ESPN.Header.Competitions[0].Competitors[0].HomeAway
	if strings.EqualFold(t.HomeAway, competitorHomeAway) {
		var teamID string = consolidatedGame.ESPN.Header.Competitions[0].Competitors[0].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
		return formattedTeamID
	} else {
		var teamID string = consolidatedGame.ESPN.Header.Competitions[0].Competitors[1].Team.DisplayName
		var formattedTeamID string = utils.FormatStringID(teamID)
		return formattedTeamID
	}

}

// Generates "GameID" field from AwayTeamID and HomeTeamID from ESPN Game Summary API
func ParseGameID(consolidatedGame extract.ConsolidatedGame) string {
	var awayTeamID string = HomeAwayTeam{HomeAway: "away"}.ParseGameSummaryTeamID(consolidatedGame)
	var homeTeamID string = HomeAwayTeam{HomeAway: "home"}.ParseGameSummaryTeamID(consolidatedGame)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = utils.FormatStringID(gameID)
	return formattedGameID
}
