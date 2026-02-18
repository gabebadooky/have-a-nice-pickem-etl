// Package utils provides string manipulation and conversion utility functions.
// These functions handle formatting, parsing, and transforming string data
// used throughout the ETL pipeline.
package utils

import (
	"strconv"
	"strings"
)

// FormatStringID normalizes a string into an ID: lowercase, spaces to hyphens, and strips .')(&é.
func FormatStringID(id string) string {
	toLower := strings.ToLower(id)
	stripSpaces := strings.ReplaceAll(toLower, " ", "-")
	stripDuplicateDashes := strings.ReplaceAll(stripSpaces, "--", "-")
	stripPeriods := strings.ReplaceAll(stripDuplicateDashes, ".", "")
	stripSingleQuote := strings.ReplaceAll(stripPeriods, "'", "")
	stripLeftParan := strings.ReplaceAll(stripSingleQuote, "(", "")
	stripRightParan := strings.ReplaceAll(stripLeftParan, ")", "")
	stripAmpersand := strings.ReplaceAll(stripRightParan, "&", "")
	stripAccentedE := strings.ReplaceAll(stripAmpersand, "é", "e")

	formatString := stripAccentedE
	return formatString
}

// StripBowlGamePrefixFromFoxGameCode removes bowl/CFP prefixes from a Fox game code string.
func StripBowlGamePrefixFromFoxGameCode(foxGameCode string) string {
	bowlString := "bowl-"
	cfpFirstRoundGameString := "cfp-first-round-game-"
	cfpNationalChampionshipString := "cfp-national-championship-"

	if strings.Contains(foxGameCode, bowlString) {
		return foxGameCode[strings.Index(foxGameCode, bowlString)+len(bowlString):]
	}
	if strings.Contains(foxGameCode, cfpFirstRoundGameString) {
		return foxGameCode[strings.Index(foxGameCode, cfpFirstRoundGameString)+len(cfpFirstRoundGameString):]
	}
	if strings.Contains(foxGameCode, cfpNationalChampionshipString) {
		return foxGameCode[strings.Index(foxGameCode, cfpNationalChampionshipString)+len(cfpNationalChampionshipString):]
	}
	return foxGameCode
}

// StripDateAndBoxScoreIDFromFoxGameCode removes the date and box score suffix from a Fox game code.
func StripDateAndBoxScoreIDFromFoxGameCode(foxGameCode string) string {
	var formattedGameCode string
	var monthAbbr string

	seasonMonthAbbreviations := [12]string{"aug", "sep", "oct", "nov", "dec", "jan", "feb", "mar", "apr", "may", "jun", "jul"}

	for i := range 12 {
		if strings.Contains(foxGameCode, seasonMonthAbbreviations[i]) {
			monthAbbr = seasonMonthAbbreviations[i]
			break
		}
	}

	idx := strings.Index(foxGameCode, monthAbbr)
	formattedGameCode = foxGameCode[:idx-1]
	return formattedGameCode
}

// ConvertStringToFloat32 parses a string as float32, returning 0 on parse error.
func ConvertStringToFloat32(str string) float32 {
	f64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		f64 = 0.00
	}

	return float32(f64)
}

// ConvertStringToInt parses a string as int via Atoi, returning 0 on parse error.
func ConvertStringToInt(str string) int {
	convertedNumber, err := strconv.Atoi(str)
	if err != nil {
		convertedNumber = 0
	}

	return convertedNumber
}

// GetNumberOfSecondsFromDurationString parses a "M:SS" duration string and returns total seconds.
func GetNumberOfSecondsFromDurationString(durationString string) int {
	colonIndex := strings.Index(durationString, ":")
	minutes := durationString[:colonIndex]
	seconds := durationString[colonIndex:]
	totalSeconds := ConvertStringToInt(minutes) + ConvertStringToInt(seconds)
	return totalSeconds
}
