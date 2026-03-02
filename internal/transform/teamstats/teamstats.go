// Package teamstats provides team statistics transformation functionality that extracts
// and structures season-long team statistics from CBS Sports HTML pages including
// passing stats (attempts, completions, yards, touchdowns, interceptions) and
// rushing stats (attempts, yards, touchdowns, yards per rush) for both team and opponent.
package teamstats

import (
	"have-a-nice-pickem-etl/internal/extract/team"

	"github.com/PuerkitoBio/goquery"
)

/*
[

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

]
*/
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

// scrapePassingStatsTable returns the first (passing) stats table from the CBS team stats page.
func scrapePassingStatsTable(teamStatsPageSelection *goquery.Selection) *goquery.Selection {
	passingStatsTable := teamStatsPageSelection.Find("div.TableBaseWrapper").Eq(0)
	return passingStatsTable
}

// scrapeRushingStatsTable returns the second (rushing) stats table from the CBS team stats page.
func scrapeRushingStatsTable(teamStatsPageSelection *goquery.Selection) *goquery.Selection {
	rushingStatsTable := teamStatsPageSelection.Find("div.TableBaseWrapper").Eq(1)
	return rushingStatsTable
}

// scrapeStatsTableTeamTotalRow returns the team total row (first total row) from a stats table.
func scrapeStatsTableTeamTotalRow(statsTableSelection *goquery.Selection) *goquery.Selection {
	teamTotalTableRow := statsTableSelection.Find("tr.TableBase-bodyTr--total").Eq(0)
	return teamTotalTableRow
}

// scrapeStatsTableOpponentTotalRow returns the opponent total row (second total row) from a stats table.
func scrapeStatsTableOpponentTotalRow(statsTableSelection *goquery.Selection) *goquery.Selection {
	opponentTotalTableRow := statsTableSelection.Find("tr.TableBase-bodyTr--total").Eq(1)
	return opponentTotalTableRow
}

// Instantiate scrapes all team and opponent stats from the CBS team stats page and returns TeamStats.
func (t New) Instantiate() TeamStats {
	cbsPage := t.CBS

	return TeamStats{
		TeamID: t.TeamID,
		Stats: []Stat{
			teamPassAttempts{page: cbsPage}.scrape(),
			oppPassAttempts{page: cbsPage}.scrape(),
			teamPassCompletions{page: cbsPage}.scrape(),
			oppPassCompletions{page: cbsPage}.scrape(),
			teamCompletionPercentage{page: cbsPage}.scrape(),
			oppCompletionPercentage{page: cbsPage}.scrape(),
			teamPassTouchdowns{page: cbsPage}.scrape(),
			oppPassTouchdowns{page: cbsPage}.scrape(),
			teamPassYards{page: cbsPage}.scrape(),
			oppPassYards{page: cbsPage}.scrape(),
			offenseInterceptions{page: cbsPage}.scrape(),
			defenseInterceptions{page: cbsPage}.scrape(),
			teamRushAttempts{page: cbsPage}.scrape(),
			oppRushAttempts{page: cbsPage}.scrape(),
			teamRushTouchdowns{page: cbsPage}.scrape(),
			oppRushTouchdowns{page: cbsPage}.scrape(),
			teamRushYards{page: cbsPage}.scrape(),
			oppRushYards{page: cbsPage}.scrape(),
			teamYardsPerRush{page: cbsPage}.scrape(),
			oppYardsPerRush{page: cbsPage}.scrape(),
		},
	}
}
