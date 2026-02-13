// Package cbsschedule provides CBS Sports schedule web scraping functionality.
// It extracts schedule data from CBS Sports HTML pages for both college football (CFB)
// and NFL, handling both regular season and postseason schedules.
package cbsschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type CbsCfbSchedule struct {
	Week uint
}

type CbsNflSchedule struct {
	Week uint
}
type cbsScheduleInstantiator interface {
	scrapeSchedule() *goquery.Selection
}

func GetScheduleForWeek(s cbsScheduleInstantiator) *goquery.Selection {
	return s.scrapeSchedule()
}

func setNflPostseasonWeekValue(week uint) string {
	switch week - utils.NFL_REG_SEASON_WEEKS {
	case 1:
		return "wild-card"
	case 2:
		return "divisional"
	case 3:
		return "championship"
	case 4:
		return "pro-bowl"
	default:
		return "super-bowl"
	}
}

// Make and handle CBS Schedule web scrape attempt
func fetchCbsSchedule(cbsSchedulePageLink string) *goquery.Selection {
	log.Printf("\nRequesting CBS Schedule page for: %s\n", cbsSchedulePageLink)

	page, err := utils.GetGoQuerySelectionBody(cbsSchedulePageLink)

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

func (sched CbsCfbSchedule) instantiateShedulePageLink() string {
	var cbsSchedulePageLink string

	if sched.Week > utils.CFB_REG_SEASON_WEEKS {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_POST_SEASON_SCHEDULE_URL, sched.Week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	}

	return cbsSchedulePageLink
}

func (sched CbsNflSchedule) instantiateShedulePageLink() string {
	var cbsSchedulePageLink string

	if sched.Week > utils.NFL_REG_SEASON_WEEKS {
		var week string = setNflPostseasonWeekValue(sched.Week)
		cbsSchedulePageLink = fmt.Sprintf("%s%s", utils.CBS_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_NFL_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	}

	return cbsSchedulePageLink
}

// Scrape Cbs CFB Schedule for a given week
func (sched CbsCfbSchedule) scrapeSchedule() *goquery.Selection {
	cbsSchedulePageLink := sched.instantiateShedulePageLink()
	cbsSchedule := fetchCbsSchedule(cbsSchedulePageLink)
	return cbsSchedule
}

// Scrape Cbs NFL Schedule for a given week
func (sched CbsNflSchedule) scrapeSchedule() *goquery.Selection {
	cbsSchedulePageLink := sched.instantiateShedulePageLink()
	cbsSchedule := fetchCbsSchedule(cbsSchedulePageLink)
	return cbsSchedule
}
