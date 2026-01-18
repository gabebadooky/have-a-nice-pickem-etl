package cbsschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type CbsSchedule interface {
	scheduleForWeek() *goquery.Selection
}

type CbsCfbSchedule struct {
	Week uint
}

type CbsNflSchedule struct {
	Week uint
}

func GetScheduleForWeek(s CbsSchedule) *goquery.Selection {
	return s.scheduleForWeek()
}

func setNflPostseasonWeek(week uint) string {
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

func scrapeCbsSchedule(cbsSchedulePageLink string) *goquery.Selection {
	log.Printf("\nRequesting CBS Schedule page for: %s\n", cbsSchedulePageLink)
	page, err := utils.GetGoQuerySelectionBody(cbsSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}
	return page
}

// Scrape Cbs CFB Schedule for a given week
func (x CbsCfbSchedule) scheduleForWeek() *goquery.Selection {
	var cbsSchedulePageLink string

	if x.Week > utils.CFB_REG_SEASON_WEEKS {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_POST_SEASON_SCHEDULE_URL, x.Week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_REGULAR_SEASON_SCHEDULE_URL, x.Week)
	}

	cbsSchedule := scrapeCbsSchedule(cbsSchedulePageLink)
	return cbsSchedule
}

// Scrape Cbs NFL Schedule for a given week
func (y CbsNflSchedule) scheduleForWeek() *goquery.Selection {
	var cbsSchedulePageLink string

	if y.Week > utils.NFL_REG_SEASON_WEEKS {
		var week string = setNflPostseasonWeek(y.Week)
		cbsSchedulePageLink = fmt.Sprintf("%s%s", utils.CBS_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_NFL_REGULAR_SEASON_SCHEDULE_URL, y.Week)
	}

	cbsSchedule := scrapeCbsSchedule(cbsSchedulePageLink)
	return cbsSchedule
}
