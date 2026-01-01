/*
Package that:
  - Scrapes the FOX Schedule page HTML for a given week number
*/
package foxschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type FoxSchedule interface {
	scheduleForWeek() *goquery.Selection
}

type FoxCfbSchedule struct {
	Week uint
}

type FoxNflSchedule struct {
	Week uint
}

func GetScheduleForWeek(s FoxSchedule) *goquery.Selection {
	return s.scheduleForWeek()
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
func (x FoxCfbSchedule) scheduleForWeek() *goquery.Selection {
	var foxCfbSchedulePageLink string

	if x.Week > utils.CFB_REG_SEASON_WEEKS {
		postSeasonweek := x.Week - utils.CFB_REG_SEASON_WEEKS
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, postSeasonweek)
	} else {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULER_SEASON_SCHEDULE_URL, x.Week)
	}

	foxSchedule := scrapeFoxSchedule(foxCfbSchedulePageLink)
	return foxSchedule
}

// Scrape Fox CFB Schedule for a given week
func (y FoxNflSchedule) scheduleForWeek() *goquery.Selection {
	var foxNflSchedulePageLink string

	if y.Week > utils.NFL_REG_SEASON_WEEKS {
		week := y.Week - utils.NFL_REG_SEASON_WEEKS
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_REGULAR_SEASON_SCHEDULE_URL, y.Week)
	}

	foxSchedule := scrapeFoxSchedule(foxNflSchedulePageLink)
	return foxSchedule
}
