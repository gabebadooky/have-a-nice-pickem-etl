package gamestats

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Instantiator interface {
	instantiate() GameStats
}

type AwayTeamStat struct {
	game.Game
}

type HomeTeamStat struct {
	game.Game
}

/*
TODO:
- Develop methods to scrape following stats from Fox Boxscore Page (Away and Home teams):
	- Kick Returns: (6th stat table)
	- Kick Return Yards: (6th stat table)

	- Punt Returns: 7th stat table
	- Punt Return Yards: 7th stat table

	- Field Goals Made: 8th stat table
	- Field Goals Attempted: 8th stat table
	- Extra Points Made: 8th stat table
	- Extra Points Attempted: 8th stat table

	- Punts: 9th stat table
*/

type Stat struct {
	Metric string
	Value  float32
}

type GameStats struct {
	GameID string
	TeamID string
	Stats  []Stat
}

func InstantiateGameStats(i Instantiator) GameStats {
	return i.instantiate()
}

const awayStatSpanIndex int = 0
const homeStatSpanIndex int = 1

var possessionStatsTableIndices map[int]string = map[int]string{
	0: "seconds_of_possession",
	1: "total_drives",
	2: "total_plays",
	3: "total_yards",
	4: "yards_per_play",
	5: "redzone_touchdowns",
	6: "redzone_attempts",
}

var passingStatsTableIndices map[int]string = map[int]string{
	0: "total_passing_yards",
	1: "passing_completions",
	2: "passing_attempts",
	3: "yards_per_pass",
	4: "passing_touchdowns",
}

var rushingStatsTableIndices map[int]string = map[int]string{
	0: "total_rushing_yards",
	1: "rushing_attempts",
	2: "yards_per_rush",
	3: "rushing_touchdowns",
}

var defenseStatsTableIndices map[int]string = map[int]string{
	0: "sacks",
	1: "tackles_for_loss",
	2: "pass_deflections",
	//3: "yards_per_rush",
}

var turnoverStatsTableIndices map[int]string = map[int]string{
	0: "total_turnovers",
	1: "fumbles_lost",
	2: "interceptions",
}

var statTableIndices map[int]map[int]string = map[int]map[int]string{
	0: possessionStatsTableIndices,
	1: passingStatsTableIndices,
	2: rushingStatsTableIndices,
	3: defenseStatsTableIndices,
	4: turnoverStatsTableIndices,
}

func getNumberOfSecondsFromDurationString(durationString string) int {
	before, after, _ := strings.Cut(durationString, ":")
	minutes := before
	seconds := after
	fmt.Printf("\nminutes: %s | seconds: %s\n", minutes, seconds)
	totalSeconds := (utils.ConvertStringToInt(minutes) * 60) + utils.ConvertStringToInt(seconds)
	return totalSeconds
}

func scrapeStatContainerRow(gameStatsPageSelection *goquery.Selection, statTableIndex int, statIndex int) *goquery.Selection {
	gameStatsContainer := gameStatsPageSelection.Find("div.event-stats-container")
	possessionStatsContainer := gameStatsContainer.Find("div.stats-comparison-container").Eq(statTableIndex).Find("div.stats-team-comparison")
	statComparisonRow := possessionStatsContainer.Find("div.stats-comparison-row").Eq(statIndex)
	return statComparisonRow
}

func scrapeStat(GameStatsPageSelection *goquery.Selection, statTableIndex int, statIndex int, statComparisonRowSpanIndex int) Stat {
	var statFloat float32

	statComparisonRow := scrapeStatContainerRow(GameStatsPageSelection, statTableIndex, statIndex)
	statSpan := statComparisonRow.Find("span.comparison-data").Eq(statComparisonRowSpanIndex)
	statText := strings.TrimSpace(statSpan.Text())
	statType := statTableIndices[statTableIndex][statIndex]

	if statTableIndex == 0 && statIndex == 0 {
		totalSeconds := getNumberOfSecondsFromDurationString(statText)
		statFloat = float32(totalSeconds)
	} else {
		statFloat = utils.ConvertStringToFloat32(statText)
	}

	return Stat{
		Metric: statType,
		Value:  statFloat,
	}
}

/*
func (a AwayTeamStat) ScrapeStat() Stat {
	statComparisonRow := scrapeStatContainerRow(a.GameStatsPageSelection, a.StatTableIndex, a.StatIndex)
	statSpan := statComparisonRow.Find("span.comparison-data").Eq(0)
	statText := strings.TrimSpace(statSpan.Text())
	statFloat := utils.ConvertStringToFloat32(statText)
	statType := statTableIndices[a.StatTableIndex][a.StatIndex]

	return Stat{
		Metric: statType,
		Value:  statFloat,
	}
}

func (h HomeTeamStat) ScrapeStat() Stat {
	statComparisonRow := scrapeStatContainerRow(h.GameStatsPageSelection, h.StatTableIndex, h.StatIndex)
	statSpan := statComparisonRow.Find("span.comparison-data").Eq(1)
	statText := strings.TrimSpace(statSpan.Text())
	statFloat := utils.ConvertStringToFloat32(statText)
	statType := statTableIndices[h.StatTableIndex][h.StatIndex]

	return Stat{
		Metric: statType,
		Value:  statFloat,
	}
}
*/

func (a AwayTeamStat) instantiate() GameStats {
	var statSlice []Stat
	var teamID string = common.ParseAwayTeamID(a.Game)

	for statTableIndex, statTable := range statTableIndices {
		for statIndex := range statTable {
			awayStat := scrapeStat(a.FOX.StatsPage, statTableIndex, statIndex, awayStatSpanIndex)
			statSlice = append(statSlice, awayStat)
		}
	}

	return GameStats{
		GameID: a.GameID,
		TeamID: teamID,
		Stats:  statSlice,
	}
}

func (h HomeTeamStat) instantiate() GameStats {
	var statSlice []Stat
	var teamID string = common.ParseAwayTeamID(h.Game)

	for statTableIndex, statTable := range statTableIndices {
		for statIndex := range statTable {
			awayStat := scrapeStat(h.FOX.StatsPage, statTableIndex, statIndex, homeStatSpanIndex)
			statSlice = append(statSlice, awayStat)
		}
	}

	return GameStats{
		GameID: h.GameID,
		TeamID: teamID,
		Stats:  statSlice,
	}
}
