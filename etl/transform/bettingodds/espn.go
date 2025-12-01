package bettingodds

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
)

type EspnOdds struct {
	pickemstructs.BettingOdds
}

func (e EspnOdds) ParseOverUnder(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) float32 {
	var overUnder float32 = consolidatedGameProperties.EspnDetails.Pickcenter[0].OverUnder
	return overUnder
}

func (e EspnOdds) ParseAwayMoneyline(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) int16 {
	var awayMoneyline int16 = consolidatedGameProperties.EspnDetails.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

func (e EspnOdds) ParseHomeMoneyline(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) int16 {
	var homeMoneyline int16 = consolidatedGameProperties.EspnDetails.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

func (e EspnOdds) ParseAwaySpread(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) float32 {
	var awaySpread float32 = consolidatedGameProperties.EspnDetails.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

func (e EspnOdds) ParseHomeSpread(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) float32 {
	var homeSpread float32 = consolidatedGameProperties.EspnDetails.Pickcenter[0].Spread
	return homeSpread
}

func (e EspnOdds) ParseAwayWinPercentage(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var awayWinPercentage string = consolidatedGameProperties.EspnDetails.Predictor.AwayTeam.GameProjection
	return awayWinPercentage
}

func (e EspnOdds) ParseHomeWinPercentage(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var homeWinPercentage string = consolidatedGameProperties.EspnDetails.Predictor.HomeTeam.GameProjection
	return homeWinPercentage
}
