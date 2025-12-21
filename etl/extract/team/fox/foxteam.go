package fox

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"
)

type FoxTeam interface {
	ExtractFoxTeamCode() string
}

type FoxAwayTeam struct {
	FoxGameCode string
}

type FoxHomeTeam struct {
	FoxGameCode string
}

// Extracts Fox team code of 'Away' team from a given Fox Game Code
func (a FoxAwayTeam) ExtractFoxTeamCode() string {
	var formattedGameCode string = utils.StripDateAndBoxScoreIDFromFoxGameCode(a.FoxGameCode)
	var teamCode string = strings.Split(formattedGameCode, "-vs-")[1]
	return teamCode
}

// Extracts Fox team code of 'Home' or 'Away' team from a given Fox Game Code
func (a FoxHomeTeam) ExtractFoxTeamCode() string {
	var formattedGameCode string = utils.StripDateAndBoxScoreIDFromFoxGameCode(a.FoxGameCode)
	var teamCode string = strings.Split(formattedGameCode, "-vs-")[0]
	return teamCode
}
