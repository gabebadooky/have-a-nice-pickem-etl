package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamCompletionPercentage struct {
	page *goquery.Selection
}

type oppCompletionPercentage struct {
	page *goquery.Selection
}

func (cp teamCompletionPercentage) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(cp.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	completionPercentageTD := teamTotalsTableRow.Find("td").Eq(4)
	completionPercentage := strings.TrimSpace(completionPercentageTD.Text())
	var formattedCompletionPercentage float32 = utils.ConvertStringToFloat32(completionPercentage)

	return Stat{
		Metric: "completion_percentage",
		Value:  formattedCompletionPercentage,
	}
}

func (cp oppCompletionPercentage) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(cp.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	completionPercentageTD := opponentTotalTableRow.Find("td").Eq(4)
	completionPercentage := strings.TrimSpace(completionPercentageTD.Text())
	var formattedCompletionPercentage float32 = utils.ConvertStringToFloat32(completionPercentage)

	return Stat{
		Metric: "opp_completion_percentage",
		Value:  formattedCompletionPercentage,
	}
}
