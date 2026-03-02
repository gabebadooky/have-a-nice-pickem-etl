// Package record provides team record transformation functionality that extracts
// and structures win-loss-tie records from ESPN team data for both conference
// and overall records.
package record

import (
	"have-a-nice-pickem-etl/internal/extract/team"
)

type ConferenceRecord struct {
	team.Team
}

type OverallRecord struct {
	team.Team
}

type Record struct {
	TeamID     string
	RecordType string
	Wins       uint
	Losses     uint
	Ties       uint
}

// instantiate builds a conference record from the team's CBS page.
func (c ConferenceRecord) Instantiate() Record {
	return Record{
		TeamID:     c.TeamID,
		RecordType: "Conference",
		Wins:       c.parseWins(),
		Losses:     c.parseLosses(),
		Ties:       c.parseTies(),
	}
}

// instantiate builds an overall record from the team's ESPN data.
func (o OverallRecord) Instantiate() Record {
	return Record{
		TeamID:     o.TeamID,
		RecordType: "Overall",
		Wins:       o.parseWins(),
		Losses:     o.parseLosses(),
		Ties:       o.parseTies(),
	}
}
