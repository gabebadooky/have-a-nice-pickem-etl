package pickemstructs

import (
	"github.com/PuerkitoBio/goquery"
)

type ConsolidatedGameProperties struct {
	EspnDetails ESPNGameDetailsResponse
	CbsPage     *goquery.Selection
	FoxPage     *goquery.Selection
}
