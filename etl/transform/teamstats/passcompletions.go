package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamPassCompletions() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passCompletionsTD := teamTotalsTableRow.Find("td").Eq(4)
	passCompletions := strings.TrimSpace(passCompletionsTD.Text())
	var formattedPassCompletions float32 = utils.ConvertStringToFloat32(passCompletions)

	return Stat{
		Metric: "pass_completions",
		Value:  formattedPassCompletions,
	}
}

func (t New) scrapeOpponentPassCompletions() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passCompletionsTD := opponentTotalTableRow.Find("td").Eq(4)
	passCompletions := strings.TrimSpace(passCompletionsTD.Text())
	var formattedPassCompletions float32 = utils.ConvertStringToFloat32(passCompletions)

	return Stat{
		Metric: "opp_pass_completions",
		Value:  formattedPassCompletions,
	}
}
