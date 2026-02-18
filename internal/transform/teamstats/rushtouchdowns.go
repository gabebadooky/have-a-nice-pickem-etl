// Package teamstats provides rushing touchdowns statistics scraping functionality.
// It extracts rushing touchdowns data from CBS Sports team stats pages for both
// team and opponent statistics.
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

// scrape extracts team rush touchdowns from the CBS rushing stats table.
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

// scrape extracts opponent rush touchdowns from the CBS rushing stats table.
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
