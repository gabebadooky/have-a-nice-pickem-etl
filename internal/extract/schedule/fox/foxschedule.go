// Package foxschedule provides Fox Sports schedule web scraping functionality.
// It extracts schedule data from Fox Sports HTML pages for both college football (CFB)
// and NFL, handling both regular season and postseason schedules.
package foxschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type FoxCfbSchedule struct {
	Week uint
}

type FoxNflSchedule struct {
	Week uint
}

// fetchFoxSchedule fetches the Fox schedule page at the given URL and returns its body as a goquery selection.
func fetchFoxSchedule(foxSchedulePageLink string) *goquery.Selection {
	log.Printf("\nRequesting Fox Schedule page for week: %s\n", foxSchedulePageLink)

	page, err := utils.GetGoQuerySelectionBody(foxSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// instantiateShedulePageLink returns the Fox college football schedule URL for the configured week.
func (sched FoxCfbSchedule) instantiateShedulePageLink() string {
	var foxSchedulePageLink string

	if sched.Week > utils.CFB_REG_SEASON_WEEKS {
		foxSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, 1)
	} else {
		foxSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	}

	return foxSchedulePageLink
}

// instantiateShedulePageLink returns the Fox NFL schedule URL for the configured week.
func (sched FoxNflSchedule) instantiateShedulePageLink() string {
	var foxSchedulePageLink string

	if sched.Week > utils.NFL_REG_SEASON_WEEKS {
		week := sched.Week - utils.NFL_REG_SEASON_WEEKS
		foxSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		foxSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	}

	return foxSchedulePageLink
}

// scrapeSchedule fetches the Fox college football schedule page for the configured week.
func (sched FoxCfbSchedule) GetSchedule() *goquery.Selection {
	foxSchedulePageLink := sched.instantiateShedulePageLink()
	foxSchedule := fetchFoxSchedule(foxSchedulePageLink)
	return foxSchedule
}

// scrapeSchedule fetches the Fox NFL schedule page for the configured week.
func (sched FoxNflSchedule) GetSchedule() *goquery.Selection {
	foxSchedulePageLink := sched.instantiateShedulePageLink()
	foxSchedule := fetchFoxSchedule(foxSchedulePageLink)
	return foxSchedule
}
