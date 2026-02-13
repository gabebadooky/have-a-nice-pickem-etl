// Package teamstats provides yards per rush statistics scraping functionality.
// It extracts yards per rush data from CBS Sports team stats pages for both
// team and opponent statistics.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamYardsPerRush struct {
	page *goquery.Selection
}

type oppYardsPerRush struct {
	page *goquery.Selection
}

func (yr teamYardsPerRush) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(yr.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	yardsPerRushTD := teamTotalsTableRow.Find("td").Eq(4)
	yardsPerRush := strings.TrimSpace(yardsPerRushTD.Text())
	var formattedYardsPerRush float32 = utils.ConvertStringToFloat32(yardsPerRush)

	return Stat{
		Metric: "yards_per_rush",
		Value:  formattedYardsPerRush,
	}
}

func (yr oppYardsPerRush) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(yr.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	yardsPerRushTD := opponentTotalTableRow.Find("td").Eq(4)
	yardsPerRush := strings.TrimSpace(yardsPerRushTD.Text())
	var formattedYardsPerRush float32 = utils.ConvertStringToFloat32(yardsPerRush)

	return Stat{
		Metric: "opp_yards_per_rush",
		Value:  formattedYardsPerRush,
	}
}
