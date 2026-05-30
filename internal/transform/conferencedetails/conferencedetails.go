package conferencedetails

import (
	"have-a-nice-pickem-etl/internal/extract/team"
	"have-a-nice-pickem-etl/internal/utils"
	"time"
)

type New struct {
	team.Team
}

type ConferenceDetails struct {
	ConferenceID    string    `gorm:"column:id"`
	Name            string    `gorm:"column:name"`
	Abbreviation    string    `gorm:"column:abbreviation"`
	PowerConference bool      `gorm:"column:power_conference"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (ConferenceDetails) TableName() string {
	return "conference"
}

func (c New) Instantiate() ConferenceDetails {
	var conferenceID string = c.Team.ESPN.Team.Groups.ID

	return ConferenceDetails{
		ConferenceID:    conferenceID,
		Name:            utils.ConferenceMapping[conferenceID].Name,
		Abbreviation:    utils.ConferenceMapping[conferenceID].Abbreviation,
		PowerConference: utils.ConferenceMapping[conferenceID].PowerConference,
	}
}
