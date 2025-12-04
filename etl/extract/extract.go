package extract

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Schedule interface {
	GetScheduleForWeek()
}

type EndpointResposne interface {
	DecodeJSONResponse()
}

// Scrape Given Schedule Page
func GetSchedulePageBody(schedulePageLink string) (*goquery.Selection, error) {
	doc, err := utils.ScrapePage(schedulePageLink)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	//log.Printf("htmlBody:\n%v\n", htmlbody)
	return htmlbody, nil
}

// Call a given API endpoint and read response
func CallEndpoint(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error occurred calling endpoint: %s: \n%s", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response code returned %s:\n%d", endpoint, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred parsing API Response: \n%s", err)
	}

	return body, nil
}
