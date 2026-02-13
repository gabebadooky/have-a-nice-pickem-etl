package foxgame

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"
)

type awayTeam struct {
	hyperlink string
}

type homeTeam struct {
	hyperlink string
}

// Extracts team string Between after league path ("/nfl/", "/college-football/")
// and "-vs-" substring in a given Fox hyperlink
func (a awayTeam) scrapeTeamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(a.hyperlink)
	formattedGameCode = utils.StripBowlGamePrefixFromFoxGameCode(formattedGameCode)
	teamCode, _, _ := strings.Cut(formattedGameCode, "-vs-")
	_, teamCodeWithoutHyperlinkPrefix, exists := strings.Cut(teamCode, "l/")
	if exists {
		return teamCodeWithoutHyperlinkPrefix
	}
	return teamCode
}

// Extracts team string AFTER "-vs-" substring in a given Fox hyperlink
func (h homeTeam) scrapeTeamCode() string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(h.hyperlink)
	formattedGameCode = utils.StripBowlGamePrefixFromFoxGameCode(formattedGameCode)
	_, teamCode, _ := strings.Cut(formattedGameCode, "-vs-")
	return teamCode
}
