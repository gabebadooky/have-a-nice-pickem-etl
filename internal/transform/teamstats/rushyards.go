// Package teamstats provides rushing yards statistics scraping functionality.
// It extracts rushing yards data from CBS Sports team stats pages for both
// team and opponent statistics.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamRushYards struct {
	page *goquery.Selection
}

type oppRushYards struct {
	page *goquery.Selection
}

// scrape extracts team rush yards from the CBS rushing stats table.
func (ry teamRushYards) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(ry.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushYardsTD := teamTotalsTableRow.Find("td").Eq(3)
	rushYards := strings.TrimSpace(rushYardsTD.Text())
	var formattedRushYards float32 = utils.ConvertStringToFloat32(rushYards)

	return Stat{
		Metric: "rush_yards",
		Value:  formattedRushYards,
	}
}

// scrape extracts opponent rush yards from the CBS rushing stats table.
func (ry oppRushYards) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(ry.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushYardsTD := opponentTotalTableRow.Find("td").Eq(3)
	rushYards := strings.TrimSpace(rushYardsTD.Text())
	var formattedRushYards float32 = utils.ConvertStringToFloat32(rushYards)

	return Stat{
		Metric: "opp_rush_yards",
		Value:  formattedRushYards,
	}
}
