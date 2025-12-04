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

type FoxCFB struct {
}

type FoxNFL struct {
}

// Scrape Fox Schedule for a given week
func getSchedule(schedulePageLink string) (*goquery.Selection, error) {
	doc, err := utils.ScrapePage(schedulePageLink)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	//log.Printf("htmlBody:\n%v\n", htmlbody)
	return htmlbody, nil
}

// Scrape Fox CFB Schedule for a given week
func (FoxCFB) GetScheduleForWeek(week uint8) *goquery.Selection {
	var foxCfbSchedulePageLink string

	if week > utils.CFB_REG_SEASON_WEEKS {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", week, foxCfbSchedulePageLink)
	} else {
		foxCfbSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULER_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", week, foxCfbSchedulePageLink)
	}

	page, err := getSchedule(foxCfbSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}

// Scrape Fox CFB Schedule for a given week
func (FoxNFL) GetScheduleForWeek(week uint8) *goquery.Selection {
	var foxNflSchedulePageLink string

	if week > utils.CFB_REG_SEASON_WEEKS {
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_POST_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for post season week %d: %s\n", week, foxNflSchedulePageLink)
	} else {
		foxNflSchedulePageLink = fmt.Sprintf("%s%d", utils.FOX_CFB_REGULER_SEASON_SCHEDULE_URL, week)
		log.Printf("\nRequesting Fox Schedule page for regular season week %d: %s\n", week, foxNflSchedulePageLink)
	}

	page, err := getSchedule(foxNflSchedulePageLink)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return page
}
