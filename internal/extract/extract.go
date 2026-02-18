// Package extract provides the main extraction orchestration layer for the ETL pipeline.
// It coordinates the extraction of games, teams, and locations from multiple data sources
// (ESPN, CBS, Fox) and consolidates them into unified data structures.
package extract

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/extract/location"
	"have-a-nice-pickem-etl/internal/extract/schedule"
	espnsched "have-a-nice-pickem-etl/internal/extract/schedule/espn"
	"have-a-nice-pickem-etl/internal/extract/team"
)

type GamesExtractor interface {
	getGames() []game.Game
}

type TeamsExtractor interface {
	getTeams() []team.Team
}

type LocationsExtractor interface {
	getLocations() []location.Location
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

// ExtractGames runs the given GamesExtractor and returns the consolidated games.
func ExtractGames(g GamesExtractor) []game.Game {
	return g.getGames()
}

// ExtractTeams runs the given TeamsExtractor and returns the consolidated teams.
func ExtractTeams(t TeamsExtractor) []team.Team {
	return t.getTeams()
}

// ExtractLocations runs the given LocationsExtractor and returns the consolidated locations.
func ExtractLocations(l LocationsExtractor) []location.Location {
	return l.getLocations()
}

// getGames fetches and consolidates college football games for the configured week.
func (g CfbGamesExtract) getGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.CfbSchedule{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game, err := game.ConsolidateGameInfo(game.CfbGame{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		})
		if err != nil {
			fmt.Printf("Skipping GameID that contatins \"tbd\"")
		}

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

// getGames fetches and consolidates NFL games for the configured week.
func (g NflGamesExtract) getGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.NflSchedule{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game, err := game.ConsolidateGameInfo(game.NflGame{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		})
		if err != nil {
			fmt.Printf("Skipping GameID that contatins \"tbd\"")
		}

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

// getTeams fetches and consolidates college football teams for the configured week.
func (t CfbTeamsExtract) getTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.CfbSchedule{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team
	fmt.Printf("len(espnWeekGames): %d", len(espnWeekGames))
	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.ConsolidateTeamInfo(team.CfbTeam{EspnCode: espnTeamCode1})
		team2 := team.ConsolidateTeamInfo(team.CfbTeam{EspnCode: espnTeamCode2})

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

// getTeams fetches and consolidates NFL teams for the configured week.
func (t NflTeamsExtract) getTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.NflSchedule{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.ConsolidateTeamInfo(team.NflTeam{EspnCode: espnTeamCode1})
		team2 := team.ConsolidateTeamInfo(team.NflTeam{EspnCode: espnTeamCode2})

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

// getLocations fetches and consolidates college football game locations for the configured week.
func (l CfbLocationsExtract) getLocations() []location.Location {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.CfbSchedule{Week: l.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var locationsThisWeek []location.Location

	for i := range espnWeekGames {
		var stadium string = espnWeekGames[i].Competitions[0].Venue.FullName
		var city string = espnWeekGames[i].Competitions[0].Venue.Address.City
		var state string = espnWeekGames[i].Competitions[0].Venue.Address.State

		opencageLocation := location.OpencageLocation{
			Stadium: stadium,
			City:    city,
			State:   state,
		}

		locationDetails := location.ConsolidateLocationInfo(opencageLocation)
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}

// getLocations fetches and consolidates NFL game locations for the configured week.
func (l NflLocationsExtract) getLocations() []location.Location {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.NflSchedule{Week: l.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var locationsThisWeek []location.Location

	for i := range espnWeekGames {
		var stadium string = espnWeekGames[i].Competitions[0].Venue.FullName
		var city string = espnWeekGames[i].Competitions[0].Venue.Address.City
		var state string = espnWeekGames[i].Competitions[0].Venue.Address.State

		opencageLocation := location.OpencageLocation{
			Stadium: stadium,
			City:    city,
			State:   state,
		}

		locationDetails := location.ConsolidateLocationInfo(opencageLocation)
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}
