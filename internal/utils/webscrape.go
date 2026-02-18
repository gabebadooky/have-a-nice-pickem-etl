// Package utils provides web scraping utilities for fetching and parsing HTML content.
// These functions handle HTTP requests and return goquery document selections
// for parsing web pages in the ETL pipeline.
package utils

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// scrapePage fetches the given URL and returns the response body as a goquery document.
func scrapePage(pageURL string) (*goquery.Document, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("error occurred navigating to %s:\n%s", pageURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response code returned from %s:\n%d", pageURL, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred instanitating goquery document:\n%s", err)
	}

	return doc, nil
}

// GetGoQuerySelectionBody fetches the page at the given URL and returns its body as a goquery selection.
func GetGoQuerySelectionBody(pageLink string) (*goquery.Selection, error) {
	doc, err := scrapePage(pageLink)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	var htmlbody *goquery.Selection = doc.Find("body").First()
	//log.Printf("htmlBody:\n%v\n", htmlbody)
	return htmlbody, nil
}
