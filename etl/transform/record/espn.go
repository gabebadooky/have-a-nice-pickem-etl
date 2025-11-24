package record

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"strconv"
	"strings"
)

func parseOverallRecordElements(espnTeamDetails pickemstructs.TeamSummaryResponse) [3]uint8 {
	var overallRecord string = espnTeamDetails.Team.OverallRecord.RecordItems[0].Summary
	var recordElements []string = strings.Split(overallRecord, "-")

	if len(recordElements) == 2 {
		recordElements = append(recordElements, "0")
	}

	wins, err := strconv.ParseUint(recordElements[0], 10, 8)
	if err != nil {
		wins = 0
	}

	losses, err := strconv.ParseUint(recordElements[1], 10, 8)
	if err != nil {
		losses = 0
	}

	ties, err := strconv.ParseUint(recordElements[2], 10, 8)
	if err != nil {
		ties = 0
	}

	return [3]uint8{uint8(wins), uint8(losses), uint8(ties)}
}

func ParseOverallWins(espnTeamDetails pickemstructs.TeamSummaryResponse) uint8 {
	var overallWins uint8 = parseOverallRecordElements(espnTeamDetails)[0]
	return overallWins
}

func ParseOverallLosses(espnTeamDetails pickemstructs.TeamSummaryResponse) uint8 {
	var overallLosses uint8 = parseOverallRecordElements(espnTeamDetails)[1]
	return overallLosses
}

func ParseOverallTies(espnTeamDetails pickemstructs.TeamSummaryResponse) uint8 {
	var overallTies uint8 = parseOverallRecordElements(espnTeamDetails)[2]
	return overallTies
}
