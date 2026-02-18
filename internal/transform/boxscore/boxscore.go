// Package boxscore provides boxscore transformation functionality that extracts
// and structures scoring data (quarter scores, total scores, overtime scores)
// from ESPN game data for both away and home teams.
package boxscore

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
)

type Instantiator interface {
	instantiate() Boxscore
}

type AwayBoxscore struct {
	game.Game
}

type HomeBoxscore struct {
	game.Game
}

type Boxscore struct {
	GameID        string
	TeamID        string
	Q1Score       uint
	Q2Score       uint
	Q3Score       uint
	Q4Score       uint
	OvertimeScore uint
	TotalScore    uint
}

// InstantiateBoxscore runs the given instantiator and returns the boxscore.
func InstantiateBoxscore(i Instantiator) Boxscore {
	return i.instantiate()
}

// instantiate builds the away team boxscore from the game's ESPN linescore data.
func (a AwayBoxscore) instantiate() Boxscore {
	return Boxscore{
		GameID:        a.GameID,
		TeamID:        common.ParseAwayTeamID(a.Game),
		Q1Score:       a.parseQuarterScore(1),
		Q2Score:       a.parseQuarterScore(2),
		Q3Score:       a.parseQuarterScore(3),
		Q4Score:       a.parseQuarterScore(4),
		OvertimeScore: a.ParseOvertimeScore(),
		TotalScore:    a.parseTotalScore(),
	}
}

// instantiate builds the home team boxscore from the game's ESPN linescore data.
func (h HomeBoxscore) instantiate() Boxscore {
	return Boxscore{
		GameID:        h.GameID,
		TeamID:        common.ParseHomeTeamID(h.Game),
		Q1Score:       h.parseQuarterScore(1),
		Q2Score:       h.parseQuarterScore(2),
		Q3Score:       h.parseQuarterScore(3),
		Q4Score:       h.parseQuarterScore(4),
		OvertimeScore: h.ParseOvertimeScore(),
		TotalScore:    h.parseTotalScore(),
	}
}
