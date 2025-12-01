package utils

import (
	"strconv"
	"strings"
)

// Replaces spaces with - and strips the following characters from a given string: .')(&é
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

// Strips Date and Box Score from Fox Game Code
func StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode string) string {
	var seasonMonthAbbreviations [12]string = [12]string{"aug", "sep", "oct", "nov", "dec", "jan", "feb", "mar", "apr", "may", "jun", "jul"}
	var formattedGameCode string
	var monthAbbr string

	for i := range 12 {
		if strings.Contains(foxGameCode, seasonMonthAbbreviations[i]) {
			monthAbbr = seasonMonthAbbreviations[i]
			break
		}
	}

	var idx int = strings.Index(foxGameCode, monthAbbr)
	formattedGameCode = foxGameCode[:idx-1]

	return formattedGameCode
}

func ConvertStringToFloat32(bettingOddString string) float32 {
	f64, err := strconv.ParseFloat(bettingOddString, 32)
	if err != nil {
		f64 = 0.00
	}

	return float32(f64)
}

func ConvertStringToInt16(bettingOddString string) int16 {
	bettingOddInt, err := strconv.Atoi(bettingOddString)
	if err != nil {
		bettingOddInt = 0.00
	}

	return int16(bettingOddInt)
}
