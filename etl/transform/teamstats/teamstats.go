package teamstats

import (
	"have-a-nice-pickem-etl/etl/extract/team"

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
	teamPassAttempts := t.scrapeTeamPassAttempts()
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
	oppRushTouchdowns := t.scrapeOpponentRushTouchdowns()

	return TeamStats{
		TeamID: t.TeamID,
		Stats: []Stat{
			teamPassAttempts,
			oppPassAttempts,
			teamCompletionPercentage,
			oppCompletionPercentage,
			teamPassYards,
			oppPassYards,
			teamPassTouchdowns,
			oppPassTouchdowns,
			offensiveInterceptions,
			defensiveInterceptions,
			teamRushYards,
			oppRushYards,
			teamRushAttempts,
			oppRushAttempts,
			teamYardsPerRush,
			oppYardsPerRush,
			teamRushTouchdowns,
			oppRushTouchdowns,
		},
	}
}
