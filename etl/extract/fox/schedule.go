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

func Schedule(league string, weeknum uint8) *goquery.Selection {
	var schedulePageLink string = fmt.Sprintf("https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week=%d", weeknum)

	resp, err := http.Get(schedulePageLink)
	if err != nil {
		log.Panicf("Error occurred navigating to %s:\n%s", schedulePageLink, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panicf("Non 200 response code returned from %s:\n%d", schedulePageLink, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicf("Error occurred instanitating goquery document:\n%s", err)
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	return htmlbody
}
