// Package bettingodds provides betting odds transformation functionality that extracts
// and structures betting data (spread, moneyline, over/under, win probability) from
// multiple sources (ESPN, CBS, Fox) for both away and home teams.
package bettingodds

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
	"time"
)

type VegasAwayBettingOdds struct {
	game.Game
}

type VegasHomeBettingOdds struct {
	game.Game
}

type BettingOdds1 struct {
	GameID         string
	TeamID         string
	Source         string
	OverUnder      float32
	Moneyline      int
	Spread         float32
	WinProbability float32
}

type BettingOdds struct {
	GameID         string    `gorm:"column:game"`
	TeamID         string    `gorm:"column:team"`
	Source         string    `gorm:"column:source"`
	OverUnder      float32   `gorm:"column:over_under"`
	Moneyline      int       `gorm:"column:moneyline"`
	Spread         float32   `gorm:"column:spread"`
	WinProbability float32   `gorm:"column:win_probability"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (BettingOdds) TableName() string {
	return "betting_odds"
}

// instantiate builds ESPN away team betting odds from the game data.
func (b EspnAwayBettingOdds) Instantiate() BettingOdds {
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

// instantiate builds ESPN home team betting odds from the game data.
func (b EspnHomeBettingOdds) Instantiate() BettingOdds {
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

// instantiate builds CBS away team betting odds from the game's CBS odds page.
func (b CbsAwayBettingOdds) Instantiate() BettingOdds {

	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseAwayTeamID(b.Game),
		Source:         "CBS",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// instantiate builds CBS home team betting odds from the game's CBS odds page.
func (b CbsHomeBettingOdds) Instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseHomeTeamID(b.Game),
		Source:         "CBS",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// instantiate builds Fox away team betting odds from the game's Fox odds page.
func (b FoxAwayBettingOdds) Instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseAwayTeamID(b.Game),
		Source:         "FOX",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}

// instantiate builds Fox home team betting odds from the game's Fox odds page.
func (b FoxHomeBettingOdds) Instantiate() BettingOdds {
	return BettingOdds{
		GameID:         b.GameID,
		TeamID:         common.ParseHomeTeamID(b.Game),
		Source:         "FOX",
		OverUnder:      b.parseOverUnder(),
		Moneyline:      b.parseMoneyline(),
		Spread:         b.parseSpread(),
		WinProbability: b.parseWinProbability(),
	}
}
