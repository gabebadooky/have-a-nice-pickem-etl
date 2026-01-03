package transform

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Instantiator interface {
	Instantiate() BettingOdds
}

type EspnBettingOdds struct {
	Game game.Game
}

type CbsBettingOdds struct {
	Game game.Game
}

type FoxBettingOdds struct {
	Game game.Game
}

type BettingOdds struct {
	GameID            string
	Source            string
	OverUnder         float32
	AwayMoneyline     int
	HomeMoneyline     int
	AwaySpread        float32
	HomeSpread        float32
	AwayWinPercentage string
	HomeWinPercentage string
}

func InstantiateBettingOdds(i Instantiator) BettingOdds {
	return i.Instantiate()
}

func (e EspnBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.Game.ESPN.Pickcenter[0].OverUnder
	return overUnder
}

func (e EspnBettingOdds) parseAwayMoneyline() int {
	var awayMoneyline int = e.Game.ESPN.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

func (e EspnBettingOdds) parseHomeMoneyline() int {
	var homeMoneyline int = e.Game.ESPN.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

func (e EspnBettingOdds) parseAwaySpread() float32 {
	var awaySpread float32 = e.Game.ESPN.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

func (e EspnBettingOdds) parseHomeSpread() float32 {
	var homeSpread float32 = e.Game.ESPN.Pickcenter[0].Spread
	return homeSpread
}

func (e EspnBettingOdds) parseAwayWinPercentage() string {
	var awayWinPercentage string = e.Game.ESPN.Predictor.AwayTeam.GameProjection
	return awayWinPercentage
}

func (e EspnBettingOdds) parseHomeWinPercentage() string {
	var homeWinPercentage string = e.Game.ESPN.Predictor.HomeTeam.GameProjection
	return homeWinPercentage
}

func (c CbsBettingOdds) parseGameOddsTable() *goquery.Selection {
	var gameExtract game.Game = c.Game
	var cbsGameCode string = common.ScrapeCbsGameCode(gameExtract)
	var gameTables *goquery.Selection = gameExtract.CBS.Find("table.OddsBlock-game")
	var gameOddsTable *goquery.Selection

	gameTables.EachWithBreak(func(i int, gameTable *goquery.Selection) bool {
		var dataAbbr string = gameTable.AttrOr("data-game-abbrev", "cbsGameCode")

		if cbsGameCode == dataAbbr {
			gameOddsTable = gameTable
			return false
		} else {
			return true
		}
	})

	return gameOddsTable
}

func (c CbsBettingOdds) parseOverUnder() float32 {
	var overUnderText string = c.Game.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	var numericOverUnder string = strings.TrimSpace(overUnderText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "o", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "u", "")
	return utils.ConvertStringToFloat32(numericOverUnder)
}

func (c CbsBettingOdds) parseAwayMoneyline() int {
	var moneylineText string = c.Game.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	var formattedMoneyline string = strings.TrimSpace(moneylineText)

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}

	numericMoneyline, err := strconv.Atoi(formattedMoneyline)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

func (c CbsBettingOdds) parseHomeMoneyline() int {
	var moneylineText string = c.Game.CBS.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	var formattedMoneyline string = strings.TrimSpace(moneylineText)

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}

	numericMoneyline, err := strconv.Atoi(formattedMoneyline)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

func (c CbsBettingOdds) parseAwaySpread() float32 {
	var spreadText string = c.Game.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	var formattedSpread string = strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}
	return utils.ConvertStringToFloat32(formattedSpread)
}

func (c CbsBettingOdds) parseHomeSpread() float32 {
	var spreadText string = c.Game.CBS.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	var formattedSpread string = strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}
	return utils.ConvertStringToFloat32(formattedSpread)
}

func (c CbsBettingOdds) parseAwayWinPercentage() string {
	return "50%"
}

func (c CbsBettingOdds) parseHomeWinPercentage() string {
	return "50%"
}

// Instantiates ESPN Betting Odds record from various sources
func (b EspnBettingOdds) CreateESPNBettingOddsRecord() BettingOdds {
	return BettingOdds{
		GameID:        b.Game.GameID,
		Source:        "ESPN",
		OverUnder:     b.parseOverUnder(),
		AwayMoneyline: b.parseAwayMoneyline(),
		HomeMoneyline: b.parseHomeMoneyline(),
		AwaySpread:    b.parseAwaySpread(),
		HomeSpread:    b.parseHomeSpread(),
	}
}

// Instantiates CBS Betting Odds record from various sources
func (b CbsBettingOdds) Instantiate() BettingOdds {
	return BettingOdds{
		GameID:            b.Game.GameID,
		Source:            "CBS",
		OverUnder:         b.parseOverUnder(),
		AwayMoneyline:     b.parseAwayMoneyline(),
		HomeMoneyline:     b.parseHomeMoneyline(),
		AwaySpread:        b.parseAwaySpread(),
		HomeSpread:        b.parseHomeSpread(),
		AwayWinPercentage: b.parseAwayWinPercentage(),
		HomeWinPercentage: b.parseHomeWinPercentage(),
	}
}
