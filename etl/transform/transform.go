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

type GameInstantiator interface {
	transformData() GameTransformations
}

type TeamInstantiator interface {
	transformData() GameTransformations
}

type LocationInstantiator interface {
	transformData() GameTransformations
}

type NewGameTransformation struct {
	GameExtract game.Game
}

type NewTeamTransformation struct {
	TeamExtract team.Team
}

type NewLocationTransformations struct {
	LocationExtract location.Location
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
	gameExtract := g.GameExtract
	newGameDetails := gamedetails.NewGameDetails{GameExtract: gameExtract}
	newEspnBettingOdds := bettingodds.EspnBettingOdds{GameExtract: gameExtract}
	newCbsBettingOdds := bettingodds.CbsBettingOdds{GameExtract: gameExtract}
	//newFoxBettingOdds := bettingodds.FoxBettingOdds{GameExtract: gameExtract}
	//newVegasBettingOdds := bettingodds.VegasBettingOdds{GameExtract: gameExtract}
	newAwayBoxscore := boxscore.AwayBoxscore{GameExtract: gameExtract}
	newHomeBoxscore := boxscore.HomeBoxscore{GameExtract: gameExtract}

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
	teamExtract := t.TeamExtract
	newTeamDetails := teamdetails.NewTeamDetails{TeamExtract: teamExtract}
	newConferenceRecord := record.ConferenceRecord{TeamExtract: teamExtract}
	newOverallRecord := record.OverallRecord{TeamExtract: teamExtract}

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
	locationExtract := l.LocationExtract
	newLocationDetails := locationdetails.NewLocation{LocationExtract: locationExtract}
	var locationsTransformation locationdetails.LocationDetails = locationdetails.InstantiateLocationDetails(newLocationDetails)

	return LocationTransformations{
		Location: locationsTransformation,
	}
}
