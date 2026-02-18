package main

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/extract"
	"have-a-nice-pickem-etl/internal/load"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"have-a-nice-pickem-etl/internal/transform/record"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
	"have-a-nice-pickem-etl/internal/utils"
)

// main runs the ETL pipeline: archives existing data, extracts games and teams,
// transforms them into loadable records, and writes CSV outputs.
func main() {
	utils.ArchiveData()

	games := extract.ExtractGames(extract.CfbGamesExtract{Week: 16})
	teams := extract.ExtractTeams(extract.CfbTeamsExtract{Week: 16})

	var allBettingOdds []bettingodds.BettingOdds
	var allBoxscores []boxscore.Boxscore
	var allGames []gamedetails.GameDetails
	var allStats []gamestats.GameStats
	var allTeamRecords []record.Record
	var allTeams []teamdetails.TeamDetails

	for i := range 4 {
		fmt.Printf("Game: %v", games[i])
		gameDetailsRow := gamedetails.New{Game: games[i]}.InstantiateGameDetails()
		espnAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.EspnAwayBettingOdds{Game: games[i]})
		espnHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.EspnHomeBettingOdds{Game: games[i]})
		cbsAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.CbsAwayBettingOdds{Game: games[i]})
		cbsHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.CbsHomeBettingOdds{Game: games[i]})
		foxAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.FoxAwayBettingOdds{Game: games[i]})
		foxHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.FoxAwayBettingOdds{})
		awayBoxscoreRow := boxscore.InstantiateBoxscore(boxscore.AwayBoxscore{Game: games[i]})
		homeBoxscoreRow := boxscore.InstantiateBoxscore(boxscore.HomeBoxscore{Game: games[i]})
		awayStatsRows := gamestats.InstantiateGameStats(gamestats.AwayTeamStat{Game: games[i]})
		homeStatsRows := gamestats.InstantiateGameStats(gamestats.HomeTeamStat{Game: games[i]})

		teamConferenceRecordRow := record.InstantiateRecord(record.ConferenceRecord{Team: teams[i]})
		teamOverallRecordRow := record.InstantiateRecord(record.OverallRecord{Team: teams[i]})
		teamDetailsRow := teamdetails.New{Team: teams[i]}.Instantiate()

		//LocationRow := locationdetails.InstantiateLocationDetails()

		allBettingOdds = append(allBettingOdds,
			espnAwayBettingOddsRow, espnHomeBettingOddsRow,
			cbsAwayBettingOddsRow, cbsHomeBettingOddsRow,
			foxAwayBettingOddsRow, foxHomeBettingOddsRow)
		allBoxscores = append(allBoxscores, awayBoxscoreRow, homeBoxscoreRow)
		allGames = append(allGames, gameDetailsRow)
		allStats = append(allStats, awayStatsRows, homeStatsRows)
		allTeamRecords = append(allTeamRecords, teamConferenceRecordRow, teamOverallRecordRow)
		allTeams = append(allTeams, teamDetailsRow)

		/*
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
			* /
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
		*/
	}

	load.BettingOdds(allBettingOdds)
	load.Boxscores(allBoxscores)
	load.GameDetails(allGames)
	load.Stats(allStats)
	load.TeamDetails(allTeams)
	load.TeamRecord(allTeamRecords)

}
