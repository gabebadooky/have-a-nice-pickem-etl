package teamdetails

import (
	"have-a-nice-pickem-etl/etl/extract/team"
	"have-a-nice-pickem-etl/etl/transform/common"
	"have-a-nice-pickem-etl/etl/utils"
)

/*
	type Instantiator interface {
		instantiate() TeamDetails
	}
*/
type New struct {
	team.Team
}

type TeamDetails struct {
	TeamID         string
	League         string
	ESPNCode       string
	CBSCode        string
	FoxCode        string
	VegasCode      string
	ConferenceID   string
	Name           string
	Mascot         string
	PrimaryColor   string
	AlternateColor string
	Ranking        uint
}

/*
	func InstantiateTeamDetails(i Instantiator) TeamDetails {
		return i.instantiate()
	}
*/

func (t New) parseConferenceID() string {
	var conferenceID string = t.ESPN.Team.Groups.ID
	return conferenceID
}

func (t New) parseTeamName() string {
	var teamName string = t.ESPN.Team.Location
	return teamName
}

func (t New) parseTeamMascot() string {
	var teamMascot string = t.ESPN.Team.Name
	return teamMascot
}

func (t New) parsePrimaryColor() string {
	var primaryColor string = t.ESPN.Team.PrimaryColor
	return primaryColor
}

func (t New) parseAlternateColor() string {
	var alternateColor string = t.ESPN.Team.AlternateColor
	return alternateColor
}

func (t New) parseRanking() uint {
	var ranking uint = t.ESPN.Team.Ranking
	return ranking
}

func (t New) Instantiate() TeamDetails {
	var teamID string = t.TeamID

	return TeamDetails{
		TeamID:         teamID,
		League:         t.League,
		ESPNCode:       common.ParseEspnTeamCode(t.Team),
		CBSCode:        utils.GetCbsTeamCode(teamID),
		FoxCode:        utils.GetFoxTeamCode(teamID),
		VegasCode:      "",
		ConferenceID:   t.parseConferenceID(),
		Name:           t.parseTeamName(),
		Mascot:         t.parseTeamMascot(),
		PrimaryColor:   t.parsePrimaryColor(),
		AlternateColor: t.parseAlternateColor(),
		Ranking:        t.parseRanking(),
	}
}
