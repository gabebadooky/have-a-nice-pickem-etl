package extract

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/extract/location"
	"have-a-nice-pickem-etl/etl/extract/schedule"
	espnsched "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/extract/team"
)

type GamesExtract interface {
	allGames() []game.Game
}

type TeamsExtract interface {
	allTeams() []team.Team
}

type LocationsExtract interface {
	allLocations() []location.Location
}

type CfbGamesExtract struct {
	Week uint
}

type NflGamesExtract struct {
	Week uint
}

type CfbTeamsExtract struct {
	Week uint
}

type NflTeamsExtract struct {
	Week uint
}

type CfbLocationsExtract struct {
	Week uint
}

type NflLocationsExtract struct {
	Week uint
}

func ExtractGames(g GamesExtract) []game.Game {
	return g.allGames()
}

func ExtractTeams(t TeamsExtract) []team.Team {
	return t.allTeams()
}

func ExtractLocations(l LocationsExtract) []location.Location {
	return l.allLocations()
}

func (g CfbGamesExtract) allGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllCfbScheduleInfo{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game := game.ConsolidateGameInfo(game.AllCfbGameInfo{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		})

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

func (g NflGamesExtract) allGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllNflScheduleInfo{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game := game.ConsolidateGameInfo(game.AllNflGameInfo{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		})

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

func (t CfbTeamsExtract) allTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllCfbScheduleInfo{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.ConsolidateTeamInfo(team.AllCfbTeamInfo{TeamCode: espnTeamCode1})
		team2 := team.ConsolidateTeamInfo(team.AllCfbTeamInfo{TeamCode: espnTeamCode2})

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

func (t NflTeamsExtract) allTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllNflScheduleInfo{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.ConsolidateTeamInfo(team.AllNflTeamInfo{TeamCode: espnTeamCode1})
		team2 := team.ConsolidateTeamInfo(team.AllNflTeamInfo{TeamCode: espnTeamCode2})

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

func (l CfbLocationsExtract) allLocations() []location.Location {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllCfbScheduleInfo{Week: l.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var locationsThisWeek []location.Location

	for i := range espnWeekGames {
		var stadium string = espnWeekGames[i].Competitions[0].Venue.FullName
		var city string = espnWeekGames[i].Competitions[0].Venue.Address.City
		var state string = espnWeekGames[i].Competitions[0].Venue.Address.State

		//location := location.OpencageLocation{Stadium: stadium, City: city, State: state}.GetLocationDetails()
		locationDetails := location.ConsolidateLocationInfo(
			location.OpencageLocationInfo{
				Stadium: stadium,
				City:    city,
				State:   state,
			},
		)
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}

func (l NflLocationsExtract) allLocations() []location.Location {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.AllNflScheduleInfo{Week: l.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var locationsThisWeek []location.Location

	for i := range espnWeekGames {
		var stadium string = espnWeekGames[i].Competitions[0].Venue.FullName
		var city string = espnWeekGames[i].Competitions[0].Venue.Address.City
		var state string = espnWeekGames[i].Competitions[0].Venue.Address.State

		//location := location.OpencageLocation{Stadium: stadium, City: city, State: state}.GetLocationDetails()
		locationDetails := location.ConsolidateLocationInfo(
			location.OpencageLocationInfo{
				Stadium: stadium,
				City:    city,
				State:   state,
			},
		)
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}
