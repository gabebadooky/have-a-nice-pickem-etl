/*
Package that:
  - Scrapes the CBS Schedule page HTML for a given week number
*/
package cbs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetSchedule(league string, week uint8, year uint16) *goquery.Selection {
	var cbsSchedulePageLink string = fmt.Sprintf("https://www.cbssports.com/college-football/scoreboard/FBS/%d/regular/%d/?layout=compact", year, week)

	log.Printf("\nRequesting CBS Schedule page for week %d: %s", week, cbsSchedulePageLink)
	resp, err := http.Get(cbsSchedulePageLink)
	if err != nil {
		log.Panicf("Error occurred navigating to %s:\n%s", cbsSchedulePageLink, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panicf("Non 200 response code returned from %s:\n%d", cbsSchedulePageLink, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicf("Error occurred instantiating goquery document:\n%s", err)
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	log.Printf("htmlBody:\n%v\n", htmlbody)
	return htmlbody
}
