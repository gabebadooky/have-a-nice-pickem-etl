package schedule

import (
	cbsschedule "have-a-nice-pickem-etl/etl/extract/schedule/cbs"
	espnschedule "have-a-nice-pickem-etl/etl/extract/schedule/espn"
	foxschedule "have-a-nice-pickem-etl/etl/extract/schedule/fox"

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

func ConsolidateScheduleInfo(s scheduleInstantiator) Schedule {
	return s.extractSchedule()
}

func (c CfbSchedule) extractSchedule() Schedule {
	var EspnSchedule espnschedule.ScoreboardEndpoint = espnschedule.GetScheduleForWeek(espnschedule.CfbEspnSchedule{Week: c.Week})
	var CbsSchedule *goquery.Selection = cbsschedule.GetScheduleForWeek(cbsschedule.CbsCfbSchedule{Week: c.Week})
	var FoxSchedule *goquery.Selection = foxschedule.GetScheduleForWeek(foxschedule.FoxCfbSchedule{Week: c.Week})

	return Schedule{
		ESPN: EspnSchedule,
		CBS:  CbsSchedule,
		FOX:  FoxSchedule,
	}
}

func (n NflSchedule) extractSchedule() Schedule {
	var EspnSchedule espnschedule.ScoreboardEndpoint = espnschedule.GetScheduleForWeek(espnschedule.NflEspnSchedule{Week: n.Week})
	var CbsSchedule *goquery.Selection = cbsschedule.GetScheduleForWeek(cbsschedule.CbsNflSchedule{Week: n.Week})
	var FoxSchedule *goquery.Selection = foxschedule.GetScheduleForWeek(foxschedule.FoxNflSchedule{Week: n.Week})

	return Schedule{
		ESPN: EspnSchedule,
		CBS:  CbsSchedule,
		FOX:  FoxSchedule,
	}
}
