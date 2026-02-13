// Package teamstats provides team statistics transformation functionality that extracts
// and structures season-long team statistics from CBS Sports HTML pages including
// passing stats (attempts, completions, yards, touchdowns, interceptions) and
// rushing stats (attempts, yards, touchdowns, yards per rush) for both team and opponent.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/extract/team"

	"github.com/PuerkitoBio/goquery"
)

/*[
	"pass_attempts", "opp_pass_attempts",
	"pass_completions", "opp_pass_completions",
	"completion_percentage", "opp_completion_percentage",
	"pass_yards", "opp_pass_yards",
	"pass_touchdowns", "opp_pass_touchdowns",
	"offense_interceptions", "defense_interceptions",
	"rush_yards", "opp_rush_yards",
	"rush_attempts", "opp_rush_attempts",
	"yards_per_rush", "opp_yards_per_rush",
	"rush_touchdowns", "opp_rush_touchdowns"
]*/

type scraper interface {
	scrape() Stat
}

type New struct {
	team.Team
}

type Stat struct {
	Metric string
	Value  float32
}

type TeamStats struct {
	TeamID string
	Stats  []Stat
}

func ScrapeStat(s scraper) Stat {
	return s.scrape()
}

func scrapePassingStatsTable(teamStatsPageSelection *goquery.Selection) *goquery.Selection {
	passingStatsTable := teamStatsPageSelection.Find("div.TableBaseWrapper").Eq(0)
	return passingStatsTable
}

func scrapeRushingStatsTable(teamStatsPageSelection *goquery.Selection) *goquery.Selection {
	rushingStatsTable := teamStatsPageSelection.Find("div.TableBaseWrapper").Eq(1)
	return rushingStatsTable
}

func scrapeStatsTableTeamTotalRow(statsTableSelection *goquery.Selection) *goquery.Selection {
	teamTotalTableRow := statsTableSelection.Find("tr.TableBase-bodyTr--total").Eq(0)
	return teamTotalTableRow
}

func scrapeStatsTableOpponentTotalRow(statsTableSelection *goquery.Selection) *goquery.Selection {
	opponentTotalTableRow := statsTableSelection.Find("tr.TableBase-bodyTr--total").Eq(1)
	return opponentTotalTableRow
}

func (t New) Instantiate() TeamStats {
	/*teamPassAttempts := t.scrapeTeamPassAttempts()
	oppPassAttempts := t.scrapeOpponentPassAttempts()
	teamCompletionPercentage := t.scrapeTeamCompletionPercentage()
	oppCompletionPercentage := t.scrapeOpponentCompletionPercentage()
	teamPassYards := t.scrapeTeamPassYards()
	oppPassYards := t.scrapeOpponentPassYards()
	teamPassTouchdowns := t.scrapeTeamPassTouchdowns()
	oppPassTouchdowns := t.scrapeOpponentPassTouchdowns()
	offensiveInterceptions := t.scrapeOffenseiveInterceptions()
	defensiveInterceptions := t.scrapeDefensiveInterceptions()
	teamRushYards := t.scrapeTeamRushYards()
	oppRushYards := t.scrapeOpponentRushYards()
	teamRushAttempts := t.scrapeTeamRushAttempts()
	oppRushAttempts := t.scrapeOpponentRushAttempts()
	teamYardsPerRush := t.scrapeTeamYardsPerRush()
	oppYardsPerRush := t.scrapeOpponentYardsPerRush()
	teamRushTouchdowns := t.scrapeTeamRushTouchdowns()
	oppRushTouchdowns := t.scrapeOpponentRushTouchdowns()*/
	cbsPage := t.CBS

	return TeamStats{
		TeamID: t.TeamID,
		Stats: []Stat{
			ScrapeStat(teamPassAttempts{page: cbsPage}),
			ScrapeStat(oppPassAttempts{page: cbsPage}),
			ScrapeStat(teamPassCompletions{page: cbsPage}),
			ScrapeStat(oppPassCompletions{page: cbsPage}),
			ScrapeStat(teamCompletionPercentage{page: cbsPage}),
			ScrapeStat(oppCompletionPercentage{page: cbsPage}),
			ScrapeStat(teamPassTouchdowns{page: cbsPage}),
			ScrapeStat(oppPassTouchdowns{page: cbsPage}),
			ScrapeStat(teamPassYards{page: cbsPage}),
			ScrapeStat(oppPassYards{page: cbsPage}),
			ScrapeStat(offenseInterceptions{page: cbsPage}),
			ScrapeStat(defenseInterceptions{page: cbsPage}),
			ScrapeStat(teamRushAttempts{page: cbsPage}),
			ScrapeStat(oppRushAttempts{page: cbsPage}),
			ScrapeStat(teamRushTouchdowns{page: cbsPage}),
			ScrapeStat(oppRushTouchdowns{page: cbsPage}),
			ScrapeStat(teamRushYards{page: cbsPage}),
			ScrapeStat(oppRushYards{page: cbsPage}),
			ScrapeStat(teamYardsPerRush{page: cbsPage}),
			ScrapeStat(oppYardsPerRush{page: cbsPage}),
		},
	}
}
