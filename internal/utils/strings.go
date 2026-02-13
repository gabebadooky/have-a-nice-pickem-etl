// Package utils provides string manipulation and conversion utility functions.
// These functions handle formatting, parsing, and transforming string data
// used throughout the ETL pipeline.
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

// Strips Date and Box Score from Fox Game Code
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

func ConvertStringToFloat32(str string) float32 {
	f64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		f64 = 0.00
	}

	return float32(f64)
}

func ConvertStringToInt(str string) int {
	convertedNumber, err := strconv.Atoi(str)
	if err != nil {
		convertedNumber = 0
	}

	return convertedNumber
}

func GetNumberOfSecondsFromDurationString(durationString string) int {
	colonIndex := strings.Index(durationString, ":")
	minutes := durationString[:colonIndex]
	seconds := durationString[colonIndex:]
	totalSeconds := ConvertStringToInt(minutes) + ConvertStringToInt(seconds)
	return totalSeconds
}
