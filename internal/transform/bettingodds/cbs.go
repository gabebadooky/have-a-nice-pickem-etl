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

// parseGameOddsTable finds the CBS odds table row matching the game's CBS game code.
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

// formatCBSOverUnder parses CBS over/under text (e.g. "45.5 o/u") into a float32.
func formatCBSOverUnder(overUnderText string) float32 {
	formattedOverUnder := strings.TrimSpace(overUnderText)
	if formattedOverUnder == "" {
		return float32(0)
	}

	formattedOverUnder = strings.ReplaceAll(formattedOverUnder, "o", "")
	formattedOverUnder = strings.ReplaceAll(formattedOverUnder, "u", "")
	var overUnderFloat32 float32 = utils.ConvertStringToFloat32(formattedOverUnder)
	return overUnderFloat32
}

// formatCBSMoneyline parses CBS moneyline text (e.g. "+120") into an int.
func formatCBSMoneyline(moneylineText string) int {
	formattedMoneyline := strings.TrimSpace(moneylineText)
	if formattedMoneyline == "" {
		return 0
	}

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}

	numericMoneyline, err := strconv.Atoi(formattedMoneyline)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

// formatCBSSpread parses CBS spread text into a float32.
func formatCBSSpread(spreadText string) float32 {
	formattedSpread := strings.TrimSpace(spreadText)
	if formattedSpread == "" {
		return float32(0)
	}

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}

	var spreadFloat32 float32 = utils.ConvertStringToFloat32(formattedSpread)
	return spreadFloat32
}

// parseOverUnder returns the over/under from the game's CBS odds table (away row).
func (c CbsAwayBettingOdds) parseOverUnder() float32 {
	var overUnderText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	overUnderFloat32 := formatCBSOverUnder(overUnderText)
	return overUnderFloat32
}

// parseOverUnder returns the over/under from the game's CBS odds table (home row).
func (c CbsHomeBettingOdds) parseOverUnder() float32 {
	var overUnderText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	overUnderFloat32 := formatCBSOverUnder(overUnderText)
	return overUnderFloat32
}

// parseMoneyline returns the away team moneyline from the game's CBS odds table.
func (c CbsAwayBettingOdds) parseMoneyline() int {
	var moneylineText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	numericMoneyline := formatCBSMoneyline(moneylineText)
	return numericMoneyline
}

// parseMoneyline returns the home team moneyline from the game's CBS odds table.
func (c CbsHomeBettingOdds) parseMoneyline() int {
	var moneylineText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	numericMoneyline := formatCBSMoneyline(moneylineText)
	return numericMoneyline
}

// parseSpread returns the away team spread from the game's CBS odds table.
func (c CbsAwayBettingOdds) parseSpread() float32 {
	var spreadText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	spreadFloat32 := formatCBSSpread(spreadText)
	return spreadFloat32
}

// parseSpread returns the home team spread from the game's CBS odds table.
func (c CbsHomeBettingOdds) parseSpread() float32 {
	var spreadText string = parseGameOddsTable(c.Game).Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	spreadFloat32 := formatCBSSpread(spreadText)
	return spreadFloat32
}

// parseWinProbability returns a placeholder 0.5 (CBS does not expose win probability).
func (c CbsAwayBettingOdds) parseWinProbability() float32 {
	return .5
}

// parseWinProbability returns a placeholder 0.5 (CBS does not expose win probability).
func (c CbsHomeBettingOdds) parseWinProbability() float32 {
	return .5
}
