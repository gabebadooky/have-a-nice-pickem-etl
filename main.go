package main

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract/cbs"
	"have-a-nice-pickem-etl/etl/extract/espn"
	"have-a-nice-pickem-etl/etl/extract/fox"
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var espnSchedule pickemstructs.ESPNScheduleResponse = espn.GetSchedule("CFB", 14)
	var cbsSchedule *goquery.Selection = cbs.GetSchedule("CFB", 14, 2025)
	var foxSchedule *goquery.Selection = fox.GetSchedule("CFB", 14)

	for i := 0; i < 2; /*len(espnSchedule.Events)*/ i++ {
		var espnGameCode string = espnSchedule.Events[i].ID
		var espnGameDetails pickemstructs.ESPNGameDetailsResponse = espn.GetGame(espnGameCode)

		var gameDetails pickemstructs.GameDetails = transform.CreateGameDetailsRecord(espnGameDetails, cbsSchedule, foxSchedule)
		var espnBettingOdds pickemstructs.BettingOdds = transform.CreateBettingOddsRecord(espnGameDetails, "ESPN")
		var awayBoxScore pickemstructs.Boxscore = transform.CreateBoxScoreRecord(espnGameDetails, "AWAY")
		var homeBoxScore pickemstructs.Boxscore = transform.CreateBoxScoreRecord(espnGameDetails, "HOME")

		fmt.Printf("\ngameDetails: %v\n", gameDetails)
		fmt.Printf("\nespnBettingOdds: %v\n", espnBettingOdds)
		fmt.Printf("\nawayBoxScore: %v\n", awayBoxScore)
		fmt.Printf("\nhomeBoxScore: %v\n", homeBoxScore)
	}
}
