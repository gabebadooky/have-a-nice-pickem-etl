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

type CfbExtract struct {
	Week uint
}

type NflExtract struct {
	Week uint
}

type Extract struct {
	GamesExtract     []game.Game
	TeamsExtract     []team.Team
	LocationsExtract []location.Location
}

// getGames fetches and consolidates college football games for the configured week.
func (g CfbExtract) compileGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.CfbSchedule{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game, err := game.CfbGame{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		}.ExtractGame()

		if err != nil {
			fmt.Printf("Skipping GameID that contatins \"tbd\"")
		}

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

// getGames fetches and consolidates NFL games for the configured week.
func (g NflExtract) compileGames() []game.Game {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.NflSchedule{Week: g.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var gamesThisWeek []game.Game

	for i := range espnWeekGames {
		var espnEvent espnsched.EventProperty = espnWeekGames[i]

		game, err := game.NflGame{
			EspnEvent:       espnEvent,
			CbsSchedulePage: weekSchedule.CBS,
			FoxSchedulePage: weekSchedule.FOX,
		}.ExtractGame()

		if err != nil {
			fmt.Printf("Skipping GameID that contatins \"tbd\"")
		}

		gamesThisWeek = append(gamesThisWeek, game)
	}

	return gamesThisWeek
}

// getTeams fetches and consolidates college football teams for the configured week.
func (t CfbExtract) compileTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.CfbSchedule{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.CfbTeam{EspnCode: espnTeamCode1}.ExtractTeam()
		team2 := team.CfbTeam{EspnCode: espnTeamCode2}.ExtractTeam()

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

// getTeams fetches and consolidates NFL teams for the configured week.
func (t NflExtract) compileTeams() []team.Team {
	weekSchedule := schedule.ConsolidateScheduleInfo(schedule.NflSchedule{Week: t.Week})
	var espnWeekGames []espnsched.EventProperty = weekSchedule.ESPN.Events
	var teamsThisWeek []team.Team

	for i := range espnWeekGames {
		espnTeamCode1 := espnWeekGames[i].Competitions[0].Competitors[0].ID
		espnTeamCode2 := espnWeekGames[i].Competitions[0].Competitors[1].ID

		team1 := team.NflTeam{EspnCode: espnTeamCode1}.ExtractTeam()
		team2 := team.NflTeam{EspnCode: espnTeamCode2}.ExtractTeam()

		teamsThisWeek = append(teamsThisWeek, team1)
		teamsThisWeek = append(teamsThisWeek, team2)
	}

	return teamsThisWeek
}

// getLocations fetches and consolidates college football game locations for the configured week.
func (l CfbExtract) compileLocations() []location.Location {
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

		locationDetails := opencageLocation.GetLocation()
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}

// getLocations fetches and consolidates NFL game locations for the configured week.
func (l NflExtract) compileLocations() []location.Location {
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

		locationDetails := opencageLocation.GetLocation()
		locationsThisWeek = append(locationsThisWeek, locationDetails)
	}

	return locationsThisWeek
}

func (e CfbExtract) PerformExtract() Extract {
	weekExtract := CfbExtract{Week: e.Week}

	return Extract{
		GamesExtract:     weekExtract.compileGames(),
		TeamsExtract:     weekExtract.compileTeams(),
		LocationsExtract: weekExtract.compileLocations(),
	}
}
