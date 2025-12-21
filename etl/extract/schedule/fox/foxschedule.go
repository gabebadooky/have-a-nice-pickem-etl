/*
Package that:
  - Scrapes the FOX Schedule page HTML for a given week number
*/
package fox

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type FoxSchedule interface {
	GetScheduleForWeek() *goquery.Selection
}

type FoxCfbSchedule struct {
	Week uint8
}

type FoxNflSchedule struct {
	Week uint8
}

func scrapeFoxSchedule(foxSchedulePageLink string) *goquery.Selection {
	log.Printf("\nRequesting Fox Schedule page for week: %s\n", foxSchedulePageLink)

	page, err := utils.GetGoQuerySelectionBody(foxSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Scrape Fox CFB Schedule for a given week
func (x FoxCfbSchedule) GetScheduleForWeek() *goquery.Selection {
	var foxCfbSchedulePageLink string

	if x.Week > utils.CFB_REG_SEASON_WEEKS {
		var postSeasonweek uint8 = x.Week - utils.CFB_REG_SEASON_WEEKS
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, postSeasonweek)
	} else {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULER_SEASON_SCHEDULE_URL, x.Week)
	}

	return scrapeFoxSchedule(foxCfbSchedulePageLink)
}

// Scrape Fox CFB Schedule for a given week
func (y FoxNflSchedule) GetScheduleForWeek() *goquery.Selection {
	var foxNflSchedulePageLink string

	if y.Week > utils.NFL_REG_SEASON_WEEKS {
		var week uint8 = y.Week - utils.NFL_REG_SEASON_WEEKS
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_REGULAR_SEASON_SCHEDULE_URL, y.Week)
	}

	return scrapeFoxSchedule(foxNflSchedulePageLink)
}
