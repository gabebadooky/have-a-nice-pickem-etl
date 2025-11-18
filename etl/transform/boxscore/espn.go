package boxscore

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"log"
	"strconv"
	"strings"
)

// Extract Points for given boxscore period
func ParsePointsForGivenCompetitorAndPeriod(homeAway string, team1homeAway string, team2homeAway string, team1points string, team2points string) uint8 {
	switch strings.ToUpper(homeAway) {
	case team1homeAway:
		points64, err := strconv.ParseInt(team1points, 10, 8)
		if err != nil {
			log.Printf("Error occurred converting string property to uint8 type: %s", err)
			points64 = 0
		}
		return uint8(points64)

	case team2homeAway:
		points64, err := strconv.ParseInt(team2points, 10, 8)
		if err != nil {
			log.Printf("Error occurred converting string property to uint8 type: %s", err)
			points64 = 0
		}
		return uint8(points64)

	default:
		log.Printf("Invalid homeAway value supplied: %s", homeAway)
		return uint8(0)
	}
}

// Parses "Linescore" field for a given competitor and quarter from the ESPN Game Summary API
func ParseQuarterScore(espnGameDetails pickemstructs.ESPNGameDetailsResponse, homeAway string, quarterNumber uint8) uint8 {
	var linescoreArraryIndex uint8 = quarterNumber - 1
	var team1HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[1].HomeAway
	var team1Points string = espnGameDetails.Header.Competitions[0].Competitors[0].Linescores[linescoreArraryIndex].DisplayValue
	var team2Points string = espnGameDetails.Header.Competitions[0].Competitors[1].Linescores[linescoreArraryIndex].DisplayValue

	return ParsePointsForGivenCompetitorAndPeriod(homeAway, team1HomeAway, team2HomeAway, team1Points, team2Points)

}

// Parses "score" field for a given competitor from the ESPN Game Summary API
func ParseTotalScore(espnGameDetails pickemstructs.ESPNGameDetailsResponse, homeAway string) uint8 {
	var team1HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[1].HomeAway
	var team1Points string = espnGameDetails.Header.Competitions[0].Competitors[0].Score
	var team2Points string = espnGameDetails.Header.Competitions[0].Competitors[1].Score

	return ParsePointsForGivenCompetitorAndPeriod(homeAway, team1HomeAway, team2HomeAway, team1Points, team2Points)

}

// Parses "Linescore" field for a given competitor and quarter from the ESPN Game Summary API
func ParseOvertimeScore(espnGameDetails pickemstructs.ESPNGameDetailsResponse, homeAway string) uint8 {
	var team1HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	var team2HomeAway string = espnGameDetails.Header.Competitions[0].Competitors[1].HomeAway
	var team1PointsSlice []pickemstructs.Linescore = espnGameDetails.Header.Competitions[0].Competitors[0].Linescores
	var team2PointsSlice []pickemstructs.Linescore = espnGameDetails.Header.Competitions[0].Competitors[1].Linescores
	var team1Points string = team1PointsSlice[4].DisplayValue
	var team2Points string = team2PointsSlice[4].DisplayValue

	if len(team1PointsSlice) == 4 {
		return uint8(0)
	} else {
		return ParsePointsForGivenCompetitorAndPeriod(homeAway, team1HomeAway, team2HomeAway, team1Points, team2Points)

	}

}
