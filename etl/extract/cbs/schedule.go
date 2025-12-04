/*
Package that:
  - Scrapes the CBS Schedule page HTML for a given week number
*/
package cbs

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/utils"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type CbsCFBSchedule struct {
	Week uint8
}

type CbsNFLSchedule struct {
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

// Scrape Cbs CFB Schedule for a given week
func (x CbsCFBSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsCfbSchedulePageLink string

	if x.Week > utils.CFB_REG_SEASON_WEEKS {
		cbsCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_POST_SEASON_SCHEDULE_URL, x.Week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", x.Week, cbsCfbSchedulePageLink)
	} else {
		cbsCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_CFB_REGULAR_SEASON_SCHEDULE_URL, x.Week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", x.Week, cbsCfbSchedulePageLink)
	}

	page, err := extract.GetSchedulePageBody(cbsCfbSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Scrape Cbs CFB Schedule for a given week
func (y CbsNFLSchedule) GetScheduleForWeek() *goquery.Selection {
	var cbsNflSchedulePageLink string

	if y.Week > utils.NFL_REG_SEASON_WEEKS {
		var week string = setNflPostseasonWeek(y.Week)
		cbsNflSchedulePageLink = fmt.Sprintf("%s%s", utils.CBS_NFL_POST_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", y.Week, cbsNflSchedulePageLink)
	} else {
		cbsNflSchedulePageLink = fmt.Sprintf("%s%d", utils.CBS_NFL_REGULAR_SEASON_SCHEDULE_URL, y.Week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", y.Week, cbsNflSchedulePageLink)
	}

	page, err := extract.GetSchedulePageBody(cbsNflSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}
