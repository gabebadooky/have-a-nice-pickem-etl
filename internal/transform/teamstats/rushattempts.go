package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamRushAttempts struct {
	page *goquery.Selection
}

type oppRushAttempts struct {
	page *goquery.Selection
}

func (ra teamRushAttempts) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(ra.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(rushingStatsTable)
	rushAttemptsTD := teamTotalsTableRow.Find("td").Eq(2)
	rushAttempts := strings.TrimSpace(rushAttemptsTD.Text())
	var formattedRushAttempts float32 = utils.ConvertStringToFloat32(rushAttempts)

	return Stat{
		Metric: "rush_attempts",
		Value:  formattedRushAttempts,
	}
}

func (ra oppRushAttempts) scrape() Stat {
	rushingStatsTable := scrapeRushingStatsTable(ra.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(rushingStatsTable)
	rushAttemptsTD := opponentTotalTableRow.Find("td").Eq(2)
	rushAttempts := strings.TrimSpace(rushAttemptsTD.Text())
	var formattedRushAttempts float32 = utils.ConvertStringToFloat32(rushAttempts)

	return Stat{
		Metric: "opp_rush_attempts",
		Value:  formattedRushAttempts,
	}
}
