// Package foxgame provides helper functions for extracting team codes from Fox game URLs.
// These functions parse team identifiers from Fox Sports hyperlinks for matching
// games with team IDs.
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

func extractFormattedGameCode(gameHyperlink string) string {
	formattedGameCode := utils.StripDateAndBoxScoreIDFromFoxGameCode(gameHyperlink)
	formattedGameCode = utils.StripBowlGamePrefixFromFoxGameCode(formattedGameCode)
	return formattedGameCode
}

// scrapeTeamCode returns the Fox away team code from the game URL (before "-vs-").
func (a awayTeam) scrapeTeamCode() string {
	formattedGameCode := extractFormattedGameCode(a.hyperlink)
	teamCode, _, _ := strings.Cut(formattedGameCode, "-vs-")
	_, teamCodeWithoutHyperlinkPrefix, exists := strings.Cut(teamCode, "l/")
	if exists {
		return teamCodeWithoutHyperlinkPrefix
	}
	return teamCode
}

// scrapeTeamCode returns the Fox home team code from the game URL (after "-vs-").
func (h homeTeam) scrapeTeamCode() string {
	formattedGameCode := extractFormattedGameCode(h.hyperlink)
	_, teamCode, _ := strings.Cut(formattedGameCode, "-vs-")
	return teamCode
}
