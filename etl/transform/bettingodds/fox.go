package bettingodds

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (odds FoxBettingOdds) scrapeProbabilityChart() *goquery.Selection {
	gameProbabilityChart := odds.FOX.OddsPage.Find("div.win-probability-chart")
	return gameProbabilityChart
}

func (odds FoxBettingOdds) scrapeGameOddsTable() *goquery.Selection {
	gameOddsTable := odds.FOX.OddsPage.Find("div.odds-sp-content")
	return gameOddsTable
}

func (odds FoxBettingOdds) scrapeAwayTeamOddsTableRow() *goquery.Selection {
	gameOddsTable := odds.scrapeGameOddsTable()
	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(0)
	return teamOddsTableRow
}

func (odds FoxBettingOdds) scrapeHomeTeamOddsTableRow() *goquery.Selection {
	gameOddsTable := odds.scrapeGameOddsTable()
	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(1)
	return teamOddsTableRow
}

func (odds FoxBettingOdds) parseOverUnder() float32 {
	oddsTableRow := odds.scrapeAwayTeamOddsTableRow()
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(2).Text()
	numericOverUnder := strings.TrimSpace(tableCellText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "O", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "U", "")
	overUnderFloat := utils.ConvertStringToFloat32(numericOverUnder)
	return overUnderFloat
}

func (odds FoxBettingOdds) parseAwayMoneyline() int {
	oddsTableRow := odds.scrapeAwayTeamOddsTableRow()
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(1).Text()
	moneylineString := strings.ReplaceAll(tableCellText, "+", "")

	numericMoneyline, err := strconv.Atoi(moneylineString)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

func (odds FoxBettingOdds) parseHomeMoneyline() int {
	oddsTableRow := odds.scrapeHomeTeamOddsTableRow()
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(1).Text()
	moneylineString := strings.ReplaceAll(tableCellText, "+", "")

	numericMoneyline, err := strconv.Atoi(moneylineString)
	if err != nil {
		numericMoneyline = 0
	}

	return numericMoneyline
}

func (odds FoxBettingOdds) parseAwaySpread() float32 {
	oddsTableRow := odds.scrapeAwayTeamOddsTableRow()
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(0).Text()
	spreadString := strings.ReplaceAll(tableCellText, "+", "")

	spreadFloat := utils.ConvertStringToFloat32(spreadString)
	return spreadFloat
}

func (odds FoxBettingOdds) parseHomeSpread() float32 {
	oddsTableRow := odds.scrapeHomeTeamOddsTableRow()
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(0).Text()
	spreadString := strings.ReplaceAll(tableCellText, "+", "")

	spreadFloat := utils.ConvertStringToFloat32(spreadString)
	return spreadFloat
}

func (odds FoxBettingOdds) parseAwayWinPercentage() string {
	probabilityChart := odds.scrapeProbabilityChart()
	probabilityText := probabilityChart.Find("div.ff-g").Eq(0).Text()
	percentage := strings.Split(probabilityText, " ")[1]
	return percentage
}

func (odds FoxBettingOdds) parseHomeWinPercentage() string {
	probabilityChart := odds.scrapeProbabilityChart()
	probabilityText := probabilityChart.Find("div.ff-g").Eq(1).Text()
	percentage := strings.Split(probabilityText, " ")[1]
	return percentage
}
