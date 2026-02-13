// Package transform provides the main transformation orchestration layer for the ETL pipeline.
// It coordinates the transformation of extracted game, team, and location data into
// structured output formats suitable for loading into the target data store.
package transform

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/extract/location"
	"have-a-nice-pickem-etl/internal/extract/team"
	"have-a-nice-pickem-etl/internal/transform/bettingodds"
	"have-a-nice-pickem-etl/internal/transform/boxscore"
	"have-a-nice-pickem-etl/internal/transform/gamedetails"
	"have-a-nice-pickem-etl/internal/transform/gamestats"
	"have-a-nice-pickem-etl/internal/transform/locationdetails"
	"have-a-nice-pickem-etl/internal/transform/record"
	"have-a-nice-pickem-etl/internal/transform/teamdetails"
)

type NewGameTransformation struct {
	game.Game
}

type NewTeamTransformation struct {
	team.Team
}

type NewLocationTransformation struct {
	location.Location
}

type GameTransformations struct {
	GameDetails         gamedetails.GameDetails
	EspnAwayBettingOdds bettingodds.BettingOdds
	EspnHomeBettingOdds bettingodds.BettingOdds
	CbsAwayBettingOdds  bettingodds.BettingOdds
	CbsHomeBettingOdds  bettingodds.BettingOdds
	FoxAwayBettingOdds  bettingodds.BettingOdds
	FoxHomeBettingOdds  bettingodds.BettingOdds
	AwayBoxscore        boxscore.Boxscore
	HomeBoxscore        boxscore.Boxscore
	AwayTeamStats       gamestats.GameStats
	HomeTeamStats       gamestats.GameStats
}

type TeamTransformations struct {
	TeamDetails      teamdetails.TeamDetails
	ConferenceRecord record.Record
	OverallRecord    record.Record
}

type LocationTransformations struct {
	Location locationdetails.LocationDetails
}

func (g NewGameTransformation) TransformData() GameTransformations {
	// Initialize transformation structs
	newGameDetails := gamedetails.New{Game: g.Game}
	newEspnAwayBettingOdds := bettingodds.EspnAwayBettingOdds{Game: g.Game}
	newEspnHomeBettingOdds := bettingodds.EspnHomeBettingOdds{Game: g.Game}
	newCbsAwayBettingOdds := bettingodds.CbsAwayBettingOdds{Game: g.Game}
	newCbsHomeBettingOdds := bettingodds.CbsHomeBettingOdds{Game: g.Game}
	newFoxAwayBettingOdds := bettingodds.FoxAwayBettingOdds{Game: g.Game}
	newFoxHomeBettingOdds := bettingodds.FoxHomeBettingOdds{Game: g.Game}
	newAwayBoxscore := boxscore.AwayBoxscore{Game: g.Game}
	newHomeBoxscore := boxscore.HomeBoxscore{Game: g.Game}
	newAwayTeamStats := gamestats.AwayTeamStat{Game: g.Game}
	newHomeTeamStats := gamestats.HomeTeamStat{Game: g.Game}

	// Perform transformations
	gameDetails := newGameDetails.InstantiateGameDetails()
	espnAwayBettingOdds := bettingodds.InstantiateBettingOdds(newEspnAwayBettingOdds)
	espnHomeBettingOdds := bettingodds.InstantiateBettingOdds(newEspnHomeBettingOdds)
	cbsAwayBettingOdds := bettingodds.InstantiateBettingOdds(newCbsAwayBettingOdds)
	cbsHomeBettingOdds := bettingodds.InstantiateBettingOdds(newCbsHomeBettingOdds)
	foxAwayBettingOdds := bettingodds.InstantiateBettingOdds(newFoxAwayBettingOdds)
	foxHomeBettingOdds := bettingodds.InstantiateBettingOdds(newFoxHomeBettingOdds)
	awayBoxscore := boxscore.InstantiateBoxscore(newAwayBoxscore)
	homeBoxscore := boxscore.InstantiateBoxscore(newHomeBoxscore)
	awayTeamStats := gamestats.InstantiateGameStats(newAwayTeamStats)
	homeTeamStats := gamestats.InstantiateGameStats(newHomeTeamStats)

	return GameTransformations{
		GameDetails:         gameDetails,
		EspnAwayBettingOdds: espnAwayBettingOdds,
		EspnHomeBettingOdds: espnHomeBettingOdds,
		CbsAwayBettingOdds:  cbsAwayBettingOdds,
		CbsHomeBettingOdds:  cbsHomeBettingOdds,
		FoxAwayBettingOdds:  foxAwayBettingOdds,
		FoxHomeBettingOdds:  foxHomeBettingOdds,
		AwayBoxscore:        awayBoxscore,
		HomeBoxscore:        homeBoxscore,
		AwayTeamStats:       awayTeamStats,
		HomeTeamStats:       homeTeamStats,
	}
}

func (t NewTeamTransformation) TransformData() TeamTransformations {
	// Initialize transformation structs
	newTeamDetails := teamdetails.New{Team: t.Team}
	newConferenceRecord := record.ConferenceRecord{Team: t.Team}
	newOverallRecord := record.OverallRecord{Team: t.Team}

	// Perform transformations
	teamDetails := newTeamDetails.Instantiate()
	conferenceRecord := record.InstantiateRecord(newConferenceRecord)
	overallRecord := record.InstantiateRecord(newOverallRecord)

	return TeamTransformations{
		TeamDetails:      teamDetails,
		ConferenceRecord: conferenceRecord,
		OverallRecord:    overallRecord,
	}
}

func (l NewLocationTransformation) TransformData() LocationTransformations {
	// Initialize transformation struct
	newLocationDetails := locationdetails.New{Location: l.Location}

	// Perform transformation
	locationDetails := locationdetails.InstantiateLocationDetails(newLocationDetails)

	return LocationTransformations{
		Location: locationDetails,
	}
}
