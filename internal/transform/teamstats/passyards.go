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
