// Package game provides game extraction functionality that consolidates game data
// from multiple sources (ESPN API, CBS web scraping, Fox web scraping) into
// a unified Game structure for both college football (CFB) and NFL.
package game

import (
	"fmt"
	cbsgame "have-a-nice-pickem-etl/internal/extract/game/cbs"
	espngame "have-a-nice-pickem-etl/internal/extract/game/espn"
	foxgame "have-a-nice-pickem-etl/internal/extract/game/fox"
	espnsched "have-a-nice-pickem-etl/internal/extract/schedule/espn"
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CfbGame struct {
	EspnEvent       espnsched.EventProperty
	CbsSchedulePage *goquery.Selection
	FoxSchedulePage *goquery.Selection
}

type NflGame struct {
	EspnEvent       espnsched.EventProperty
	CbsSchedulePage *goquery.Selection
	FoxSchedulePage *goquery.Selection
}

type Game struct {
	GameID string
	ESPN   espngame.GameSummaryEndpoint
	CBS    *goquery.Selection
	FOX    foxgame.FoxGamePages
}

// instantiateGameID builds a GameID from the event name and week (or post-season).
func instantiateGameID(sched espnsched.EventProperty) string {
	var eventNameAndWeek string
	if sched.Season.Type == 3 {
		eventNameAndWeek = fmt.Sprintf("%s week post season", sched.Name)
	} else {
		eventNameAndWeek = fmt.Sprintf("%s week %d", sched.Name, sched.Week.Number)
	}
	gameID := utils.FormatStringID(eventNameAndWeek)
	return gameID
}

// extractGame fetches and consolidates college football game data from ESPN, CBS, and Fox.
func (c CfbGame) extractGame() (Game, error) {
	gameID := instantiateGameID(c.EspnEvent)
	fmt.Printf("\nEvent: %s", gameID)

	if strings.Contains(gameID, "-tbd-") {
		return Game{}, fmt.Errorf("GameID includes \"tbd\"")
	}

	return Game{
		GameID: gameID,
		ESPN:   espngame.EspnCfbGame{GameCode: c.EspnEvent.ID}.GetGameSummary(),
		CBS:    cbsgame.CbsGame{CbsOddsPage: c.CbsSchedulePage, GameId: gameID}.GetGamePage(),
		FOX:    foxgame.FoxGame{FoxSchedulePage: c.FoxSchedulePage, GameID: gameID}.GetGamePage(),
	}, nil
}

// extractGame fetches and consolidates NFL game data from ESPN, CBS, and Fox.
func (n NflGame) extractGame() (Game, error) {
	gameID := instantiateGameID(n.EspnEvent)
	fmt.Printf("\nEvent: %s", gameID)

	if strings.Contains(gameID, "-tbd-") {
		return Game{}, fmt.Errorf("GameID includes \"tbd\"")
	}
	return Game{
		GameID: gameID,
		ESPN:   espngame.EspnNflGame{GameCode: n.EspnEvent.ID}.GetGameSummary(),
		CBS:    cbsgame.CbsGame{CbsOddsPage: n.CbsSchedulePage, GameId: gameID}.GetGamePage(),
		FOX:    foxgame.FoxGame{FoxSchedulePage: n.FoxSchedulePage, GameID: gameID}.GetGamePage(),
	}, nil
}
