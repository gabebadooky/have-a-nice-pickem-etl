// Package record provides overall record parsing functionality.
// It extracts overall win-loss-tie records from ESPN Team Summary API responses.
package record

import (
	espnteam "have-a-nice-pickem-etl/internal/extract/team/espn"
	"strconv"
	"strings"
)

// parseOverallRecordElements parses the ESPN overall record summary into wins, losses, and ties.
func (o OverallRecord) parseOverallRecordElements() [3]uint {
	var recordItems []espnteam.RecordItemProperty = o.ESPN.Team.OverallRecord.RecordItems
	if len(recordItems) == 0 {
		return [3]uint{0, 0, 0}
	}

	var overallRecord string = recordItems[0].Summary
	recordElements := strings.Split(overallRecord, "-")

	if len(recordElements) == 2 {
		recordElements = append(recordElements, "0")
	}

	wins, err := strconv.Atoi(recordElements[0])
	if err != nil {
		wins = 0
	}

	losses, err := strconv.Atoi(recordElements[1])
	if err != nil {
		losses = 0
	}

	ties, err := strconv.Atoi(recordElements[2])
	if err != nil {
		ties = 0
	}

	return [3]uint{uint(wins), uint(losses), uint(ties)}
}

// parseOverallWins returns overall wins from the team's ESPN record.
func (o OverallRecord) parseWins() uint {
	overallWins := o.parseOverallRecordElements()[0]
	return overallWins
}

// parseOverallLosses returns overall losses from the team's ESPN record.
func (o OverallRecord) parseLosses() uint {
	overallLosses := o.parseOverallRecordElements()[1]
	return overallLosses
}

// parseOverallTies returns overall ties from the team's ESPN record.
func (o OverallRecord) parseTies() uint {
	overallTies := o.parseOverallRecordElements()[2]
	return overallTies
}
