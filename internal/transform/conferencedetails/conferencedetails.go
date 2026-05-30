package conferencedetails

import (
	"have-a-nice-pickem-etl/internal/extract/team"
	"time"
)

type New struct {
	team.Team
}

type ConferenceDetails struct {
	ConferenceID    string    `gorm:"column:id"`
	Name            string    `gorm:"column:name"`
	Abbreviation    string    `gorm:"column:abbreviation"`
	PowerConference bool    `gorm:"column:power_conference"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (ConferenceDetails) TableName() string {
	return "conference"
}

func (c New) parseConferenceID() string {
	
}

func (c New) Instantiate() ConferenceDetails {
	return ConferenceDetails{
		ConferenceID: ,
		Name: ,
		Abbreviation: ,
		PowerConference: ,
	}
}
