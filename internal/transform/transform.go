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

type GameTransformer interface {
	transformData() GameTransformations
}

type TeamTransformer interface {
	transformData() TeamTransformations
}

type LocationTransformer interface {
	transformData() LocationTransformations
}

type NewGameTransformation struct {
	game.Game
}

type NewTeamTransformation struct {
	team.Team
}

type NewLocationTransformations struct {
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
	//VegasBettingOdds bettingodds.BettingOdds
	//VegasBettingOdds bettingodds.BettingOdds
	AwayBoxscore  boxscore.Boxscore
	HomeBoxscore  boxscore.Boxscore
	AwayTeamStats gamestats.GameStats
	HomeTeamStats gamestats.GameStats
}

type TeamTransformations struct {
	TeamDetails      teamdetails.TeamDetails
	ConferenceRecord record.Record
	OverallRecord    record.Record
}

type LocationTransformations struct {
	Location locationdetails.LocationDetails
}

func PerformGameTransformations(g NewGameTransformation) GameTransformations {
	return g.transformData()
}

func PerformTeamTransformations(t NewTeamTransformation) TeamTransformations {
	return t.transformData()
}

func PerformLocationTransformations(l NewLocationTransformations) LocationTransformations {
	return l.transformData()
}

func (g NewGameTransformation) transformData() GameTransformations {
	newGameDetails := gamedetails.New{Game: g.Game}
	newEspnAwayBettingOdds := bettingodds.EspnAwayBettingOdds{Game: g.Game}
	newEspnHomeBettingOdds := bettingodds.EspnHomeBettingOdds{Game: g.Game}
	newCbsAwayBettingOdds := bettingodds.CbsAwayBettingOdds{Game: g.Game}
	newCbsHomeBettingOdds := bettingodds.CbsHomeBettingOdds{Game: g.Game}
	newFoxAwayBettingOdds := bettingodds.FoxAwayBettingOdds{Game: g.Game}
	newFoxHomeBettingOdds := bettingodds.FoxHomeBettingOdds{Game: g.Game}
	//newVegasBettingOdds := bettingodds.VegasBettingOdds{GameExtract: gameExtract}
	//newVegasBettingOdds := bettingodds.VegasBettingOdds{GameExtract: gameExtract}
	newAwayBoxscore := boxscore.AwayBoxscore{Game: g.Game}
	newHomeBoxscore := boxscore.HomeBoxscore{Game: g.Game}

	//var gameDetailsTransformation gamedetails.GameDetails = gamedetails.InstantiateGameDetails(newGameDetails)
	var gameDetailsTransformation gamedetails.GameDetails = newGameDetails.InstantiateGameDetails()
	var espnAwayBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newEspnAwayBettingOdds)
	var espnHomeBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newEspnHomeBettingOdds)
	var cbsAwayBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newCbsAwayBettingOdds)
	var cbsHomeBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newCbsHomeBettingOdds)
	var foxAwayBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newFoxAwayBettingOdds)
	var foxHomeBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newFoxHomeBettingOdds)
	var awayBoxscoreTransformation boxscore.Boxscore = boxscore.InstantiateBoxscore(newAwayBoxscore)
	var homeBoxscoreTransformation boxscore.Boxscore = boxscore.InstantiateBoxscore(newHomeBoxscore)

	return GameTransformations{
		GameDetails:         gameDetailsTransformation,
		EspnAwayBettingOdds: espnAwayBettingOddsTransformation,
		EspnHomeBettingOdds: espnHomeBettingOddsTransformation,
		CbsAwayBettingOdds:  cbsAwayBettingOddsTransformation,
		CbsHomeBettingOdds:  cbsHomeBettingOddsTransformation,
		FoxAwayBettingOdds:  foxAwayBettingOddsTransformation,
		FoxHomeBettingOdds:  foxHomeBettingOddsTransformation,
		AwayBoxscore:        awayBoxscoreTransformation,
		HomeBoxscore:        homeBoxscoreTransformation,
	}
}

func (t NewTeamTransformation) transformData() TeamTransformations {
	newTeamDetails := teamdetails.New{Team: t.Team}
	newConferenceRecord := record.ConferenceRecord{Team: t.Team}
	newOverallRecord := record.OverallRecord{Team: t.Team}

	//var teamDetailsTransformation teamdetails.TeamDetails = teamdetails.InstantiateTeamDetails(newTeamDetails)
	var teamDetailsTransformation teamdetails.TeamDetails = newTeamDetails.Instantiate()
	var conferenceRecordTransformation record.Record = record.InstantiateRecord(newConferenceRecord)
	var overallRecordTransformation record.Record = record.InstantiateRecord(newOverallRecord)

	return TeamTransformations{
		TeamDetails:      teamDetailsTransformation,
		ConferenceRecord: conferenceRecordTransformation,
		OverallRecord:    overallRecordTransformation,
	}
}

func (l NewLocationTransformations) transformData() LocationTransformations {
	newLocationDetails := locationdetails.New{Location: l.Location}
	var locationsTransformation locationdetails.LocationDetails = locationdetails.InstantiateLocationDetails(newLocationDetails)

	return LocationTransformations{
		Location: locationsTransformation,
	}
}
