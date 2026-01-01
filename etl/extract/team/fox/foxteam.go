package foxteam

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"
)

type FoxTeam interface {
	teamCode() string
}

type FoxAwayTeam struct {
	FoxGameCode string
}

type FoxHomeTeam struct {
	FoxGameCode string
}

func ExtractFoxTeamCode(t FoxTeam) string {
	return t.teamCode()
}

// Extracts team string BEFORE "-vs-" substring in a given Fox Game Code
func (t FoxAwayTeam) teamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(t.FoxGameCode)
	// teamCode := strings.Split(formattedGameCode, "-vs-")[0]
	teamCode, _, _ := strings.Cut(formattedGameCode, "-vs-")
	return teamCode
}

// Extracts team string AFTER "-vs-" substring in a given Fox Game Code
func (t FoxHomeTeam) teamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(t.FoxGameCode)
	// teamCode := strings.Split(formattedGameCode, "-vs-")[1]
	_, teamCode, _ := strings.Cut(formattedGameCode, "-vs-")
	return teamCode
}
