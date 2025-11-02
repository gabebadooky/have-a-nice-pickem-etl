/*
Package that:
  - Scrapes the CBS Schedule page HTML for a given week number
*/
package cbs

import (
	"fmt"
	"log"
	"net/http"
)

func Schedule(league string, weeknum uint8, year uint16) *http.Response {
	var schedulePageLink string = fmt.Sprintf("https://www.cbssports.com/college-football/scoreboard/FBS/%d/regular/%d/?layout=compact", year, weeknum)

	resp, err := http.Get(schedulePageLink)
	if err != nil {
		log.Printf("Error occurred navigating to %s:\n%s", schedulePageLink, err)
	}
	defer resp.Body.Close()

	return resp
}
