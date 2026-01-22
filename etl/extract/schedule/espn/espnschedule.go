/*
Package that:

  - Calls the Scoreboard endpoint (espnHiddenScoreboardBaseURL),
    for a given week, from the ESPN hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response
*/
package espnschedule

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnSchedule interface {
	scheduleForWeek() ScoreboardEndpoint
}

type CfbEspnSchedule struct {
	Week uint
}

type NflEspnSchedule struct {
	Week uint
}

func GetScheduleForWeek(s EspnSchedule) ScoreboardEndpoint {
	return s.scheduleForWeek()
}

func makeAndHandleScoreboardEndpointCall(week uint, espnScoreboardEndpoint string) ScoreboardEndpoint {
	log.Printf("\nCalling Scoreboard endpoint for week %d: %s\n", week, espnScoreboardEndpoint)

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

// Call CFB ESPN Scoreboard Summary API Endpoint
func (e CfbEspnSchedule) scheduleForWeek() ScoreboardEndpoint {
	var espnScoreboardEndpoint string

	if e.Week <= utils.CFB_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL, e.Week)
	} else {
		// SeasonType 3 only has one week for CFB Postseason
		//espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_CFB_POST_SEASON_SCHEDULE_URL, 1)
		espnScoreboardEndpoint = fmt.Sprintf("%s", utils.ESPN_CFB_POST_SEASON_SCHEDULE_URL)
	}

	espnScoreboard := makeAndHandleScoreboardEndpointCall(e.Week, espnScoreboardEndpoint)
	return espnScoreboard
}

// Call NFL ESPN Scoreboard Summary API Endpoint
func (e NflEspnSchedule) scheduleForWeek() ScoreboardEndpoint {
	var espnScoreboardEndpoint string

	if e.Week <= utils.NFL_REG_SEASON_WEEKS {
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL, e.Week)
	} else {
		// SeasonType 3 weeks begin at 1
		var postSeasonWeek uint = e.Week - utils.NFL_REG_SEASON_WEEKS
		espnScoreboardEndpoint = fmt.Sprintf("%s%d", utils.ESPN_NFL_POST_SEASON_SCHEDULE_URL, postSeasonWeek)
	}

	espnScoreboard := makeAndHandleScoreboardEndpointCall(e.Week, espnScoreboardEndpoint)
	return espnScoreboard
}
