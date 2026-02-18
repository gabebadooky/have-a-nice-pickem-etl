// Package common provides shared parsing and extraction utilities used across
// transformation packages. These functions extract identifiers and codes from
// game and team data structures.
package common

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/extract/team"
	"log"
	"strings"
)

// ParseEspnGameCode returns the ESPN game code from the game's ESPN header.
func ParseEspnGameCode(gameExtract game.Game) string {
	var gameCode string = gameExtract.ESPN.Header.ESPNGameCode
	return gameCode
}

// ScrapeCbsGameCode extracts the CBS game abbreviation from the game's CBS odds block.
func ScrapeCbsGameCode(gameExtract game.Game) string {
	gameCode, exists := gameExtract.CBS.Find("table.OddsBlock-game").Attr("data-game-abbrev")
	if !exists {
		log.Panicf("Could not locate CBS Game Code for %v", &gameExtract.CBS)
	}
	return gameCode
}

// ScrapeFoxGameCode extracts the Fox game path from the game's Fox odds page nav link.
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

// ParseAwayTeamID returns the away team ID from the game ID (segment before "-at-").
func ParseAwayTeamID(gameExtract game.Game) string {
	before, _, _ := strings.Cut(gameExtract.GameID, "-at-")
	return before
}

// ParseHomeTeamID returns the home team ID from the game ID (between "-at-" and "-week-").
func ParseHomeTeamID(gameExtract game.Game) string {
	fmt.Printf("gameExtract.GameID: %s", gameExtract.GameID)
	startIndex := strings.Index(gameExtract.GameID, "-at-") + 4
	endIndex := strings.LastIndex(gameExtract.GameID, "-week-")
	homeTeamID := gameExtract.GameID[startIndex:endIndex]
	return homeTeamID
}

// ParseEspnTeamCode returns the ESPN team code from the team's ESPN summary.
func ParseEspnTeamCode(teamExtract team.Team) string {
	var espnTeamCode string = teamExtract.ESPN.Team.Code
	return espnTeamCode
}
