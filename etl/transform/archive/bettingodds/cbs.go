package bettingodds

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CbsOdds struct {
	pickemstructs.BettingOdds
}

func (c CbsOdds) ParseGameOddsTable(cbsOddsPage *goquery.Selection, cbsGameCode string) *goquery.Selection {
	var gameTables *goquery.Selection = cbsOddsPage.Find("table.OddsBlock-game")
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

func (c CbsOdds) ParseOverUnder(cbsGameOddsTable *goquery.Selection) float32 {
	var overUnderText string = cbsGameOddsTable.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--total").Find("div.BetButton-text").Text()
	//fmt.Printf("overUnderText: %s", overUnderText)
	var numericOverUnder string = strings.TrimSpace(overUnderText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "o", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "u", "")
	return utils.ConvertStringToFloat32(numericOverUnder)
}

func (c CbsOdds) ParseAwayMoneyline(cbsGameOddsTable *goquery.Selection) int16 {
	var moneylineText string = cbsGameOddsTable.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	var formattedMoneyline string = strings.TrimSpace(moneylineText)

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}
	return utils.ConvertStringToInt16(formattedMoneyline)
}

func (c CbsOdds) ParseHomeMoneyline(cbsGameOddsTable *goquery.Selection) int16 {
	var moneylineText string = cbsGameOddsTable.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--moneyline").Find("div.BetButton-text").Text()
	var formattedMoneyline string = strings.TrimSpace(moneylineText)

	if strings.Contains(formattedMoneyline, "+") {
		formattedMoneyline = strings.ReplaceAll(formattedMoneyline, "+", "")
	}
	return utils.ConvertStringToInt16(formattedMoneyline)
}

func (c CbsOdds) ParseAwaySpread(cbsGameOddsTable *goquery.Selection) float32 {
	var spreadText string = cbsGameOddsTable.Find("tbody").Find("tr").Eq(0).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	var formattedSpread string = strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}
	return utils.ConvertStringToFloat32(formattedSpread)
}

func (c CbsOdds) ParseHomeSpread(cbsGameOddsTable *goquery.Selection) float32 {
	var spreadText string = cbsGameOddsTable.Find("tbody").Find("tr").Eq(1).Find("td.OddsBlock-betOdds--spread").Find("div.BetButton-text").Text()
	var formattedSpread string = strings.TrimSpace(spreadText)

	if strings.Contains(formattedSpread, "+") {
		formattedSpread = strings.ReplaceAll(formattedSpread, "+", "")
	}
	return utils.ConvertStringToFloat32(formattedSpread)
}

/*
func ParseAwayWinPercentage(cbsGameOddsTable *goquery.Selection) string {
	return "50%"
}

func ParseHomeWinPercentage(cbsGameOddsTable *goquery.Selection) string {
	return "50%"
}
*/
