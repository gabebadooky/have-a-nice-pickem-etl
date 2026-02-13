package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamPassCompletions struct {
	page *goquery.Selection
}

type oppPassCompletions struct {
	page *goquery.Selection
}

func (pc teamPassCompletions) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pc.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passCompletionsTD := teamTotalsTableRow.Find("td").Eq(4)
	passCompletions := strings.TrimSpace(passCompletionsTD.Text())
	var formattedPassCompletions float32 = utils.ConvertStringToFloat32(passCompletions)

	return Stat{
		Metric: "pass_completions",
		Value:  formattedPassCompletions,
	}
}

func (pc oppPassCompletions) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pc.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passCompletionsTD := opponentTotalTableRow.Find("td").Eq(4)
	passCompletions := strings.TrimSpace(passCompletionsTD.Text())
	var formattedPassCompletions float32 = utils.ConvertStringToFloat32(passCompletions)

	return Stat{
		Metric: "opp_pass_completions",
		Value:  formattedPassCompletions,
	}
}
