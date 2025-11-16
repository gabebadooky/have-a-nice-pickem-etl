/*
Package that:
  - Scrapes the FOX Schedule page HTML for a given week number
*/
package fox

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetSchedule(league string, week uint8) *goquery.Selection {
	var foxSchedulePageLink string = fmt.Sprintf("https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week=%d", week)

	log.Printf("\nRequesting Fox Schedule page for week %d: %s", week, foxSchedulePageLink)
	resp, err := http.Get(foxSchedulePageLink)
	if err != nil {
		log.Panicf("Error occurred navigating to %s:\n%s", foxSchedulePageLink, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panicf("Non 200 response code returned from %s:\n%d", foxSchedulePageLink, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicf("Error occurred instanitating goquery document:\n%s", err)
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	log.Printf("htmlBody:\n%v\n", htmlbody)
	return htmlbody
}
