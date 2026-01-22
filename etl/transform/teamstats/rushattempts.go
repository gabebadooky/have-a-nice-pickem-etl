package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamRushAttempts() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushAttemptsTD := teamTotalsTableRow.Find("td").Eq(2)
	rushAttempts := strings.TrimSpace(rushAttemptsTD.Text())
	var formattedRushAttempts float32 = utils.ConvertStringToFloat32(rushAttempts)

	return Stat{
		Metric: "rush_attempts",
		Value:  formattedRushAttempts,
	}
}

func (t New) scrapeOpponentRushAttempts() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushAttemptsTD := opponentTotalTableRow.Find("td").Eq(2)
	rushAttempts := strings.TrimSpace(rushAttemptsTD.Text())
	var formattedRushAttempts float32 = utils.ConvertStringToFloat32(rushAttempts)

	return Stat{
		Metric: "opp_rush_attempts",
		Value:  formattedRushAttempts,
	}
}
