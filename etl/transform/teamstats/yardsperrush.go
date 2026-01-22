package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamYardsPerRush() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	yardsPerRushTD := teamTotalsTableRow.Find("td").Eq(4)
	yardsPerRush := strings.TrimSpace(yardsPerRushTD.Text())
	var formattedYardsPerRush float32 = utils.ConvertStringToFloat32(yardsPerRush)

	return Stat{
		Metric: "yards_per_rush",
		Value:  formattedYardsPerRush,
	}
}

func (t New) scrapeOpponentYardsPerRush() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	yardsPerRushTD := opponentTotalTableRow.Find("td").Eq(4)
	yardsPerRush := strings.TrimSpace(yardsPerRushTD.Text())
	var formattedYardsPerRush float32 = utils.ConvertStringToFloat32(yardsPerRush)

	return Stat{
		Metric: "opp_yards_per_rush",
		Value:  formattedYardsPerRush,
	}
}
