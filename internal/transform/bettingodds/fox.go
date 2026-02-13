// Package bettingodds provides Fox Sports betting odds parsing functionality.
// It extracts betting odds data from Fox Sports HTML pages including spread,
// moneyline, over/under, and win probability for both away and home teams.
package bettingodds

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type FoxAwayBettingOdds struct {
	game.Game
}

type FoxHomeBettingOdds struct {
	game.Game
}

func scrapeProbabilityChart(odds game.Game) *goquery.Selection {
	gameProbabilityChart := odds.FOX.OddsPage.Find("div.win-probability-chart")
	return gameProbabilityChart
}

func scrapeGameOddsTable(odds game.Game) *goquery.Selection {
	gameOddsTable := odds.FOX.OddsPage.Find("div.odds-sp-content")
	return gameOddsTable
}

func scrapeAwayTeamOddsTableRow(odds game.Game) *goquery.Selection {
	gameOddsTable := scrapeGameOddsTable(odds)
	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(0)
	return teamOddsTableRow
}

func scrapeHomeTeamOddsTableRow(odds game.Game) *goquery.Selection {
	gameOddsTable := scrapeGameOddsTable(odds)
	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(1)
	return teamOddsTableRow
}

func formatFoxOverUnder(oddsTableRow *goquery.Selection) float32 {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(2).Text()
	numericOverUnder := strings.TrimSpace(tableCellText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "O", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "U", "")
	overUnderFloat32 := utils.ConvertStringToFloat32(numericOverUnder)
	return overUnderFloat32
}

func formatFoxMoneyline(oddsTableRow *goquery.Selection) int {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(1).Text()
	moneylineString := strings.ReplaceAll(tableCellText, "+", "")

	numericMoneyline, err := strconv.Atoi(moneylineString)
	if err != nil {
		numericMoneyline = 0
	}
	return numericMoneyline
}

func formatFoxSpread(oddsTableRow *goquery.Selection) float32 {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(0).Text()
	spreadString := strings.ReplaceAll(tableCellText, "+", "")
	spreadFloat := utils.ConvertStringToFloat32(spreadString)
	return spreadFloat
}

func formatFoxWinProbability(probabilityText string) float32 {
	probabilityText = strings.Split(probabilityText, " ")[1]
	probabilityText = strings.ReplaceAll(probabilityText, "%", "")
	percentageFloat32 := utils.ConvertStringToFloat32(probabilityText)
	return percentageFloat32
}

func (odds FoxAwayBettingOdds) parseOverUnder() float32 {
	oddsTableRow := scrapeAwayTeamOddsTableRow(odds.Game)
	overUnderFloat32 := formatFoxOverUnder(oddsTableRow)
	return overUnderFloat32
}

func (odds FoxHomeBettingOdds) parseOverUnder() float32 {
	oddsTableRow := scrapeHomeTeamOddsTableRow(odds.Game)
	overUnderFloat32 := formatFoxOverUnder(oddsTableRow)
	return overUnderFloat32
}

func (odds FoxAwayBettingOdds) parseMoneyline() int {
	oddsTableRow := scrapeAwayTeamOddsTableRow(odds.Game)
	numericMoneyline := formatFoxMoneyline(oddsTableRow)
	return numericMoneyline
}

func (odds FoxHomeBettingOdds) parseMoneyline() int {
	oddsTableRow := scrapeHomeTeamOddsTableRow(odds.Game)
	numericMoneyline := formatFoxMoneyline(oddsTableRow)
	return numericMoneyline
}

func (odds FoxAwayBettingOdds) parseSpread() float32 {
	oddsTableRow := scrapeAwayTeamOddsTableRow(odds.Game)
	spreadFloat32 := formatFoxSpread(oddsTableRow)
	return spreadFloat32
}

func (odds FoxHomeBettingOdds) parseSpread() float32 {
	oddsTableRow := scrapeHomeTeamOddsTableRow(odds.Game)
	spreadFloat32 := formatFoxSpread(oddsTableRow)
	return spreadFloat32
}

func (odds FoxAwayBettingOdds) parseWinProbability() float32 {
	probabilityChart := scrapeProbabilityChart(odds.Game)
	probabilityText := probabilityChart.Find("div.ff-g").Eq(0).Text()
	percentageFloat32 := formatFoxWinProbability(probabilityText)
	return percentageFloat32
}

func (odds FoxHomeBettingOdds) parseWinProbability() float32 {
	probabilityChart := scrapeProbabilityChart(odds.Game)
	probabilityText := probabilityChart.Find("div.ff-g").Eq(1).Text()
	percentageFloat32 := formatFoxWinProbability(probabilityText)
	return percentageFloat32
}
