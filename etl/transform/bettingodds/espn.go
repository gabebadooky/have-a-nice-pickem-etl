package bettingodds

func (e EspnBettingOdds) parseOverUnder() float32 {
	var overUnder float32 = e.GameExtract.ESPN.Pickcenter[0].OverUnder
	return overUnder
}

func (e EspnBettingOdds) parseAwayMoneyline() int {
	var awayMoneyline int = e.GameExtract.ESPN.Pickcenter[0].AwayTeamOdds.Moneyline
	return awayMoneyline
}

func (e EspnBettingOdds) parseHomeMoneyline() int {
	var homeMoneyline int = e.GameExtract.ESPN.Pickcenter[0].HomeTeamOdds.Moneyline
	return homeMoneyline
}

func (e EspnBettingOdds) parseAwaySpread() float32 {
	var awaySpread float32 = e.GameExtract.ESPN.Pickcenter[0].Spread
	if awaySpread > 0 {
		awaySpread -= (awaySpread * 2)
	} else {
		awaySpread += (awaySpread * 2)
	}
	return awaySpread
}

func (e EspnBettingOdds) parseHomeSpread() float32 {
	var homeSpread float32 = e.GameExtract.ESPN.Pickcenter[0].Spread
	return homeSpread
}

func (e EspnBettingOdds) parseAwayWinPercentage() string {
	var awayWinPercentage string = e.GameExtract.ESPN.Predictor.AwayTeam.GameProjection
	return awayWinPercentage
}

func (e EspnBettingOdds) parseHomeWinPercentage() string {
	var homeWinPercentage string = e.GameExtract.ESPN.Predictor.HomeTeam.GameProjection
	return homeWinPercentage
}
