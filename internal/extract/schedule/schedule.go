// Package schedule provides schedule extraction functionality that consolidates
// schedule data from multiple sources (ESPN API, CBS web scraping, Fox web scraping)
// for both college football (CFB) and NFL seasons.
package schedule

import (
	cbsschedule "have-a-nice-pickem-etl/internal/extract/schedule/cbs"
	espnschedule "have-a-nice-pickem-etl/internal/extract/schedule/espn"
	foxschedule "have-a-nice-pickem-etl/internal/extract/schedule/fox"

	"github.com/PuerkitoBio/goquery"
)

type CfbSchedule struct {
	Week uint
}

type NflSchedule struct {
	Week uint
}

type Schedule struct {
	ESPN espnschedule.ScoreboardEndpoint
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

type scheduleInstantiator interface {
	extractSchedule() Schedule
}

// ConsolidateScheduleInfo runs the given schedule instantiator and returns the consolidated Schedule.
func ConsolidateScheduleInfo(s scheduleInstantiator) Schedule {
	return s.extractSchedule()
}

// extractSchedule fetches college football schedule data from ESPN, CBS, and Fox for the week.
func (c CfbSchedule) extractSchedule() Schedule {
	return Schedule{
		ESPN: espnschedule.CfbEspnSchedule{Week: c.Week}.GetSchedule(),
		CBS:  cbsschedule.CbsCfbSchedule{Week: c.Week}.GetSchedule(),
		FOX:  foxschedule.FoxCfbSchedule{Week: c.Week}.GetSchedule(),
	}
}

// extractSchedule fetches NFL schedule data from ESPN, CBS, and Fox for the week.
func (n NflSchedule) extractSchedule() Schedule {
	return Schedule{
		ESPN: espnschedule.NflEspnSchedule{Week: n.Week}.GetSchedule(),
		CBS:  cbsschedule.CbsNflSchedule{Week: n.Week}.GetSchedule(),
		FOX:  foxschedule.FoxNflSchedule{Week: n.Week}.GetSchedule(),
	}
}
