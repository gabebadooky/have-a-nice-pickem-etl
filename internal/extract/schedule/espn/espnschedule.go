// Package espnschedule provides ESPN schedule API client functionality.
// It calls the ESPN Scoreboard API endpoint to retrieve schedule data for both
// college football (CFB) and NFL, handling both regular season and postseason schedules.
package espnschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
)

type CfbEspnSchedule struct {
	Week uint
}

type NflEspnSchedule struct {
	Week uint
}

// fetchEspnSchedule calls the ESPN scoreboard endpoint and decodes the JSON response.
func fetchEspnSchedule(espnScoreboardEndpoint string) ScoreboardEndpoint {
	log.Printf("\nCalling Scoreboard endpoint: %s\n", espnScoreboardEndpoint)

	body, err := utils.CallEndpoint(espnScoreboardEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	scheduleDetails, err := utils.DecodeJSON[ScoreboardEndpoint](body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return scheduleDetails
}

// instantiateScoreboardEndpoint returns the ESPN college football scoreboard URL for the configured week.
func (sched CfbEspnSchedule) setScoreboardEndpoint() string {
	var espnScoreboardEndpoint string

	if sched.Week <= utils.CFB_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	} else {
		// SeasonType 3 only has one week for CFB Postseason
		espnScoreboardEndpoint = fmt.Sprintf("%s", utils.ESPN_CFB_POST_SEASON_SCHEDULE_URL)
	}

	return espnScoreboardEndpoint
}

// instantiateScoreboardEndpoint returns the ESPN NFL scoreboard URL for the configured week.
func (sched NflEspnSchedule) setScoreboardEndpoint() string {
	var espnScoreboardEndpoint string

	if sched.Week <= utils.NFL_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL, sched.Week)
	} else {
		// SeasonType 3 weeks begin at 1
		var postSeasonWeek uint = sched.Week - utils.NFL_REG_SEASON_WEEKS
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_POST_SEASON_SCHEDULE_URL, postSeasonWeek)
	}

	return espnScoreboardEndpoint
}

// callSchedule fetches the ESPN college football scoreboard for the configured week.
func (sched CfbEspnSchedule) GetSchedule() ScoreboardEndpoint {
	espnScoreboardEndpoint := sched.setScoreboardEndpoint()
	espnScoreboard := fetchEspnSchedule(espnScoreboardEndpoint)
	return espnScoreboard
}

// callSchedule fetches the ESPN NFL scoreboard for the configured week.
func (sched NflEspnSchedule) GetSchedule() ScoreboardEndpoint {
	espnScoreboardEndpoint := sched.setScoreboardEndpoint()
	espnScoreboard := fetchEspnSchedule(espnScoreboardEndpoint)
	return espnScoreboard
}
