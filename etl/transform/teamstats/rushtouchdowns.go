package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeTeamRushTouchdowns() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushTouchdownsTD := teamTotalsTableRow.Find("td").Eq(5)
	rushTouchdowns := strings.TrimSpace(rushTouchdownsTD.Text())
	var formattedRushTouchdowns float32 = utils.ConvertStringToFloat32(rushTouchdowns)

	return Stat{
		Metric: "rush_touchdowns",
		Value:  formattedRushTouchdowns,
	}
}

func (t New) scrapeOpponentRushTouchdowns() Stat {
	rushingStatsTable := scrapeRushingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushTouchdownsTD := opponentTotalTableRow.Find("td").Eq(5)
	rushTouchdowns := strings.TrimSpace(rushTouchdownsTD.Text())
	var formattedRushTouchdowns float32 = utils.ConvertStringToFloat32(rushTouchdowns)

	return Stat{
		Metric: "opp_rush_touchdowns",
		Value:  formattedRushTouchdowns,
	}
}
