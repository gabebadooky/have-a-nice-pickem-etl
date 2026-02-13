package teamstats

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type offenseInterceptions struct {
	page *goquery.Selection
}

type defenseInterceptions struct {
	page *goquery.Selection
}

func (oi offenseInterceptions) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(oi.page)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	offensiveInterceptionsTD := teamTotalsTableRow.Find("td").Eq(7)
	offensiveInterceptions := strings.TrimSpace(offensiveInterceptionsTD.Text())
	var formattedOffensiveInterceptions float32 = utils.ConvertStringToFloat32(offensiveInterceptions)

	return Stat{
		Metric: "offense_interceptions",
		Value:  formattedOffensiveInterceptions,
	}
}

func (di defenseInterceptions) scrape() Stat {
	passingStatsTable := scrapePassingStatsTable(di.page)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	defensiveInterceptionsTD := opponentTotalTableRow.Find("td").Eq(7)
	defensiveInterceptions := strings.TrimSpace(defensiveInterceptionsTD.Text())
	var formattedDefensiveInterceptions float32 = utils.ConvertStringToFloat32(defensiveInterceptions)

	return Stat{
		Metric: "defense_interceptions",
		Value:  formattedDefensiveInterceptions,
	}
}
