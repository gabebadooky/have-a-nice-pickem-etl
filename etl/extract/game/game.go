package game

import (
	"fmt"
	cbsgame "have-a-nice-pickem-etl/etl/extract/game/cbs"
	espngame "have-a-nice-pickem-etl/etl/extract/game/espn"
	foxgame "have-a-nice-pickem-etl/etl/extract/game/fox"
	espnsched "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AllCfbGameInfo struct {
	EspnEvent       espnsched.EventProperty
	CbsSchedulePage *goquery.Selection
	FoxSchedulePage *goquery.Selection
}

type AllNflGameInfo struct {
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
type instantiator interface {
	gameInfo() (Game, error)
}

func ConsolidateGameInfo(g instantiator) (Game, error) {
	return g.gameInfo()
}

// Instantiate GameID based on whether a game is a regular season or post-season matchup
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

func (c AllCfbGameInfo) gameInfo() (Game, error) {
	gameID := instantiateGameID(c.EspnEvent)
	fmt.Printf("\nEvent: %s", gameID)

	if strings.Contains(gameID, "-tbd-") {
		return Game{}, fmt.Errorf("GameID includes \"tbd\"")
	}

	var espnGame espngame.GameSummaryEndpoint = espngame.GetGameSummary(espngame.EspnCfbGame{GameCode: c.EspnEvent.ID})
	var cbsGame *goquery.Selection = cbsgame.GetGamePage(cbsgame.CbsGame{CbsOddsPage: c.CbsSchedulePage, GameId: gameID})
	var foxGame foxgame.FoxGamePages = foxgame.GetGamePages(foxgame.FoxGame{FoxSchedulePage: c.FoxSchedulePage, GameID: gameID})

	return Game{
		GameID: gameID,
		ESPN:   espnGame,
		CBS:    cbsGame,
		FOX:    foxGame,
	}, nil
}

func (n AllNflGameInfo) gameInfo() (Game, error) {
	gameID := instantiateGameID(n.EspnEvent)
	fmt.Printf("\nEvent: %s", gameID)

	if strings.Contains(gameID, "-tbd-") {
		return Game{}, fmt.Errorf("GameID includes \"tbd\"")
	}

	var espnGame espngame.GameSummaryEndpoint = espngame.GetGameSummary(espngame.EspnNflGame{GameCode: n.EspnEvent.ID})
	var cbsGame *goquery.Selection = cbsgame.GetGamePage(cbsgame.CbsGame{CbsOddsPage: n.CbsSchedulePage, GameId: gameID})
	var foxGame foxgame.FoxGamePages = foxgame.GetGamePages(foxgame.FoxGame{FoxSchedulePage: n.FoxSchedulePage, GameID: gameID})

	return Game{
		GameID: gameID,
		ESPN:   espnGame,
		CBS:    cbsGame,
		FOX:    foxGame,
	}, nil
}
