package teamstats

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t New) scrapeOffenseiveInterceptions() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var teamTotalsTableRow *goquery.Selection = scrapeStatsTableTeamTotalRow(passingStatsTable)
	offensiveInterceptionsTD := teamTotalsTableRow.Find("td").Eq(7)
	offensiveInterceptions := strings.TrimSpace(offensiveInterceptionsTD.Text())
	var formattedOffensiveInterceptions float32 = utils.ConvertStringToFloat32(offensiveInterceptions)

	return Stat{
		Metric: "offense_interceptions",
		Value:  formattedOffensiveInterceptions,
	}
}

func (t New) scrapeDefensiveInterceptions() Stat {
	passingStatsTable := scrapePassingStatsTable(t.CBS)
	var opponentTotalTableRow *goquery.Selection = scrapeStatsTableOpponentTotalRow(passingStatsTable)
	defensiveInterceptionsTD := opponentTotalTableRow.Find("td").Eq(7)
	defensiveInterceptions := strings.TrimSpace(defensiveInterceptionsTD.Text())
	var formattedDefensiveInterceptions float32 = utils.ConvertStringToFloat32(defensiveInterceptions)

	return Stat{
		Metric: "defense_interceptions",
		Value:  formattedDefensiveInterceptions,
	}
}
