package transform

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/extract/location"
	"have-a-nice-pickem-etl/etl/extract/team"
	"have-a-nice-pickem-etl/etl/transform/bettingodds"
	"have-a-nice-pickem-etl/etl/transform/boxscore"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"
	"have-a-nice-pickem-etl/etl/transform/locationdetails"
	"have-a-nice-pickem-etl/etl/transform/record"
	"have-a-nice-pickem-etl/etl/transform/teamdetails"
)

type GameTransformer interface {
	transformData() GameTransformations
}

type TeamTransformer interface {
	transformData() GameTransformations
}

type LocationTransformer interface {
	transformData() GameTransformations
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
	GameDetails     gamedetails.GameDetails
	EspnBettingOdds bettingodds.BettingOdds
	CbsBettingOdds  bettingodds.BettingOdds
	//FoxBettingOdds   bettingodds.BettingOdds
	//VegasBettingOdds bettingodds.BettingOdds
	AwayBoxscore boxscore.Boxscore
	HomeBoxscore boxscore.Boxscore
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
	newEspnBettingOdds := bettingodds.EspnBettingOdds{Game: g.Game}
	newCbsBettingOdds := bettingodds.CbsBettingOdds{Game: g.Game}
	//newFoxBettingOdds := bettingodds.FoxBettingOdds{GameExtract: gameExtract}
	//newVegasBettingOdds := bettingodds.VegasBettingOdds{GameExtract: gameExtract}
	newAwayBoxscore := boxscore.AwayBoxscore{Game: g.Game}
	newHomeBoxscore := boxscore.HomeBoxscore{Game: g.Game}

	var gameDetailsTransformation gamedetails.GameDetails = gamedetails.InstantiateGameDetails(newGameDetails)
	var espnBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newEspnBettingOdds)
	var cbsBettingOddsTransformation bettingodds.BettingOdds = bettingodds.InstantiateBettingOdds(newCbsBettingOdds)
	var awayBoxscoreTransformation boxscore.Boxscore = boxscore.InstantiateBoxscore(newAwayBoxscore)
	var homeBoxscoreTransformation boxscore.Boxscore = boxscore.InstantiateBoxscore(newHomeBoxscore)

	return GameTransformations{
		GameDetails:     gameDetailsTransformation,
		EspnBettingOdds: espnBettingOddsTransformation,
		CbsBettingOdds:  cbsBettingOddsTransformation,
		AwayBoxscore:    awayBoxscoreTransformation,
		HomeBoxscore:    homeBoxscoreTransformation,
	}
}

func (t NewTeamTransformation) transformData() TeamTransformations {
	newTeamDetails := teamdetails.New{Team: t.Team}
	newConferenceRecord := record.ConferenceRecord{Team: t.Team}
	newOverallRecord := record.OverallRecord{Team: t.Team}

	var teamDetailsTransformation teamdetails.TeamDetails = teamdetails.InstantiateTeamDetails(newTeamDetails)
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
