package main

import (
	"fmt"
	extractCBS "have-a-nice-pickem-etl/etl/extract/cbs"
)

type ESPNGameDetailsResponse struct {
	Header Header `json:"header"`
}

type Header struct {
	Week         int8           `json:"week"`
	Season       Season         `json:"season"`
	ESPNGameCode string         `json:"id"`
	Competitions []Competitions `[]json:"competitions"`
}

type Season struct {
	Year uint16 `json:"year"`
}

type Competitions struct {
	Competitors []Competitors `[]json:"competitors"`
	Date        string        `json:"date"`
	Broadcasts  []Media       `[]json:"broadcasts"`
	Status      Status        `json:"status"`
}

type Competitors struct {
	HomeAway string `json:"homeAway"`
	TeamID   string `json:"slug"`
}

type Broadcasts struct {
	Media Media `json:"media"`
}

type Media struct {
	ShortName string `json:"shortName"`
}

type Status struct {
	Type Type `json:"type"`
}

type Type struct {
	Completed string `json:"completed"`
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
	//fmt.Println("Hello, World")
	//espnGameSummary := extract.Game("401754528")
	//fmt.Println(transform.Game("CFB", espnGameSummary))
	fmt.Println(extractCBS.Week("CFB", 11, 2025))
	//espnGameSummary := extractESPN.GameSummary("401754528")
	//extractESPN.TeamSummary("158")
	//fmt.Println(gameDetails.GameDetails(espnGameSummary))
}
