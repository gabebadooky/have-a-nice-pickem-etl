package extract

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/extract/location"
	"have-a-nice-pickem-etl/etl/extract/schedule"
	espnsched "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/extract/team"
)

type GamesExtract interface {
	ExtractGames() []game.Game
}

type TeamsExtract interface {
	ExtractTeams() []team.Team
}

type CfbGamesExtract struct {
	Week uint8
}

type NflGamesExtract struct {
	Week uint8
}

type CfbTeamsExtract struct {
	Week uint8
}

type NflTeamsExtract struct {
	Week uint8
}

type LocationsExtract struct {
	Week uint8
}

func (g CfbGamesExtract) ExtractGames() []game.Game {
	weekSchedule := schedule.AllCfbScheduleInfo{Week: g.Week}.ConsolidateScheduleInfo()
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game := game.AllCfbGameInfo{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		}.ConsolidateGameInfo()

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

func (g NflGamesExtract) ExtractGames() []game.Game {
	weekSchedule := schedule.AllNflScheduleInfo{Week: g.Week}.ConsolidateScheduleInfo()
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game := game.AllNflGameInfo{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		}.ConsolidateGameInfo()

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

func (t CfbTeamsExtract) ExtractTeams() []team.Team {
	weekSchedule := schedule.AllCfbScheduleInfo{Week: t.Week}.ConsolidateScheduleInfo()
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.AllCfbTeamInfo{TeamCode: espnTeamCode1}.ConsolidateTeamInfo()
		team2 := team.AllCfbTeamInfo{TeamCode: espnTeamCode2}.ConsolidateTeamInfo()

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

func (t NflTeamsExtract) ExtractTeams() []team.Team {
	weekSchedule := schedule.AllNflScheduleInfo{Week: t.Week}.ConsolidateScheduleInfo()
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.AllNflTeamInfo{TeamCode: espnTeamCode1}.ConsolidateTeamInfo()
		team2 := team.AllNflTeamInfo{TeamCode: espnTeamCode2}.ConsolidateTeamInfo()

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

func (l LocationsExtract) ExtractLocations() []location.OpencageEndpoint {
	weekSchedule := schedule.AllCfbScheduleInfo{Week: l.Week}.ConsolidateScheduleInfo()
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var locationsThisWeek []location.OpencageEndpoint

	for i := range espnWeekGames {
		var stadium string = espnWeekGames[i].Competitions[0].Venue.FullName
		var city string = espnWeekGames[i].Competitions[0].Venue.Address.City
		var state string = espnWeekGames[i].Competitions[0].Venue.Address.State

		location := location.OpencageLocation{Stadium: stadium, City: city, State: state}.GetLocationDetails()
		locationsThisWeek = append(locationsThisWeek, location)
	}

	return locationsThisWeek
}
