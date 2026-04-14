// Package transform provides the main transformation orchestration layer for the ETL pipeline.
// It coordinates the transformation of extracted game, team, and location data into
// structured output formats suitable for loading into the target data store.
package transform

import (
	"have-a-nice-pickem-etl/internal/extract"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"have-a-nice-pickem-etl/internal/transform/locationdetails"
	"have-a-nice-pickem-etl/internal/transform/record"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
	"log"
	"slices"
)

type New struct {
	extract.Extract
}

type GameTransformations struct {
	AllBettingOdds []bettingodds.BettingOdds
	AllBoxscores   []boxscore.Boxscore
	AllGameDetails []gamedetails.GameDetails
	AllGameStats   []gamestats.GameStats
}

type TeamTransformations struct {
	AllTeams       []teamdetails.TeamDetails
	AllTeamRecords []record.Record
}

type LocationTransformations struct {
	AllLocations []locationdetails.LocationDetails
}

type Transformation struct {
	GameTransformations     GameTransformations
	TeamTransformations     TeamTransformations
	LocationTransformations LocationTransformations
}

func keyExistsInSlice[T any](
	slice []T,
	target T,
	equals func(T, T) bool,
) bool {
	idx := slices.IndexFunc(slice, func(item T) bool {
		return equals(item, target)
	})
	return idx != -1
}

// TransformData produces all game-level transformations (details, odds, boxscore, stats) from the extracted game.
func (t New) transformGameData() GameTransformations {
	var allBettingOdds []bettingodds.BettingOdds
	var allBoxscores []boxscore.Boxscore
	var allGameDetails []gamedetails.GameDetails
	var allGameStats []gamestats.GameStats

	for i := range t.GamesExtract {
		game := t.GamesExtract[i]
		log.Printf("\nTransforming Game: %v", game)

		gameDetailsRow := gamedetails.New{Game: game, Locations: t.LocationsExtract}.Instantiate()
		espnAwayBettingOddsRow := bettingodds.EspnAwayBettingOdds{Game: game}.Instantiate()
		espnHomeBettingOddsRow := bettingodds.EspnHomeBettingOdds{Game: game}.Instantiate()
		cbsAwayBettingOddsRow := bettingodds.CbsAwayBettingOdds{Game: game}.Instantiate()
		cbsHomeBettingOddsRow := bettingodds.CbsHomeBettingOdds{Game: game}.Instantiate()
		foxAwayBettingOddsRow := bettingodds.FoxAwayBettingOdds{Game: game}.Instantiate()
		foxHomeBettingOddsRow := bettingodds.FoxHomeBettingOdds{Game: game}.Instantiate()
		awayBoxScoreRow := boxscore.AwayBoxscore{Game: game}.Instantiate()
		homeBoxScoreRow := boxscore.HomeBoxscore{Game: game}.Instantiate()
		awayTeamGameStats := gamestats.AwayTeamStat{Game: game}.Instantiate()
		homeTeamGameStats := gamestats.HomeTeamStat{Game: game}.Instantiate()

		allGameDetails = append(allGameDetails, gameDetailsRow)
		allBoxscores = append(allBoxscores, awayBoxScoreRow, homeBoxScoreRow)
		if len(awayTeamGameStats.Stats) > 0 {
			allGameStats = append(allGameStats, awayTeamGameStats)
		}
		if len(homeTeamGameStats.Stats) > 0 {
			allGameStats = append(allGameStats, homeTeamGameStats)
		}
		allBettingOdds = append(
			allBettingOdds,
			espnAwayBettingOddsRow,
			espnHomeBettingOddsRow,
			cbsAwayBettingOddsRow,
			cbsHomeBettingOddsRow,
			foxAwayBettingOddsRow,
			foxHomeBettingOddsRow,
		)
	}

	return GameTransformations{
		AllBettingOdds: allBettingOdds,
		AllBoxscores:   allBoxscores,
		AllGameDetails: allGameDetails,
		AllGameStats:   allGameStats,
	}
}

// TransformData produces team details and conference/overall records from the extracted team.
func (t New) transformTeamData() TeamTransformations {
	var allTeamRecords []record.Record
	var allTeams []teamdetails.TeamDetails

	for i := range t.TeamsExtract {
		team := t.TeamsExtract[i]
		log.Printf("\nTransforming Team: %v", team)

		teamDetailsRow := teamdetails.New{Team: team}.Instantiate()
		teamConferenceRecordRow := record.ConferenceRecord{Team: team}.Instantiate()
		teamOverallRecordRow := record.OverallRecord{Team: team}.Instantiate()

		if !keyExistsInSlice(allTeams, teamDetailsRow, func(a, b teamdetails.TeamDetails) bool {
			return a.TeamID == b.TeamID
		}) {
			allTeams = append(allTeams, teamDetailsRow)
		}

		if !keyExistsInSlice(allTeamRecords, teamOverallRecordRow, func(a, b record.Record) bool {
			return a.TeamID == b.TeamID && a.RecordType == b.RecordType
		}) {
			allTeamRecords = append(allTeamRecords, teamOverallRecordRow)
		}

		if !keyExistsInSlice(allTeamRecords, teamConferenceRecordRow, func(a, b record.Record) bool {
			return a.TeamID == b.TeamID && a.RecordType == b.RecordType
		}) {
			allTeamRecords = append(allTeamRecords, teamConferenceRecordRow)
		}
	}

	return TeamTransformations{
		AllTeams:       allTeams,
		AllTeamRecords: allTeamRecords,
	}
}

// TransformData produces location details from the extracted location data.
func (t New) transformLocationData() LocationTransformations {
	var allLocations []locationdetails.LocationDetails

	for i := range t.LocationsExtract {
		location := t.LocationsExtract[i]
		log.Printf("\nTransforming Location: %v", location)

		locationDetailsRow := locationdetails.New{Location: location}.Instantiate()

		if !keyExistsInSlice(allLocations, locationDetailsRow, func(a, b locationdetails.LocationDetails) bool {
			return a.LocationID == b.LocationID
		}) {
			allLocations = append(allLocations, locationDetailsRow)
		}
	}

	return LocationTransformations{
		AllLocations: allLocations,
	}
}

func (t New) PerformTransformation() Transformation {
	return Transformation{
		GameTransformations:     t.transformGameData(),
		TeamTransformations:     t.transformTeamData(),
		LocationTransformations: t.transformLocationData(),
	}
}
