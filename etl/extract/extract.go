package extract

import (
	"have-a-nice-pickem-etl/etl/extract/cbs"
	"have-a-nice-pickem-etl/etl/extract/espn"
	"have-a-nice-pickem-etl/etl/extract/fox"
	"have-a-nice-pickem-etl/etl/sharedtypes"

	"github.com/PuerkitoBio/goquery"
)

type CFBSchedule struct {
	Week uint8
}

type NFLSchedule struct {
	Week uint8
}

type CFBGame struct {
	EspnCode string
	CBS      *goquery.Selection
	FOX      *goquery.Selection
}

type NFLGame struct {
	EspnCode string
	CBS      *goquery.Selection
	FOX      *goquery.Selection
}

type CFBTeam struct {
	EspnCode string
}

type NFLTeam struct {
	EspnCode string
}

type ConsolidatedSchedule struct {
	ESPN sharedtypes.ESPNScheduleResponse
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

type ConsolidatedGame struct {
	ESPN sharedtypes.ESPNGameDetailsResponse
	CBS  *goquery.Selection
	FOX  *goquery.Selection
}

func (cfb CFBSchedule) ExtractSchedule() ConsolidatedSchedule {
	return ConsolidatedSchedule{
		ESPN: espn.EspnCFBSchedule{Week: cfb.Week}.GetScheduleForWeek(),
		CBS:  cbs.CbsCFBSchedule{Week: cfb.Week}.GetScheduleForWeek(),
		FOX:  fox.FoxCFBSchedule{Week: cfb.Week}.GetScheduleForWeek(),
	}
}

func (nfl NFLSchedule) ExtractSchedule() ConsolidatedSchedule {
	return ConsolidatedSchedule{
		ESPN: espn.EspnNFLSchedule{Week: nfl.Week}.GetScheduleForWeek(),
		CBS:  cbs.CbsNFLSchedule{Week: nfl.Week}.GetScheduleForWeek(),
		FOX:  fox.FoxNFLSchedule{Week: nfl.Week}.GetScheduleForWeek(),
	}
}

func (cfb CFBGame) ExtractGame() ConsolidatedGame {
	return ConsolidatedGame{
		ESPN: espn.EspnCFBGame{GameCode: cfb.EspnCode}.GetGameSummary(),
	}
}

func (nfl NFLGame) ExtractGame() ConsolidatedGame {
	return ConsolidatedGame{
		ESPN: espn.EspnCFBGame{GameCode: nfl.EspnCode}.GetGameSummary(),
	}
}

func (cfb CFBTeam) ExtractGame() ConsolidatedGame {
	return ConsolidatedGame{
		ESPN: espn.EspnCFBGame{GameCode: cfb.EspnCode}.GetGameSummary(),
	}
}

func (nfl NFLTeam) ExtractGame() ConsolidatedGame {
	return ConsolidatedGame{
		ESPN: espn.EspnCFBGame{GameCode: nfl.EspnCode}.GetGameSummary(),
	}
}
