/*
Package that:
  - Scrapes the CBS Schedule page HTML for a given week number
*/
package cbs

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type CbsSchedule interface {
	GetScheduleForWeek() *goquery.Selection
}

type CbsCfbSchedule struct {
	Week uint8
}

type CbsNflSchedule struct {
	Week uint8
}

func setNflPostseasonWeek(week uint8) string {
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
func (x CbsCfbSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsSchedulePageLink string

	if x.Week > utils.CFB_REG_SEASON_WEEKS {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_POST_SEASON_SCHEDULE_URL, x.Week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_REGULAR_SEASON_SCHEDULE_URL, x.Week)
	}

	return scrapeCbsSchedule(cbsSchedulePageLink)
}

// Scrape Cbs CFB Schedule for a given week
func (y CbsNflSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsSchedulePageLink string

	if y.Week > utils.NFL_REG_SEASON_WEEKS {
		var week string = setNflPostseasonWeek(y.Week)
		cbsSchedulePageLink = fmt.Sprintf("%s%s", utils.CBS_NFL_POST_SEASON_SCHEDULE_URL, week)
	} else {
		cbsSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_NFL_REGULAR_SEASON_SCHEDULE_URL, y.Week)
	}

	return scrapeCbsSchedule(cbsSchedulePageLink)
}
