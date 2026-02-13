// Package bettingodds provides CBS Sports betting odds parsing functionality.
// It extracts betting odds data from CBS Sports HTML pages including spread,
// moneyline, and over/under for both away and home teams.
package bettingodds

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
	"have-a-nice-pickem-etl/internal/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsAwayBettingOdds struct {
	game.Game
}

type CbsHomeBettingOdds struct {
	game.Game
}

func parseGameOddsTable(c game.Game) *goquery.Selection {
	var cbsGameCode string = common.ScrapeCbsGameCode(c)
	var gameTables *goquery.Selection = c.CBS.Find("table.OddsBlock-game")
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

func formatCBSOverUnder(overUnderText string) float32 {
	formattedOverUnder := strings.TrimSpace(overUnderText)
	formattedOverUnder = strings.ReplaceAll(formattedOverUnder, "o", "")
	formattedOverUnder = strings.ReplaceAll(formattedOverUnder, "u", "")
	var overUnderFloat32 float32 = utils.ConvertStringToFloat32(formattedOverUnder)
	return overUnderFloat32
}

func formatCBSMoneyline(moneylineText string) int {
	formattedMoneyline := strings.TrimSpace(moneylineText)

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}

	numericMoneyline, err := strconv.Atoi(formattedMoneyline)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

func formatCBSSpread(spreadText string) float32 {
	formattedSpread := strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}

	var spreadFloat32 float32 = utils.ConvertStringToFloat32(formattedSpread)
	return spreadFloat32
}

func (c CbsAwayBettingOdds) parseOverUnder() float32 {
	// var overUnderText string = c.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	var overUnderText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	overUnderFloat32 := formatCBSOverUnder(overUnderText)
	return overUnderFloat32
}

func (c CbsHomeBettingOdds) parseOverUnder() float32 {
	// var overUnderText string = c.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	var overUnderText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	overUnderFloat32 := formatCBSOverUnder(overUnderText)
	return overUnderFloat32
}

func (c CbsAwayBettingOdds) parseMoneyline() int {
	var moneylineText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	numericMoneyline := formatCBSMoneyline(moneylineText)
	return numericMoneyline
}

func (c CbsHomeBettingOdds) parseMoneyline() int {
	var moneylineText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	numericMoneyline := formatCBSMoneyline(moneylineText)
	return numericMoneyline
}

func (c CbsAwayBettingOdds) parseSpread() float32 {
	var spreadText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	spreadFloat32 := formatCBSSpread(spreadText)
	return spreadFloat32
}

func (c CbsHomeBettingOdds) parseSpread() float32 {
	var spreadText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	spreadFloat32 := formatCBSSpread(spreadText)
	return spreadFloat32
}

func (c CbsAwayBettingOdds) parseWinProbability() float32 {
	return .5
}

func (c CbsHomeBettingOdds) parseWinProbability() float32 {
	return .5
}
