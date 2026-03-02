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
	Locations []location.Location
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

// TransformData produces all game-level transformations (details, odds, boxscore, stats) from the extracted game.
func (g NewGameTransformation) TransformData() GameTransformations {
	return GameTransformations{
		GameDetails:         gamedetails.New{Game: g.Game, Locations: g.Locations}.Instantiate(),
		EspnAwayBettingOdds: bettingodds.EspnAwayBettingOdds{Game: g.Game}.Instantiate(),
		EspnHomeBettingOdds: bettingodds.EspnHomeBettingOdds{Game: g.Game}.Instantiate(),
		CbsAwayBettingOdds:  bettingodds.CbsAwayBettingOdds{Game: g.Game}.Instantiate(),
		CbsHomeBettingOdds:  bettingodds.CbsHomeBettingOdds{Game: g.Game}.Instantiate(),
		FoxAwayBettingOdds:  bettingodds.FoxAwayBettingOdds{Game: g.Game}.Instantiate(),
		FoxHomeBettingOdds:  bettingodds.FoxHomeBettingOdds{Game: g.Game}.Instantiate(),
		AwayBoxscore:        boxscore.AwayBoxscore{Game: g.Game}.Instantiate(),
		HomeBoxscore:        boxscore.HomeBoxscore{Game: g.Game}.Instantiate(),
		AwayTeamStats:       gamestats.AwayTeamStat{Game: g.Game}.Instantiate(),
		HomeTeamStats:       gamestats.HomeTeamStat{Game: g.Game}.Instantiate(),
	}
}

// TransformData produces team details and conference/overall records from the extracted team.
func (t NewTeamTransformation) TransformData() TeamTransformations {
	return TeamTransformations{
		TeamDetails:      teamdetails.New{Team: t.Team}.Instantiate(),
		ConferenceRecord: record.ConferenceRecord{Team: t.Team}.Instantiate(),
		OverallRecord:    record.OverallRecord{Team: t.Team}.Instantiate(),
	}
}

// TransformData produces location details from the extracted location data.
func (l NewLocationTransformation) TransformData() LocationTransformations {
	return LocationTransformations{
		Location: locationdetails.New{Location: l.Location}.Instantiate(),
	}
}
