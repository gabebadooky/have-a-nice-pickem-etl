package main

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/gamestats"
)

func main() {
	games := extract.ExtractGames(extract.CfbGamesExtract{Week: 16})
	//var distinctTeams []string

	for i := range 4 {
		// fmt.Printf("Game: %v", games[i])
		//gamedetailsrow := gamedetails.New{Game: games[i]}.InstantiateGameDetails()
		//espnBettingOdds := bettingodds.InstantiateBettingOdds(bettingodds.EspnBettingOdds{Game: games[i]})
		//cbsBettingOdds := bettingodds.InstantiateBettingOdds(bettingodds.CbsBettingOdds{Game: games[i]})
		//awayBoxScore := boxscore.InstantiateBoxscore(boxscore.AwayBoxscore{Game: games[i]})
		//homeBoxScore := boxscore.InstantiateBoxscore(boxscore.HomeBoxscore{Game: games[i]})
		//awayGameStats := gamestats.InstantiateGameStats(gamestats.AwayTeamStat{Game: games[i]})
		homeGomeStats := gamestats.InstantiateGameStats(gamestats.HomeTeamStat{Game: games[i]})

		/*if !slices.Contains(distinctTeams, gamedetailsrow.AwayTeamID) {
				distinctTeams = append(distinctTeams, gamedetailsrow.AwayTeamID)
			}
			if !slices.Contains(distinctTeams, gamedetailsrow.HomeTeamID) {
				distinctTeams = append(distinctTeams, gamedetailsrow.HomeTeamID)
			}*

			fmt.Println()
			fmt.Printf(`
				GameID: %s,
				EspnCode: %s,
				CbsCode: %s,
				FoxCode: %s
			`,
				gamedetailsrow.GameID,
				gamedetailsrow.EspnCode,
				gamedetailsrow.CbsCode,
				gamedetailsrow.FoxCode,
			)
			fmt.Println()
		}

		/*
			teams := extract.ExtractTeams(extract.CfbTeamsExtract{Week: 16})
			fmt.Printf("len(teams): %d", len(teams))

			for i := range len(teams) {
				//teamdetailsrow := teamdetails.New{Team: teams[i]}.Instantiate()
				//conferenceRecord := record.InstantiateRecord(record.ConferenceRecord{Team: teams[i]})
				//overallRecord := record.InstantiateRecord(record.OverallRecord{Team: teams[i]})
				teamStats := teamstats.New{Team: teams[i]}.Instantiate()

				fmt.Println()
				/*fmt.Printf(`
					TeamID: %s
					League: %s
					ESPNCode: %s
					CBSCode: %s
					FoxCode: %s
					VegasCode: %s
					ConferenceID: %s
					Name: %s
					Mascot: %s
					PrimaryColor: %s
					AlternateColor: %s
					Ranking: %d
				`,
					teamdetailsrow.TeamID,
					teamdetailsrow.League,
					teamdetailsrow.ESPNCode,
					teamdetailsrow.CBSCode,
					teamdetailsrow.FoxCode,
					teamdetailsrow.VegasCode,
					teamdetailsrow.ConferenceID,
					teamdetailsrow.Name,
					teamdetailsrow.Mascot,
					teamdetailsrow.PrimaryColor,
					teamdetailsrow.AlternateColor,
					teamdetailsrow.Ranking,
				)*/
		/*fmt.Printf(`
			TeamID: %s
			RecordType: %s
			Wins: %d
			Losses: %d
			Ties: %d
		`,
			conferenceRecord.TeamID,
			conferenceRecord.RecordType,
			conferenceRecord.Wins,
			conferenceRecord.Losses,
			conferenceRecord.Ties,
		)*
		fmt.Printf(`
			TeamID: %s
			RecordType: %s
			Wins: %d
			Losses: %d
			Ties: %d
		`,
			overallRecord.TeamID,
			overallRecord.RecordType,
			overallRecord.Wins,
			overallRecord.Losses,
			overallRecord.Ties,
		)*/

		/*
				for i := range len(teamStats.Stats) {
					fmt.Printf(`
						TeamID: %s
						Metric: %s
						Value: %f
					`,
						teamStats.TeamID,
						teamStats.Stats[i].Metric,
						teamStats.Stats[i].Value,
					)
				}
				fmt.Println()
				fmt.Println()
				* /
			}
		*/
		for i := range len(homeGomeStats.Stats) {
			fmt.Printf(`
				TeamID: %s
				Metric: %s
				Value: %f
			`,
				homeGomeStats.TeamID,
				homeGomeStats.Stats[i].Metric,
				homeGomeStats.Stats[i].Value,
			)
			fmt.Println()
			fmt.Println()
		}
	}
}
