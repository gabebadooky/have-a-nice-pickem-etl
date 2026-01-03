package bettingodds

import (
	"have-a-nice-pickem-etl/etl/extract/game"
)

type Instantiator interface {
	instantiate() BettingOdds
}

type EspnBettingOdds struct {
	GameExtract game.Game
}

type CbsBettingOdds struct {
	GameExtract game.Game
}

type FoxBettingOdds struct {
	GameExtract game.Game
}

type BettingOdds struct {
	GameID            string
	Source            string
	OverUnder         float32
	AwayMoneyline     int
	HomeMoneyline     int
	AwaySpread        float32
	HomeSpread        float32
	AwayWinPercentage string
	HomeWinPercentage string
}

func InstantiateBettingOdds(i Instantiator) BettingOdds {
	return i.instantiate()
}

// Instantiates ESPN Betting Odds record from various sources
func (b EspnBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:        b.GameExtract.GameID,
		Source:        "ESPN",
		OverUnder:     b.parseOverUnder(),
		AwayMoneyline: b.parseAwayMoneyline(),
		HomeMoneyline: b.parseHomeMoneyline(),
		AwaySpread:    b.parseAwaySpread(),
		HomeSpread:    b.parseHomeSpread(),
	}
}

// Instantiates CBS Betting Odds record from various sources
func (b CbsBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:            b.GameExtract.GameID,
		Source:            "CBS",
		OverUnder:         b.parseOverUnder(),
		AwayMoneyline:     b.parseAwayMoneyline(),
		HomeMoneyline:     b.parseHomeMoneyline(),
		AwaySpread:        b.parseAwaySpread(),
		HomeSpread:        b.parseHomeSpread(),
		AwayWinPercentage: b.parseAwayWinPercentage(),
		HomeWinPercentage: b.parseHomeWinPercentage(),
	}
}
