// Package boxscore provides away team boxscore parsing functionality.
// It extracts scoring data from ESPN Game Summary API responses for away teams
// including quarter-by-quarter scores, total scores, and overtime scores.
package boxscore

import (
	espngame "have-a-nice-pickem-etl/internal/extract/game/espn"
	"log"
	"strconv"
	"strings"
)

// Extract Points for given boxscore period
func (a AwayBoxscore) parsePointsForGivenQuarter(team1homeAway string, team2homeAway string, team1points string, team2points string) uint {
	switch "AWAY" {
	case strings.ToUpper(team1homeAway):
		points, err := strconv.Atoi(team2points)
		if err != nil {
			points = 0
		}
		return uint(points)

	case strings.ToUpper(team2homeAway):
		points, err := strconv.Atoi(team2points)
		if err != nil {
			points = 0
		}
		return uint(points)

	default:
		log.Printf("Invalid homeAway value supplied: %s, %s\n", team1homeAway, team2homeAway)
		return uint(0)
	}
}

// Parses "Linescore" field for a given competitor and quarter from the ESPN Game Summary API
func (a AwayBoxscore) parseQuarterScore(quarterNumber uint) uint {
	var linescoreArraryIndex uint = quarterNumber - 1

	if len(a.ESPN.Header.Competitions[0].Competitors[0].Linescores) == 0 || len(a.ESPN.Header.Competitions[0].Competitors[1].Linescores) == 0 {
		return 0
	} else {
		var team1HomeAway string = a.ESPN.Header.Competitions[0].Competitors[0].HomeAway
		var team2HomeAway string = a.ESPN.Header.Competitions[0].Competitors[1].HomeAway
		var team1Points string = a.ESPN.Header.Competitions[0].Competitors[0].Linescores[linescoreArraryIndex].DisplayValue
		var team2Points string = a.ESPN.Header.Competitions[0].Competitors[1].Linescores[linescoreArraryIndex].DisplayValue
		return a.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}

// Parses "score" field for a given competitor from the ESPN Game Summary API
func (a AwayBoxscore) parseTotalScore() uint {
	var team1HomeAway string = a.ESPN.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = a.ESPN.Header.Competitions[0].Competitors[1].HomeAway
	var team1Points string = a.ESPN.Header.Competitions[0].Competitors[0].Score
	var team2Points string = a.ESPN.Header.Competitions[0].Competitors[1].Score

	if team1Points == "" || team2Points == "" {
		return 0
	} else {
		return a.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}

// Parses "Linescore" field for a given competitor and quarter from the ESPN Game Summary API
func (a AwayBoxscore) ParseOvertimeScore() uint {
	var team1HomeAway string = a.ESPN.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = a.ESPN.Header.Competitions[0].Competitors[1].HomeAway
	var team1PointsSlice []espngame.LinescoreProperty = a.ESPN.Header.Competitions[0].Competitors[0].Linescores
	var team2PointsSlice []espngame.LinescoreProperty = a.ESPN.Header.Competitions[0].Competitors[1].Linescores

	if len(team1PointsSlice) <= 4 {
		return uint(0)
	} else {
		var team1Points string = team1PointsSlice[4].DisplayValue
		var team2Points string = team2PointsSlice[4].DisplayValue
		return a.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}
