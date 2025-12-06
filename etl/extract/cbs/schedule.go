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

type CbsCFBSchedule struct {
	week uint8
}

type CbsNFLSchedule struct {
	week uint8
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

// Scrape Cbs CFB Schedule for a given week
func (x CbsCFBSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsCfbSchedulePageLink string

	if x.week > utils.CFB_REG_SEASON_WEEKS {
		cbsCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_POST_SEASON_SCHEDULE_URL, x.week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", x.week, cbsCfbSchedulePageLink)
	} else {
		cbsCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_REGULAR_SEASON_SCHEDULE_URL, x.week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", x.week, cbsCfbSchedulePageLink)
	}

	page, err := utils.GetGoQuerySelectionBody(cbsCfbSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Scrape Cbs CFB Schedule for a given week
func (y CbsNFLSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsNflSchedulePageLink string

	if y.week > utils.NFL_REG_SEASON_WEEKS {
		var week string = setNflPostseasonWeek(y.week)
		cbsNflSchedulePageLink = fmt.Sprintf("%s%s", utils.CBS_NFL_POST_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", y.week, cbsNflSchedulePageLink)
	} else {
		cbsNflSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_NFL_REGULAR_SEASON_SCHEDULE_URL, y.week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", y.week, cbsNflSchedulePageLink)
	}

	page, err := utils.GetGoQuerySelectionBody(cbsNflSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}
