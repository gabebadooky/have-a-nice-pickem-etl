package bettingodds

import "have-a-nice-pickem-etl/etl/pickemstructs"

func ParseOverUnder(espnGameDetails pickemstructs.ESPNGameDetailsResponse) float32 {
	var overUnder float32 = espnGameDetails.Pickcenter[0].OverUnder
	return overUnder
}

func ParseAwayMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse) uint16 {
	var awayMoneyline uint16 = espnGameDetails.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

func ParseHomeMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse) uint16 {
	var homeMoneyline uint16 = espnGameDetails.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

func ParseAwaySpread(espnGameDetails pickemstructs.ESPNGameDetailsResponse) float32 {
	var awaySpread float32 = espnGameDetails.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

func ParseHomeSpread(espnGameDetails pickemstructs.ESPNGameDetailsResponse) float32 {
	var homeSpread float32 = espnGameDetails.Pickcenter[0].Spread
	return homeSpread
}

func ParseAwayWinPercentage(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var awayWinPercentage string = espnGameDetails.Predictor.AwayTeam.GameProjection
	return awayWinPercentage
}

func ParseHomeWinPercentage(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var homeWinPercentage string = espnGameDetails.Predictor.HomeTeam.GameProjection
	return homeWinPercentage
}
