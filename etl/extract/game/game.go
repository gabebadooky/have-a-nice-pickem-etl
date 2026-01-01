package game

import (
	cbsgame "have-a-nice-pickem-etl/etl/extract/game/cbs"
	espngame "have-a-nice-pickem-etl/etl/extract/game/espn"
	foxgame "have-a-nice-pickem-etl/etl/extract/game/fox"
	espnsched "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/utils"

	"github.com/PuerkitoBio/goquery"
)

type AllGameInfo interface {
	gameInfo() Game
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
	ESPN espngame.GameSummaryEndpoint
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

func ConsolidateGameInfo(g AllGameInfo) Game {
	return g.gameInfo()
}

func (c AllCfbGameInfo) gameInfo() Game {
	var gameID string = utils.FormatStringID(c.EspnEvent.Name)
	var espnGame espngame.GameSummaryEndpoint = espngame.GetGameSummary(espngame.EspnCfbGame{GameCode: gameID})
	var cbsGame *goquery.Selection = cbsgame.CbsGame{CbsOddsPage: c.CbsSchedulePage, GameId: gameID}.ExtractCbsGameHTML()
	var foxGame *goquery.Selection = foxgame.FoxGame{FoxSchedulePage: c.FoxSchedulePage, GameID: gameID}.ExtractFoxGameHTML()

	return Game{
		ESPN: espnGame,
		CBS:  cbsGame,
		FOX:  foxGame,
	}
}

func (n AllNflGameInfo) gameInfo() Game {
	var gameID string = utils.FormatStringID(n.EspnEvent.Name)
	var EspnGame espngame.GameSummaryEndpoint = espngame.GetGameSummary(espngame.EspnNflGame{GameCode: gameID})
	var CbsGame *goquery.Selection = cbsgame.CbsGame{CbsOddsPage: n.CbsSchedulePage, GameId: gameID}.ExtractCbsGameHTML()
	var FoxGame *goquery.Selection = foxgame.FoxGame{FoxSchedulePage: n.FoxSchedulePage, GameID: gameID}.ExtractFoxGameHTML()

	return Game{
		ESPN: EspnGame,
		CBS:  CbsGame,
		FOX:  FoxGame,
	}
}
