package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamPassTouchdowns() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passTouchdownsTD := teamTotalsTableRow.Find("td").Eq(6)
	passTouchdowns := strings.TrimSpace(passTouchdownsTD.Text())
	var formattedPassTouchdowns float32 = utils.ConvertStringToFloat32(passTouchdowns)

	return Stat{
		Metric: "pass_touchdowns",
		Value:  formattedPassTouchdowns,
	}
}

func (t New) scrapeOpponentPassTouchdowns() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passTouchdownsTD := opponentTotalTableRow.Find("td").Eq(6)
	passTouchdowns := strings.TrimSpace(passTouchdownsTD.Text())
	var formattedPassTouchdowns float32 = utils.ConvertStringToFloat32(passTouchdowns)

	return Stat{
		Metric: "opp_pass_touchdowns",
		Value:  formattedPassTouchdowns,
	}
}
