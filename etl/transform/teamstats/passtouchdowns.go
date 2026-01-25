package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamPassTouchdowns struct {
	page *goquery.Selection
}

type oppPassTouchdowns struct {
	page *goquery.Selection
}

func (pt teamPassTouchdowns) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pt.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passTouchdownsTD := teamTotalsTableRow.Find("td").Eq(6)
	passTouchdowns := strings.TrimSpace(passTouchdownsTD.Text())
	var formattedPassTouchdowns float32 = utils.ConvertStringToFloat32(passTouchdowns)

	return Stat{
		Metric: "pass_touchdowns",
		Value:  formattedPassTouchdowns,
	}
}

func (pt oppPassTouchdowns) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pt.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passTouchdownsTD := opponentTotalTableRow.Find("td").Eq(6)
	passTouchdowns := strings.TrimSpace(passTouchdownsTD.Text())
	var formattedPassTouchdowns float32 = utils.ConvertStringToFloat32(passTouchdowns)

	return Stat{
		Metric: "opp_pass_touchdowns",
		Value:  formattedPassTouchdowns,
	}
}
