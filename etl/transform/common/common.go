package common

import (
	"have-a-nice-pickem-etl/etl/extract/game"
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
		log.Panicf("Could not locate CBS Game Code for %s", gameExtract.CBS)
	}

	return gameCode
}

func ScrapeFoxGameCode(gameExtract game.Game) string {
	broadcastCell := gameExtract.FOX.Find("td.broadcast")
	gameCode, exists := broadcastCell.Attr("")
	if !exists {
		log.Panicf("Could not locate CBS Game Code for %s", gameExtract.FOX)
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
