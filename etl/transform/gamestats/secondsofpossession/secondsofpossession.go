package secondsofpossession

import (
	"have-a-nice-pickem-etl/etl/transform/gamestats"
	"have-a-nice-pickem-etl/etl/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type awayTeamSecondsOfPossession struct {
	page *goquery.Selection
}

type homeTeamSecondsOfPossession struct {
	page *goquery.Selection
}

func getNumberOfSecondsFromDurationString(durationString string) int {
	colonIndex := strings.Index(durationString, ":")
	minutes := durationString[:colonIndex]
	seconds := durationString[colonIndex:]
	totalSeconds := utils.ConvertStringToInt(minutes) + utils.ConvertStringToInt(seconds)
	return totalSeconds
}

func (sp awayTeamSecondsOfPossession) scrape() gamestats.Stat {
	statContainer := gamestats.ScrapePossessionStatsContainer(sp.page)
	comparisonRow := statContainer.Find("div.stats-comparison-row").Eq(0)
	statSpan := gamestats.ScrapeAwayTeamStatSpan(comparisonRow)
	statText := strings.TrimSpace(statSpan.Text())
	totalSeconds := getNumberOfSecondsFromDurationString(statText)

	return gamestats.Stat{
		Metric: "seconds_of_possession",
		Value:  float32(totalSeconds),
	}
}

func (sp homeTeamSecondsOfPossession) scrape() gamestats.Stat {
	statContainer := gamestats.ScrapePossessionStatsContainer(sp.page)
	comparisonRow := statContainer.Find("div.stats-comparison-row").Eq(0)
	statSpan := gamestats.ScrapeHomeTeamStatSpan(comparisonRow)
	statText := strings.TrimSpace(statSpan.Text())
	totalSeconds := getNumberOfSecondsFromDurationString(statText)

	return gamestats.Stat{
		Metric: "seconds_of_possession",
		Value:  float32(totalSeconds),
	}
}
