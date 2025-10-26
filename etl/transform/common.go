package transform

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
