package bettingodds

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
)

type Instantiator interface {
	instantiate() BettingOdds
}

type VegasAwayBettingOdds struct {
	game.Game
}

type VegasHomeBettingOdds struct {
	game.Game
}

type BettingOdds struct {
	GameID         string
	TeamID         string
	Source         string
	OverUnder      float32
	Moneyline      int
	Spread         float32
	WinProbability float32
}

func InstantiateBettingOdds(i Instantiator) BettingOdds {
	return i.instantiate()
}

// Instantiates Away Team ESPN Betting Odds record
func (b EspnAwayBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseAwayTeamID(b.Game),
		Source:         "ESPN",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// Instantiates Home Team ESPN Betting Odds record
func (b EspnHomeBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseHomeTeamID(b.Game),
		Source:         "ESPN",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// Instantiates Away Team CBS Betting Odds record
func (b CbsAwayBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		Source:         "CBS",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// Instantiates Home Team CBS Betting Odds record
func (b CbsHomeBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		Source:         "CBS",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// Instantiates Away Team Fox Betting Odds record
func (b FoxAwayBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		Source:         "FOX",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// Instantiates Home Team Fox Betting Odds record
func (b FoxHomeBettingOdds) instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		Source:         "FOX",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}
