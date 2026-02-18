// Package bettingodds provides ESPN betting odds parsing functionality.
// It extracts betting odds data from ESPN Game Summary API responses including
// spread, moneyline, over/under, and win probability for both away and home teams.
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

// parseOverUnder returns the over/under total from the ESPN pickcenter (away or home row).
func (e EspnAwayBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.ESPN.Pickcenter[0].OverUnder
	return overUnder
}

// parseOverUnder returns the over/under total from the ESPN pickcenter for the home team.
func (e EspnHomeBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.ESPN.Pickcenter[1].OverUnder
	return overUnder
}

// parseMoneyline returns the away team moneyline from the ESPN pickcenter.
func (e EspnAwayBettingOdds) parseMoneyline() int {
	var awayMoneyline int = e.ESPN.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

// parseMoneyline returns the home team moneyline from the ESPN pickcenter.
func (e EspnHomeBettingOdds) parseMoneyline() int {
	var homeMoneyline int = e.ESPN.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

// parseSpread returns the away team spread from the ESPN pickcenter (negated for away).
func (e EspnAwayBettingOdds) parseSpread() float32 {
	var awaySpread float32 = e.ESPN.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

// parseSpread returns the home team spread from the ESPN pickcenter.
func (e EspnHomeBettingOdds) parseSpread() float32 {
	var homeSpread float32 = e.ESPN.Pickcenter[0].Spread
	return homeSpread
}

// parseWinProbability returns the away team win probability from the ESPN predictor.
func (e EspnAwayBettingOdds) parseWinProbability() float32 {
	var winProbabilityText string = e.ESPN.Predictor.AwayTeam.GameProjection
	winProbabilityFloat32 := utils.ConvertStringToFloat32(winProbabilityText)
	return winProbabilityFloat32
}

// parseWinProbability returns the home team win probability from the ESPN predictor.
func (e EspnHomeBettingOdds) parseWinProbability() float32 {
	var winProbabilityText string = e.ESPN.Predictor.HomeTeam.GameProjection
	winProbabilityFloat32 := utils.ConvertStringToFloat32(winProbabilityText)
	return winProbabilityFloat32
}
