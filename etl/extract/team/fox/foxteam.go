package foxteam

import (
	"have-a-nice-pickem-etl/etl/utils"
	"strings"
)

type FoxTeam interface {
	teamCode() string
}

type FoxAwayTeam struct {
	FoxGameHyperlink string
}

type FoxHomeTeam struct {
	FoxGameHyperlink string
}

func ExtractFoxTeamCode(t FoxTeam) string {
	return t.teamCode()
}

// Extracts team string BEFORE "-vs-" substring in a given Fox Game Code
func (t FoxAwayTeam) teamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(t.FoxGameHyperlink)
	formattedGameCode = utils.StripBowlGamePrefixFromFoxGameCode(formattedGameCode)
	teamCode, _, _ := strings.Cut(formattedGameCode, "-vs-")
	_, teamCodeWithoutHyperlinkPrefix, exists := strings.Cut(teamCode, "l/")
	if exists {
		return teamCodeWithoutHyperlinkPrefix
	}
	return teamCode
}

// Extracts team string AFTER "-vs-" substring in a given Fox Game Code
func (t FoxHomeTeam) teamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(t.FoxGameHyperlink)
	formattedGameCode = utils.StripBowlGamePrefixFromFoxGameCode(formattedGameCode)
	_, teamCode, _ := strings.Cut(formattedGameCode, "-vs-")
	return teamCode
}
