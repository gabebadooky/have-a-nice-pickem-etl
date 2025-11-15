package main

import (
	"fmt"
	extractcbs "have-a-nice-pickem-etl/etl/extract/cbs"
	extractespn "have-a-nice-pickem-etl/etl/extract/espn"
	transform "have-a-nice-pickem-etl/etl/transform"
	etltypes "have-a-nice-pickem-etl/etl/types"

	"github.com/PuerkitoBio/goquery"
)

type boxScore struct {
	awayQ1score       uint8
	awayQ2score       uint8
	awayQ3score       uint8
	awayQ4score       uint8
	awayOvertimeScore uint8
	awayTotalscore    uint8
	homeQ1score       uint8
	homeQ2score       uint8
	homeQ3score       uint8
	homeQ4score       uint8
	homeOvertimeScore uint8
	homeTotalscore    uint8
}

type location struct {
	stadium   string
	city      string
	state     string
	latitude  float32
	longitude float32
}

type odds struct {
	awayMoneyline string
	homeMoneyline string
	awaySpread    string
	homeSpread    string
	overUnder     string
}

func main() {
	var espnGameSummary etltypes.ESPNGameDetailsResponse = extractespn.Game("401754528")
	var cbsSchedulePage *goquery.Selection = extractcbs.Schedule("CFB", 11, 2025)
	fmt.Println(transform.Game(espnGameSummary, cbsSchedulePage))
}
