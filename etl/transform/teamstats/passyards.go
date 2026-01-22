package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamPassYards() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passYardsTD := teamTotalsTableRow.Find("td").Eq(5)
	passYards := strings.TrimSpace(passYardsTD.Text())
	var formattedPassYards float32 = utils.ConvertStringToFloat32(passYards)

	return Stat{
		Metric: "pass_yards",
		Value:  formattedPassYards,
	}
}

func (t New) scrapeOpponentPassYards() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passYardsTD := opponentTotalTableRow.Find("td").Eq(5)
	passYards := strings.TrimSpace(passYardsTD.Text())
	var formattedPassYards float32 = utils.ConvertStringToFloat32(passYards)

	return Stat{
		Metric: "opp_pass_yards",
		Value:  formattedPassYards,
	}
}
