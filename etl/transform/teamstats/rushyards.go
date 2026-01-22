package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamRushYards() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushYardsTD := teamTotalsTableRow.Find("td").Eq(3)
	rushYards := strings.TrimSpace(rushYardsTD.Text())
	var formattedRushYards float32 = utils.ConvertStringToFloat32(rushYards)

	return Stat{
		Metric: "rush_yards",
		Value:  formattedRushYards,
	}
}

func (t New) scrapeOpponentRushYards() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushYardsTD := opponentTotalTableRow.Find("td").Eq(3)
	rushYards := strings.TrimSpace(rushYardsTD.Text())
	var formattedRushYards float32 = utils.ConvertStringToFloat32(rushYards)

	return Stat{
		Metric: "opp_rush_yards",
		Value:  formattedRushYards,
	}
}
