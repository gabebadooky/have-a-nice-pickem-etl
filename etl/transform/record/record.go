package record

import (
	"have-a-nice-pickem-etl/etl/extract/team"
)

type Instantiator interface {
	instantiate() Record
}

type ConferenceRecord struct {
	TeamExtract team.Team
}

type OverallRecord struct {
	TeamExtract team.Team
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
		TeamID:     c.TeamExtract.TeamID,
		RecordType: "Conference",
		Wins:       0,
		Losses:     0,
		Ties:       0,
	}
}

func (o OverallRecord) instantiate() Record {
	return Record{
		TeamID:     o.TeamExtract.TeamID,
		RecordType: "Overall",
		Wins:       0,
		Losses:     0,
		Ties:       0,
	}
}
