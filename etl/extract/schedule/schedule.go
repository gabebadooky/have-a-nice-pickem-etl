package schedule

import (
	"have-a-nice-pickem-etl/etl/extract/schedule/cbs"
	"have-a-nice-pickem-etl/etl/extract/schedule/espn"
	"have-a-nice-pickem-etl/etl/extract/schedule/fox"

	"github.com/PuerkitoBio/goquery"
)

type AllScheduleInfo interface {
	ConsolidateScheduleInfo() Schedule
}

type AllCfbScheduleInfo struct {
	Week uint8
}

type AllNflScheduleInfo struct {
	Week uint8
}

type Schedule struct {
	ESPN espn.ScoreboardEndpoint
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

func (c AllCfbScheduleInfo) ConsolidateScheduleInfo() Schedule {
	var EspnSchedule espn.ScoreboardEndpoint = espn.CfbEspnSchedule{Week: c.Week}.GetScheduleForWeek()
	var CbsSchedule *goquery.Selection = cbs.CbsCfbSchedule{Week: c.Week}.GetScheduleForWeek()
	var FoxSchedule *goquery.Selection = fox.FoxCfbSchedule{Week: c.Week}.GetScheduleForWeek()

	return Schedule{
		ESPN: EspnSchedule,
		CBS:  CbsSchedule,
		FOX:  FoxSchedule,
	}
}

func (n AllNflScheduleInfo) ConsolidateScheduleInfo() Schedule {
	var EspnSchedule espn.ScoreboardEndpoint = espn.NflEspnSchedule{Week: n.Week}.GetScheduleForWeek()
	var CbsSchedule *goquery.Selection = cbs.CbsNflSchedule{Week: n.Week}.GetScheduleForWeek()
	var FoxSchedule *goquery.Selection = fox.FoxNflSchedule{Week: n.Week}.GetScheduleForWeek()

	return Schedule{
		ESPN: EspnSchedule,
		CBS:  CbsSchedule,
		FOX:  FoxSchedule,
	}
}
