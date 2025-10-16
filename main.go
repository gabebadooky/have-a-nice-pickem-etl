package main

import (
	"fmt"

	extractESPN "github.com/gabebadooky/have-a-nice-pickem-etl/etl/extract/espn"
)

type details struct {
	gameID     string
	league     string
	week       uint8
	year       uint16
	espnCode   string
	cbsCode    string
	foxCode    string
	vegasCode  string
	awayTeamID string
	homeTeamID string
	date       string
	time       string
	tvCoverage string
	finished   string
}

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
	fmt.Println("Hello, World")
	extractESPN.GameSummary("401754528")
	extractESPN.TeamSummary("158")
}
