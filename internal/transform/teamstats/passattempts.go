// Package teamstats provides passing attempts statistics scraping functionality.
// It extracts pass attempts data from CBS Sports team stats pages for both
// team and opponent statistics.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type teamPassAttempts struct {
	page *goquery.Selection
}

type oppPassAttempts struct {
	page *goquery.Selection
}

// scrape extracts team pass attempts from the CBS passing stats table.
func (pa teamPassAttempts) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pa.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	passAttemptsTD := teamTotalsTableRow.Find("td.TableBase-bodyTd").Eq(2)
	passAttempts := strings.TrimSpace(passAttemptsTD.Text())
	var formattedPassAttempts float32 = utils.ConvertStringToFloat32(passAttempts)

	return Stat{
		Metric: "pass_attempts",
		Value:  formattedPassAttempts,
	}
}

// scrape extracts opponent pass attempts from the CBS passing stats table.
func (pa oppPassAttempts) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(pa.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	passAttemptsTD := opponentTotalTableRow.Find("td.TableBase-bodyTd").Eq(2)
	passAttempts := strings.TrimSpace(passAttemptsTD.Text())
	var formattedPassAttempts float32 = utils.ConvertStringToFloat32(passAttempts)

	return Stat{
		Metric: "opp_pass_attempts",
		Value:  formattedPassAttempts,
	}
}
