package main

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/extract"
	"have-a-nice-pickem-etl/internal/load"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"have-a-nice-pickem-etl/internal/transform/locationdetails"
	"have-a-nice-pickem-etl/internal/transform/record"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
)

// main runs the ETL pipeline: archives existing data, extracts games and teams,
// transforms them into loadable records, and writes CSV outputs.
func main() {
	var weekNum uint = 16
	games := extract.ExtractGames(extract.CfbGamesExtract{Week: weekNum})
	teams := extract.ExtractTeams(extract.CfbTeamsExtract{Week: weekNum})
	locations := extract.ExtractLocations(extract.CfbLocationsExtract{Week: weekNum})

	var allBettingOdds []bettingodds.BettingOdds
	var allBoxscores []boxscore.Boxscore
	var allGames []gamedetails.GameDetails
	var allStats []gamestats.GameStats
	var allTeamRecords []record.Record
	var allTeams []teamdetails.TeamDetails
	var allLocations []locationdetails.LocationDetails

	for i := range games {
		fmt.Printf("\nTransforming Game: %v", games[i])
		gameDetailsRow := gamedetails.New{Game: games[i]}.InstantiateGameDetails()
		espnAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.EspnAwayBettingOdds{Game: games[i]})
		espnHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.EspnHomeBettingOdds{Game: games[i]})
		cbsAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.CbsAwayBettingOdds{Game: games[i]})
		cbsHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.CbsHomeBettingOdds{Game: games[i]})
		foxAwayBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.FoxAwayBettingOdds{Game: games[i]})
		foxHomeBettingOddsRow := bettingodds.InstantiateBettingOdds(bettingodds.FoxHomeBettingOdds{Game: games[i]})
		awayBoxscoreRow := boxscore.InstantiateBoxscore(boxscore.AwayBoxscore{Game: games[i]})
		homeBoxscoreRow := boxscore.InstantiateBoxscore(boxscore.HomeBoxscore{Game: games[i]})
		awayStatsRows := gamestats.InstantiateGameStats(gamestats.AwayTeamStat{Game: games[i]})
		homeStatsRows := gamestats.InstantiateGameStats(gamestats.HomeTeamStat{Game: games[i]})

		allBettingOdds = append(allBettingOdds,
			espnAwayBettingOddsRow, espnHomeBettingOddsRow,
			cbsAwayBettingOddsRow, cbsHomeBettingOddsRow,
			foxAwayBettingOddsRow, foxHomeBettingOddsRow)
		allBoxscores = append(allBoxscores, awayBoxscoreRow, homeBoxscoreRow)
		allGames = append(allGames, gameDetailsRow)
		allStats = append(allStats, awayStatsRows, homeStatsRows)
	}

	for j := range teams {
		fmt.Printf("\nTransforming Team: %v", teams[j])
		teamConferenceRecordRow := record.InstantiateRecord(record.ConferenceRecord{Team: teams[j]})
		teamOverallRecordRow := record.InstantiateRecord(record.OverallRecord{Team: teams[j]})
		teamDetailsRow := teamdetails.New{Team: teams[j]}.Instantiate()

		allTeamRecords = append(allTeamRecords, teamConferenceRecordRow, teamOverallRecordRow)
		allTeams = append(allTeams, teamDetailsRow)
	}

	for k := range locations {
		fmt.Printf("nTransforming Location: %v", locations[k])
		locationRow := locationdetails.InstantiateLocationDetails(locationdetails.New{Location: locations[k]})

		allLocations = append(allLocations, locationRow)
	}

	csvLoadFolderPath := load.InstantiateLoadDirectory()
	load.BettingOdds(allBettingOdds, csvLoadFolderPath)
	load.Boxscores(allBoxscores, csvLoadFolderPath)
	load.GameDetails(allGames, csvLoadFolderPath)
	load.Stats(allStats, csvLoadFolderPath)
	load.TeamDetails(allTeams, csvLoadFolderPath)
	load.TeamRecord(allTeamRecords, csvLoadFolderPath)
	load.LocationDetails(allLocations, csvLoadFolderPath)

}
