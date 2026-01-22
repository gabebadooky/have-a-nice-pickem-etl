package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamPassAttempts() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passAttemptsTD := teamTotalsTableRow.Find("td.TableBase-bodyTd").Eq(2)
	passAttempts := strings.TrimSpace(passAttemptsTD.Text())
	var formattedPassAttempts float32 = utils.ConvertStringToFloat32(passAttempts)

	return Stat{
		Metric: "pass_attempts",
		Value:  formattedPassAttempts,
	}
}

func (t New) scrapeOpponentPassAttempts() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passAttemptsTD := opponentTotalTableRow.Find("td.TableBase-bodyTd").Eq(2)
	passAttempts := strings.TrimSpace(passAttemptsTD.Text())
	var formattedPassAttempts float32 = utils.ConvertStringToFloat32(passAttempts)

	return Stat{
		Metric: "opp_pass_attempts",
		Value:  formattedPassAttempts,
	}
}
