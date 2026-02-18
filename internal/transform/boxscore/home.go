// Package boxscore provides home team boxscore parsing functionality.
// It extracts scoring data from ESPN Game Summary API responses for home teams
// including quarter-by-quarter scores, total scores, and overtime scores.
package boxscore

import (
	espngame "have-a-nice-pickem-etl/internal/extract/game/espn"
	"log"
	"strconv"
	"strings"
)

// parsePointsForGivenQuarter returns the home team's points for the period given competitor home/away and point strings.
func (h HomeBoxscore) parsePointsForGivenQuarter(team1homeAway string, team2homeAway string, team1points string, team2points string) uint {
	switch "HOME" {
	case strings.ToUpper(team1homeAway):
		points, err := strconv.Atoi(team1points)
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

// parseQuarterScore returns the home team's score for the given quarter (1-4) from the ESPN linescore.
func (h HomeBoxscore) parseQuarterScore(quarterNumber uint) uint {
	var linescoreArraryIndex uint = quarterNumber - 1

	if len(h.ESPN.Header.Competitions[0].Competitors[0].Linescores) == 0 || len(h.ESPN.Header.Competitions[0].Competitors[1].Linescores) == 0 {
		return 0
	} else {
		var team1HomeAway string = h.ESPN.Header.Competitions[0].Competitors[0].HomeAway
		var team2HomeAway string = h.ESPN.Header.Competitions[0].Competitors[1].HomeAway
		var team1Points string = h.ESPN.Header.Competitions[0].Competitors[0].Linescores[linescoreArraryIndex].DisplayValue
		var team2Points string = h.ESPN.Header.Competitions[0].Competitors[1].Linescores[linescoreArraryIndex].DisplayValue
		return h.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}

// parseTotalScore returns the home team's total score from the ESPN competition data.
func (h HomeBoxscore) parseTotalScore() uint {
	var team1HomeAway string = h.ESPN.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = h.ESPN.Header.Competitions[0].Competitors[1].HomeAway
	var team1Points string = h.ESPN.Header.Competitions[0].Competitors[0].Score
	var team2Points string = h.ESPN.Header.Competitions[0].Competitors[1].Score

	if team1Points == "" || team2Points == "" {
		return 0
	} else {
		return h.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}

// ParseOvertimeScore returns the home team's overtime score from the ESPN linescore (fifth period if present).
func (h HomeBoxscore) ParseOvertimeScore() uint {
	var team1HomeAway string = h.ESPN.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = h.ESPN.Header.Competitions[0].Competitors[1].HomeAway
	var team1PointsSlice []espngame.LinescoreProperty = h.ESPN.Header.Competitions[0].Competitors[0].Linescores
	var team2PointsSlice []espngame.LinescoreProperty = h.ESPN.Header.Competitions[0].Competitors[1].Linescores

	if len(team1PointsSlice) <= 4 {
		return uint(0)
	} else {
		var team1Points string = team1PointsSlice[4].DisplayValue
		var team2Points string = team2PointsSlice[4].DisplayValue
		return h.parsePointsForGivenQuarter(team1HomeAway, team2HomeAway, team1Points, team2Points)
	}
}
