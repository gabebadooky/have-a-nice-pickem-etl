// Package teamstats provides passing yards statistics scraping functionality.
// It extracts passing yards data from CBS Sports team stats pages for both
// team and opponent statistics.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamPassYards struct {
	page *goquery.Selection
}

type oppPassYards struct {
	page *goquery.Selection
}

// scrape extracts team pass yards from the CBS passing stats table.
func (py teamPassYards) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(py.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passYardsTD := teamTotalsTableRow.Find("td").Eq(5)
	passYards := strings.TrimSpace(passYardsTD.Text())
	var formattedPassYards float32 = utils.ConvertStringToFloat32(passYards)

	return Stat{
		Metric: "pass_yards",
		Value:  formattedPassYards,
	}
}

// scrape extracts opponent pass yards from the CBS passing stats table.
func (py oppPassYards) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(py.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passYardsTD := opponentTotalTableRow.Find("td").Eq(5)
	passYards := strings.TrimSpace(passYardsTD.Text())
	var formattedPassYards float32 = utils.ConvertStringToFloat32(passYards)

	return Stat{
		Metric: "opp_pass_yards",
		Value:  formattedPassYards,
	}
}
