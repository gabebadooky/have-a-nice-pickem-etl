package bettingodds

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/utils"
)

type EspnAwayBettingOdds struct {
	game.Game
}

type EspnHomeBettingOdds struct {
	game.Game
}

func (e EspnAwayBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.ESPN.Pickcenter[0].OverUnder
	return overUnder
}

func (e EspnHomeBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.ESPN.Pickcenter[1].OverUnder
	return overUnder
}

func (e EspnAwayBettingOdds) parseMoneyline() int {
	var awayMoneyline int = e.ESPN.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

func (e EspnHomeBettingOdds) parseMoneyline() int {
	var homeMoneyline int = e.ESPN.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

func (e EspnAwayBettingOdds) parseSpread() float32 {
	var awaySpread float32 = e.ESPN.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

func (e EspnHomeBettingOdds) parseSpread() float32 {
	var homeSpread float32 = e.ESPN.Pickcenter[0].Spread
	return homeSpread
}

func (e EspnAwayBettingOdds) parseWinProbability() float32 {
	var winProbabilityText string = e.ESPN.Predictor.AwayTeam.GameProjection
	winProbabilityFloat32 := utils.ConvertStringToFloat32(winProbabilityText)
	return winProbabilityFloat32
}

func (e EspnHomeBettingOdds) parseWinProbability() float32 {
	var winProbabilityText string = e.ESPN.Predictor.HomeTeam.GameProjection
	winProbabilityFloat32 := utils.ConvertStringToFloat32(winProbabilityText)
	return winProbabilityFloat32
}
