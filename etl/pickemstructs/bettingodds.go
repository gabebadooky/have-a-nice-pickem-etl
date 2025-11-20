package pickemstructs

type BettingOdds struct {
	GameID            string
	Source            string
	OverUnder         float32
	AwayMoneyline     uint16
	HomeMoneyline     uint16
	AwaySpread        float32
	HomeSpread        float32
	AwayWinPercentage string
	HomeWinPercentage string
}
