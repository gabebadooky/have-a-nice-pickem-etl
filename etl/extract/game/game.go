package game

import (
	"have-a-nice-pickem-etl/etl/extract/game/cbs"
	"have-a-nice-pickem-etl/etl/extract/game/espn"
	"have-a-nice-pickem-etl/etl/extract/game/fox"
	espnsched "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/utils"

	"github.com/PuerkitoBio/goquery"
)

type AllGameInfo interface {
	ConsolidateGameInfo() Game
}

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
	ESPN espn.GameSummaryEndpoint
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

func (c AllCfbGameInfo) ConsolidateGameInfo() Game {
	var gameID string = utils.FormatStringID(c.EspnEvent.Name)
	var EspnGame espn.GameSummaryEndpoint = espn.EspnCfbGame{GameCode: c.EspnEvent.ID}.GetGameSummary()
	var CbsGame *goquery.Selection = cbs.CbsGame{CbsOddsPage: c.CbsSchedulePage, GameId: gameID}.ExtractCbsGameHTML()
	var FoxGame *goquery.Selection = fox.FoxGame{FoxSchedulePage: c.FoxSchedulePage, GameID: gameID}.ExtractFoxGameHTML()

	return Game{
		ESPN: EspnGame,
		CBS:  CbsGame,
		FOX:  FoxGame,
	}
}

func (n AllNflGameInfo) ConsolidateGameInfo() Game {
	var gameID string = utils.FormatStringID(n.EspnEvent.Name)
	var EspnGame espn.GameSummaryEndpoint = espn.EspnNflGame{GameCode: n.EspnEvent.ID}.GetGameSummary()
	var CbsGame *goquery.Selection = cbs.CbsGame{CbsOddsPage: n.CbsSchedulePage, GameId: gameID}.ExtractCbsGameHTML()
	var FoxGame *goquery.Selection = fox.FoxGame{FoxSchedulePage: n.FoxSchedulePage, GameID: gameID}.ExtractFoxGameHTML()

	return Game{
		ESPN: EspnGame,
		CBS:  CbsGame,
		FOX:  FoxGame,
	}
}
