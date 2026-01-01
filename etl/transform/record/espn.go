package record

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"strconv"
	"strings"
)

func parseOverallRecordElements(espnTeamDetails pickemstructs.TeamSummaryResponse) [3]uint {
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

	return [3]uint{uint(wins), uint(losses), uint(ties)}
}

func ParseOverallWins(espnTeamDetails pickemstructs.TeamSummaryResponse) uint {
	var overallWins uint = parseOverallRecordElements(espnTeamDetails)[0]
	return overallWins
}

func ParseOverallLosses(espnTeamDetails pickemstructs.TeamSummaryResponse) uint {
	var overallLosses uint = parseOverallRecordElements(espnTeamDetails)[1]
	return overallLosses
}

func ParseOverallTies(espnTeamDetails pickemstructs.TeamSummaryResponse) uint {
	var overallTies uint = parseOverallRecordElements(espnTeamDetails)[2]
	return overallTies
}
