package main

import (
	"have-a-nice-pickem-etl/etl/extract"
)

func main() {
	var week uint8 = 7
	var league string = "CFB"
	var consolidatedSchedule extract.ConsolidatedSchedule

	if league == "CFB" {
		consolidatedSchedule = extract.CFBSchedule{Week: week}.ExtractSchedule()
	} else {
		consolidatedSchedule = extract.NFLSchedule{Week: week}.ExtractSchedule()
	}

	for i := 0; i < 2; i++ {
		var espnGameCode string
		espnGameCode = consolidatedSchedule.ESPN.Events[i].ID

		if league == "CFB" {
			game := extract.CFBGame{EspnCode: espnGameCode}
		} else {
			game := extract.NFLGame{EspnCode: espnGameCode}
		}
	}

	/*
		for i := 0; i < 2; /*len(espnSchedule.Events)* / i++ {
			var espnGameCode string = espnSchedule.Events[i].ID
			var espnGameDetails pickemstructs.ESPNGameDetailsResponse = espn.GetGame(espnGameCode)
			consolidatedGameProperties := pickemstructs.ConsolidatedGameProperties{
				EspnDetails: espnGameDetails,
				CbsPage:     cbsSchedule,
				FoxPage:     foxSchedule,
			}
			gameProperties := GameProperties{
				GameDetails:     transform.CreateGameDetailsRecord(consolidatedGameProperties),
				EspnBettingOdds: transform.CreateESPNBettingOddsRecord(consolidatedGameProperties),
				CbsBettingOdds:  transform.CreateCBSBettingOddsRecord(consolidatedGameProperties),
				AwayBoxScore:    transform.CreateBoxScoreRecord(consolidatedGameProperties, "AWAY"),
				HomeBoxScore:    transform.CreateBoxScoreRecord(consolidatedGameProperties, "HOME"),
				Location:        transform.CreateLocationRecord(consolidatedGameProperties),
			}

			fmt.Printf("\ngameDetails: %v\n", gameProperties.GameDetails)
			fmt.Printf("\nespnBettingOdds: %v\n", gameProperties.EspnBettingOdds)
			fmt.Printf("\ncbsBettingOdds: %v\n", gameProperties.CbsBettingOdds)
			fmt.Printf("\nawayBoxScore: %v\n", gameProperties.AwayBoxScore)
			fmt.Printf("\nhomeBoxScore: %v\n", gameProperties.HomeBoxScore)
			fmt.Printf("\nlocation: %v\n", gameProperties.Location)

			if !slices.Contains(distinct.Teams, gameProperties.GameDetails.AwayTeamID) {
				distinct.Teams = append(distinct.Teams, gameProperties.GameDetails.AwayTeamID)
			}
			if !slices.Contains(distinct.Teams, gameProperties.GameDetails.HomeTeamID) {
				distinct.Teams = append(distinct.Teams, gameProperties.GameDetails.HomeTeamID)
			}
			if !slices.Contains(distinct.GameDetails, gameProperties.GameDetails) {
				distinct.GameDetails = append(distinct.GameDetails, gameProperties.GameDetails)
			}
			if !slices.Contains(distinct.BettingOdds, gameProperties.EspnBettingOdds) {
				distinct.BettingOdds = append(distinct.BettingOdds, gameProperties.EspnBettingOdds)
			}
			if !slices.Contains(distinct.BettingOdds, gameProperties.CbsBettingOdds) {
				distinct.BettingOdds = append(distinct.BettingOdds, gameProperties.CbsBettingOdds)
			}
			if !slices.Contains(distinct.BoxScores, gameProperties.AwayBoxScore) {
				distinct.BoxScores = append(distinct.BoxScores, gameProperties.AwayBoxScore)
			}
			if !slices.Contains(distinct.BoxScores, gameProperties.HomeBoxScore) {
				distinct.BoxScores = append(distinct.BoxScores, gameProperties.HomeBoxScore)
			}
			if !slices.Contains(distinct.Locations, gameProperties.Location) {
				distinct.Locations = append(distinct.Locations, gameProperties.Location)
			}

			fmt.Printf("\ndistinctTeams: %v\n", distinct.Teams)
			fmt.Printf("\ndistinctGameDetails: %v\n", distinct.GameDetails)
			fmt.Printf("\ndistinctBettingOdds: %v\n", distinct.BettingOdds)
			fmt.Printf("\ndistinctBoxScoress: %v\n", distinct.BoxScores)
			fmt.Printf("\ndistinctLocations: %v\n", distinct.Locations)
		}
	*/
}
