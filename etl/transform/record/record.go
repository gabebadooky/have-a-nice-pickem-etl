package record

import (
	"have-a-nice-pickem-etl/etl/extract/team"
)

type Instantiator interface {
	instantiate() Record
}

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

func InstantiateRecord(i Instantiator) Record {
	return i.instantiate()
}

func (c ConferenceRecord) instantiate() Record {
	return Record{
		TeamID:     c.TeamID,
		RecordType: "Conference",
		Wins:       c.parseWins(),
		Losses:     c.parseLosses(),
		Ties:       c.parseTies(),
	}
}

func (o OverallRecord) instantiate() Record {
	return Record{
		TeamID:     o.TeamID,
		RecordType: "Overall",
		Wins:       o.parseOverallWins(),
		Losses:     o.parseOverallLosses(),
		Ties:       o.parseOverallTies(),
	}
}
