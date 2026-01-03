package teamdetails

import (
	"have-a-nice-pickem-etl/etl/extract/team"
	"have-a-nice-pickem-etl/etl/transform/common"
)

type Instantiator interface {
	instantiate() TeamDetails
}

type NewTeamDetails struct {
	TeamExtract team.Team
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

func InstantiateTeamDetails(i Instantiator) TeamDetails {
	return i.instantiate()
}

func (t NewTeamDetails) parseConferenceID() string {
	var conferenceID string = t.TeamExtract.ESPN.Team.Groups.ID
	return conferenceID
}

func (t NewTeamDetails) parseTeamName() string {
	var teamName string = t.TeamExtract.ESPN.Team.Location
	return teamName
}

func (t NewTeamDetails) parseTeamMascot() string {
	var teamMascot string = t.TeamExtract.ESPN.Team.Name
	return teamMascot
}

func (t NewTeamDetails) parsePrimaryColor() string {
	var primaryColor string = t.TeamExtract.ESPN.Team.PrimaryColor
	return primaryColor
}

func (t NewTeamDetails) parseAlternateColor() string {
	var alternateColor string = t.TeamExtract.ESPN.Team.AlternateColor
	return alternateColor
}

func (t NewTeamDetails) parseRanking() uint {
	var ranking uint = t.TeamExtract.ESPN.Team.Ranking
	return ranking
}

func (t NewTeamDetails) instantiate() TeamDetails {
	var teamExtract team.Team = t.TeamExtract
	var teamID string = teamExtract.TeamID

	return TeamDetails{
		TeamID:         teamID,
		League:         teamExtract.League,
		ESPNCode:       common.ParseEspnTeamCode(teamExtract),
		CBSCode:        common.GetCbsTeamCode(teamID),
		FoxCode:        common.GetFoxTeamCode(teamID),
		VegasCode:      "",
		ConferenceID:   t.parseConferenceID(),
		Name:           t.parseTeamName(),
		Mascot:         t.parseTeamMascot(),
		PrimaryColor:   t.parsePrimaryColor(),
		AlternateColor: t.parseAlternateColor(),
		Ranking:        t.parseRanking(),
	}
}
