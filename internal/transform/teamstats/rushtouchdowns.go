package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamRushTouchdowns struct {
	page *goquery.Selection
}

type oppRushTouchdowns struct {
	page *goquery.Selection
}

func (rt teamRushTouchdowns) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(rt.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushTouchdownsTD := teamTotalsTableRow.Find("td").Eq(5)
	rushTouchdowns := strings.TrimSpace(rushTouchdownsTD.Text())
	var formattedRushTouchdowns float32 = utils.ConvertStringToFloat32(rushTouchdowns)

	return Stat{
		Metric: "rush_touchdowns",
		Value:  formattedRushTouchdowns,
	}
}

func (rt oppRushTouchdowns) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(rt.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushTouchdownsTD := opponentTotalTableRow.Find("td").Eq(5)
	rushTouchdowns := strings.TrimSpace(rushTouchdownsTD.Text())
	var formattedRushTouchdowns float32 = utils.ConvertStringToFloat32(rushTouchdowns)

	return Stat{
		Metric: "opp_rush_touchdowns",
		Value:  formattedRushTouchdowns,
	}
}
