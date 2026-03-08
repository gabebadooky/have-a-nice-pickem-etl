// Package bettingodds provides Fox Sports betting odds parsing functionality.
// It extracts betting odds data from Fox Sports HTML pages including spread,
// moneyline, over/under, and win probability for both away and home teams.
package bettingodds

import (
	"fmt"
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

// scrapeProbabilityChart returns the Fox win probability chart selection from the odds page.
func scrapeProbabilityChart(odds game.Game) (*goquery.Selection, error) {
	if odds.FOX.OddsPage == nil {
		return nil, fmt.Errorf("Fox page does not exist for game %s", odds.GameID)
	}

	gameProbabilityChart := odds.FOX.OddsPage.Find("div.win-probability-chart")
	return gameProbabilityChart, nil
}

// scrapeGameOddsTable returns the Fox odds content selection from the odds page.
func scrapeGameOddsTable(odds game.Game) (*goquery.Selection, error) {
	if odds.FOX.OddsPage == nil {
		return nil, fmt.Errorf("Fox page does not exist for game %s", odds.GameID)
	}

	gameOddsTable := odds.FOX.OddsPage.Find("div.odds-sp-content")
	return gameOddsTable, nil
}

// scrapeAwayTeamOddsTableRow returns the away team row from the Fox odds table.
func scrapeAwayTeamOddsTableRow(odds game.Game) (*goquery.Selection, error) {
	gameOddsTable, err := scrapeGameOddsTable(odds)
	if err != nil {
		return nil, err
	}

	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(0)
	return teamOddsTableRow, nil
}

// scrapeHomeTeamOddsTableRow returns the home team row from the Fox odds table.
func scrapeHomeTeamOddsTableRow(odds game.Game) (*goquery.Selection, error) {
	gameOddsTable, err := scrapeGameOddsTable(odds)
	if err != nil {
		return nil, err
	}

	teamOddsTableRow := gameOddsTable.Find("div.sp-rows").Eq(1)
	return teamOddsTableRow, nil
}

// formatFoxOverUnder extracts and parses the over/under value from a Fox odds row.
func formatFoxOverUnder(oddsTableRow *goquery.Selection) float32 {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(2).Text()
	if tableCellText == "" {
		return float32(0)
	}

	numericOverUnder := strings.TrimSpace(tableCellText)
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "O", "")
	numericOverUnder = strings.ReplaceAll(numericOverUnder, "U", "")
	overUnderFloat32 := utils.ConvertStringToFloat32(numericOverUnder)
	return overUnderFloat32
}

// formatFoxMoneyline extracts and parses the moneyline from a Fox odds row.
func formatFoxMoneyline(oddsTableRow *goquery.Selection) int {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(1).Text()
	if tableCellText == "" {
		return 0
	}

	moneylineString := strings.ReplaceAll(tableCellText, "+", "")
	numericMoneyline, err := strconv.Atoi(moneylineString)
	if err != nil {
		numericMoneyline = 0
	}
	return numericMoneyline
}

// formatFoxSpread extracts and parses the spread from a Fox odds row.
func formatFoxSpread(oddsTableRow *goquery.Selection) float32 {
	tableCellText := oddsTableRow.Find("div.sp-row-data").Eq(0).Text()
	if tableCellText == "" {
		return float32(0)
	}

	spreadString := strings.ReplaceAll(tableCellText, "+", "")
	spreadFloat := utils.ConvertStringToFloat32(spreadString)
	return spreadFloat
}

// formatFoxWinProbability parses the win probability percentage from Fox probability text.
func formatFoxWinProbability(probabilityText string) float32 {
	probabilityStringElements := strings.Split(probabilityText, " ")
	if len(probabilityStringElements) < 2 {
		return float32(0)
	}

	probabilityText = strings.ReplaceAll(probabilityStringElements[1], "%", "")
	if probabilityText == "" {
		return float32(0)
	}

	percentageFloat32 := utils.ConvertStringToFloat32(probabilityText)
	return percentageFloat32
}

// parseOverUnder returns the over/under from the game's Fox odds page (away row).
func (odds FoxAwayBettingOdds) parseOverUnder() float32 {
	oddsTableRow, err := scrapeAwayTeamOddsTableRow(odds.Game)
	if err != nil {
		return float32(0)
	}

	overUnderFloat32 := formatFoxOverUnder(oddsTableRow)
	return overUnderFloat32
}

// parseOverUnder returns the over/under from the game's Fox odds page (home row).
func (odds FoxHomeBettingOdds) parseOverUnder() float32 {
	oddsTableRow, err := scrapeHomeTeamOddsTableRow(odds.Game)
	if err != nil {
		return float32(0)
	}

	overUnderFloat32 := formatFoxOverUnder(oddsTableRow)
	return overUnderFloat32
}

// parseMoneyline returns the away team moneyline from the game's Fox odds page.
func (odds FoxAwayBettingOdds) parseMoneyline() int {
	oddsTableRow, err := scrapeAwayTeamOddsTableRow(odds.Game)
	if err != nil {
		return 0
	}

	numericMoneyline := formatFoxMoneyline(oddsTableRow)
	return numericMoneyline
}

// parseMoneyline returns the home team moneyline from the game's Fox odds page.
func (odds FoxHomeBettingOdds) parseMoneyline() int {
	oddsTableRow, err := scrapeHomeTeamOddsTableRow(odds.Game)
	if err != nil {
		return 0
	}

	numericMoneyline := formatFoxMoneyline(oddsTableRow)
	return numericMoneyline
}

// parseSpread returns the away team spread from the game's Fox odds page.
func (odds FoxAwayBettingOdds) parseSpread() float32 {
	oddsTableRow, err := scrapeAwayTeamOddsTableRow(odds.Game)
	if err != nil {
		return float32(0)
	}

	spreadFloat32 := formatFoxSpread(oddsTableRow)
	return spreadFloat32
}

// parseSpread returns the home team spread from the game's Fox odds page.
func (odds FoxHomeBettingOdds) parseSpread() float32 {
	oddsTableRow, err := scrapeHomeTeamOddsTableRow(odds.Game)
	if err != nil {
		return float32(0)
	}

	spreadFloat32 := formatFoxSpread(oddsTableRow)
	return spreadFloat32
}

// parseWinProbability returns the away team win probability from the game's Fox odds page.
func (odds FoxAwayBettingOdds) parseWinProbability() float32 {
	probabilityChart, err := scrapeProbabilityChart(odds.Game)
	if err != nil {
		return float32(0)
	}

	probabilityText := probabilityChart.Find("div.ff-g").Eq(0).Text()
	percentageFloat32 := formatFoxWinProbability(probabilityText)
	return percentageFloat32
}

// parseWinProbability returns the home team win probability from the game's Fox odds page.
func (odds FoxHomeBettingOdds) parseWinProbability() float32 {
	probabilityChart, err := scrapeProbabilityChart(odds.Game)
	if err != nil {
		return float32(0)
	}

	probabilityText := probabilityChart.Find("div.ff-g").Eq(1).Text()
	percentageFloat32 := formatFoxWinProbability(probabilityText)
	return percentageFloat32
}
