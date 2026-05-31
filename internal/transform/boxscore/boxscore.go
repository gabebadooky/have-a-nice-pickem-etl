// Package boxscore provides boxscore transformation functionality that extracts
// and structures scoring data (quarter scores, total scores, overtime scores)
// from ESPN game data for both away and home teams.
package boxscore

import (
	"have-a-nice-pickem-etl/internal/extract/game"
	"have-a-nice-pickem-etl/internal/transform/common"
	"time"
)

type AwayBoxscore struct {
	game.Game
}

type HomeBoxscore struct {
	game.Game
}

type Boxscore struct {
	GameID        string    `gorm:"column:game"`
	TeamID        string    `gorm:"column:team"`
	Q1Score       uint      `gorm:"column:quarter1"`
	Q2Score       uint      `gorm:"column:quarter2"`
	Q3Score       uint      `gorm:"column:quarter3"`
	Q4Score       uint      `gorm:"column:quarter4"`
	OvertimeScore uint      `gorm:"column:overtime"`
	TotalScore    uint      `gorm:"column:total"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (Boxscore) TableName() string {
	return "box_score"
}

// instantiate builds the away team boxscore from the game's ESPN linescore data.
func (a AwayBoxscore) Instantiate() Boxscore {
	return Boxscore{
		GameID:        a.GameID,
		TeamID:        common.ParseAwayTeamID(a.Game),
		Q1Score:       a.parseQuarterScore(1),
		Q2Score:       a.parseQuarterScore(2),
		Q3Score:       a.parseQuarterScore(3),
		Q4Score:       a.parseQuarterScore(4),
		OvertimeScore: a.parseOvertimeScore(),
		TotalScore:    a.parseTotalScore(),
	}
}

// instantiate builds the home team boxscore from the game's ESPN linescore data.
func (h HomeBoxscore) Instantiate() Boxscore {
	return Boxscore{
		GameID:        h.GameID,
		TeamID:        common.ParseHomeTeamID(h.Game),
		Q1Score:       h.parseQuarterScore(1),
		Q2Score:       h.parseQuarterScore(2),
		Q3Score:       h.parseQuarterScore(3),
		Q4Score:       h.parseQuarterScore(4),
		OvertimeScore: h.parseOvertimeScore(),
		TotalScore:    h.parseTotalScore(),
	}
}
