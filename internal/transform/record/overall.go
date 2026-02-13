// Package record provides overall record parsing functionality.
// It extracts overall win-loss-tie records from ESPN Team Summary API responses.
package record

import (
	"strconv"
	"strings"
)

func (o OverallRecord) parseOverallRecordElements() [3]uint {
	var overallRecord string = o.ESPN.Team.OverallRecord.RecordItems[0].Summary
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

func (o OverallRecord) parseOverallWins() uint {
	var overallWins uint = o.parseOverallRecordElements()[0]
	return overallWins
}

func (o OverallRecord) parseOverallLosses() uint {
	var overallLosses uint = o.parseOverallRecordElements()[1]
	return overallLosses
}

func (o OverallRecord) parseOverallTies() uint {
	var overallTies uint = o.parseOverallRecordElements()[2]
	return overallTies
}
