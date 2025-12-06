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

type FoxCFBSchedule struct {
	week uint8
}

type FoxNFLSchedule struct {
	week uint8
}

// Scrape Fox CFB Schedule for a given week
func (x FoxCFBSchedule) GetScheduleForWeek() *goquery.Selection {
	var foxCfbSchedulePageLink string

	if x.week > utils.CFB_REG_SEASON_WEEKS {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, x.week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", x.week, foxCfbSchedulePageLink)
	} else {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULER_SEASON_SCHEDULE_URL, x.week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", x.week, foxCfbSchedulePageLink)
	}

	page, err := utils.GetGoQuerySelectionBody(foxCfbSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Scrape Fox CFB Schedule for a given week
func (y FoxNFLSchedule) GetScheduleForWeek() *goquery.Selection {
	var foxNflSchedulePageLink string

	if y.week > utils.NFL_REG_SEASON_WEEKS {
		var week uint8 = y.week - utils.NFL_REG_SEASON_WEEKS
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_POST_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", y.week, foxNflSchedulePageLink)
	} else {
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_NFL_REGULAR_SEASON_SCHEDULE_URL, y.week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", y.week, foxNflSchedulePageLink)
	}

	page, err := utils.GetGoQuerySelectionBody(foxNflSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}
