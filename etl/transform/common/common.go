package common

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/extract/team"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
	"strings"
)

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(gameExtract game.Game) string {
	var gameCode string = gameExtract.ESPN.Header.ESPNGameCode
	return gameCode
}

func ScrapeCbsGameCode(gameExtract game.Game) string {
	gameCode, exists := gameExtract.CBS.Attr("data-game-abbrev")
	if !exists {
		log.Panicf("Could not locate CBS Game Code for %v", gameExtract.CBS)
	}

	return gameCode
}

func ScrapeFoxGameCode(gameExtract game.Game) string {
	broadcastCell := gameExtract.FOX.Find("td.broadcast")
	gameCode, exists := broadcastCell.Attr("")
	if !exists {
		log.Panicf("Could not locate CBS Game Code for %v", gameExtract.FOX)
	}

	return gameCode
}

func ParseAwayTeamID(gameExtract game.Game) string {
	before, _, _ := strings.Cut(gameExtract.GameID, "-vs-")
	return before
}

func ParseHomeTeamID(gameExtract game.Game) string {
	startIndex := strings.Index(gameExtract.GameID, "-vs-") + 4
	endIndex := strings.LastIndex(gameExtract.GameID, "-week-")
	homeTeamID := gameExtract.GameID[startIndex:endIndex]
	return homeTeamID
}

func ParseEspnTeamCode(teamExtract team.Team) string {
	var espnTeamCode string = teamExtract.ESPN.Team.Code
	return espnTeamCode
}

func GetCbsTeamCode(teamID string) string {
	cbsCode, cbsMappingExists := utils.TeamIDtoCbsTeamCode[teamID]
	if cbsMappingExists {
		return cbsCode
	}
	return teamID

}

func GetFoxTeamCode(teamID string) string {
	foxCode, foxMappingExists := utils.TeamIDtoFoxTeamCode[teamID]
	if foxMappingExists {
		return foxCode
	}
	return teamID
}
