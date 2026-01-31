package common

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/extract/team"
	"log"
	"strings"
)

// Parses "ID" field from ESPN Game Summary API
func ParseEspnGameCode(gameExtract game.Game) string {
	var gameCode string = gameExtract.ESPN.Header.ESPNGameCode
	return gameCode
}

func ScrapeCbsGameCode(gameExtract game.Game) string {
	gameCode, exists := gameExtract.CBS.Find("table.OddsBlock-game").Attr("data-game-abbrev")
	if !exists {
		log.Panicf("Could not locate CBS Game Code for %v", &gameExtract.CBS)
	}
	return gameCode
}

func ScrapeFoxGameCode(gameExtract game.Game) string {
	matchupCell := gameExtract.FOX.OddsPage.Find("div.nav-horizontal").Find("a").First()
	gameCode, exists := matchupCell.Attr("href")
	if !exists {
		log.Panicf("Could not locate Fox Game Code for %v", gameExtract.FOX)
	}
	formattedGameCode := gameCode[1:]
	_, stringAfterSportPrefix, _ := strings.Cut(formattedGameCode, "/")
	stringBeforeQueryParams, _, _ := strings.Cut(stringAfterSportPrefix, "?")
	//gameCodeWithoutBowlPrefix := utils.StripBowlGamePrefixFromFoxGameCode(stringBeforeQueryParams)
	//gameCodeWithDateSuffix := utils.StripDateAndBoxScoreIDFromFoxGameCode(gameCodeWithoutBowlPrefix)
	return stringBeforeQueryParams
}

func ParseAwayTeamID(gameExtract game.Game) string {
	before, _, _ := strings.Cut(gameExtract.GameID, "-at-")
	return before
}

func ParseHomeTeamID(gameExtract game.Game) string {
	fmt.Printf("gameExtract.GameID: %s", gameExtract.GameID)
	startIndex := strings.Index(gameExtract.GameID, "-at-") + 4
	endIndex := strings.LastIndex(gameExtract.GameID, "-week-")
	homeTeamID := gameExtract.GameID[startIndex:endIndex]
	return homeTeamID
}

func ParseEspnTeamCode(teamExtract team.Team) string {
	var espnTeamCode string = teamExtract.ESPN.Team.Code
	return espnTeamCode
}
