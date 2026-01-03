package boxscore

import (
	"have-a-nice-pickem-etl/etl/extract/game"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type Instantiator interface {
	instantiate() Boxscore
}

type AwayBoxscore struct {
	GameExtract game.Game
}

type HomeBoxscore struct {
	GameExtract game.Game
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

func InstantiateBoxscore(i Instantiator) Boxscore {
	return i.instantiate()
}

func (a AwayBoxscore) instantiate() Boxscore {
	var gameExtract game.Game = a.GameExtract

	return Boxscore{
		GameID:        gameExtract.GameID,
		TeamID:        common.ParseAwayTeamID(gameExtract),
		Q1Score:       a.parseQuarterScore(1),
		Q2Score:       a.parseQuarterScore(2),
		Q3Score:       a.parseQuarterScore(3),
		Q4Score:       a.parseQuarterScore(4),
		OvertimeScore: a.ParseOvertimeScore(),
		TotalScore:    a.parseTotalScore(),
	}
}

func (h HomeBoxscore) instantiate() Boxscore {
	var gameExtract game.Game = h.GameExtract

	return Boxscore{
		GameID:        gameExtract.GameID,
		TeamID:        common.ParseAwayTeamID(gameExtract),
		Q1Score:       h.parseQuarterScore(1),
		Q2Score:       h.parseQuarterScore(2),
		Q3Score:       h.parseQuarterScore(3),
		Q4Score:       h.parseQuarterScore(4),
		OvertimeScore: h.ParseOvertimeScore(),
		TotalScore:    h.parseTotalScore(),
	}
}
