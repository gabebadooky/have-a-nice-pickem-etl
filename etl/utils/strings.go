package utils

import "strings"

func FormatStringID(id string) string {
	var toLower string = strings.ToLower(id)
	var stripSpaces string = strings.ReplaceAll(toLower, " ", "-")
	var stripDuplicateDashes string = strings.ReplaceAll(stripSpaces, "--", "-")
	var stripPeriods string = strings.ReplaceAll(stripDuplicateDashes, ".", "")
	var stripSingleQuote string = strings.ReplaceAll(stripPeriods, "'", "")
	var stripLeftParan string = strings.ReplaceAll(stripSingleQuote, "(", "")
	var stripRightParan string = strings.ReplaceAll(stripLeftParan, ")", "")
	var stripAmpersand string = strings.ReplaceAll(stripRightParan, "&", "")
	var stripAccentedE string = strings.ReplaceAll(stripAmpersand, "é", "e")

	var formatString string = stripAccentedE
	return formatString
}

func StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode string) string {
	var seasonMonthAbbreviations [12]string = [12]string{"aug", "sep", "oct", "nov", "dec", "jan", "feb", "mar", "apr", "may", "jun", "jul"}
	var formattedGameCode string
	var monthAbbr string

	for i := 0; i < 12; i++ {
		if strings.Contains(foxGameCode, monthAbbr) {
			monthAbbr = seasonMonthAbbreviations[i]
			break
		}
	}

	var idx int = strings.Index(foxGameCode, monthAbbr)
	formattedGameCode = foxGameCode[:idx-1]

	return formattedGameCode
}
