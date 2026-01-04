package bettingodds

import (
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c CbsBettingOdds) parseGameOddsTable() *goquery.Selection {
	var cbsGameCode string = common.ScrapeCbsGameCode(c.Game)
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

func (c CbsBettingOdds) parseOverUnder() float32 {
	var overUnderText string = c.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	var numericOverUnder string = strings.TrimSpace(overUnderText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "o", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "u", "")
	return utils.ConvertStringToFloat32(numericOverUnder)
}

func (c CbsBettingOdds) parseAwayMoneyline() int {
	var moneylineText string = c.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
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
	var moneylineText string = c.CBS.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
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
	var spreadText string = c.CBS.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	var formattedSpread string = strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}
	return utils.ConvertStringToFloat32(formattedSpread)
}

func (c CbsBettingOdds) parseHomeSpread() float32 {
	var spreadText string = c.CBS.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
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
