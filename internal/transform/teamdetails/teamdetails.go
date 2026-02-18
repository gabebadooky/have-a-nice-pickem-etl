// Package teamdetails provides team details transformation functionality that extracts
// and structures comprehensive team information including team codes from multiple sources,
// conference ID, name, mascot, colors, and ranking from ESPN team data.
package teamdetails

import (
	"have-a-nice-pickem-etl/internal/extract/team"
	"have-a-nice-pickem-etl/internal/transform/common"
	"have-a-nice-pickem-etl/internal/utils"
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

// parseConferenceID returns the conference/group ID from the team's ESPN data.
func (t New) parseConferenceID() string {
	var conferenceID string = t.ESPN.Team.Groups.ID
	return conferenceID
}

// parseTeamName returns the team location/name from the team's ESPN data.
func (t New) parseTeamName() string {
	var teamName string = t.ESPN.Team.Location
	return teamName
}

// parseTeamMascot returns the team mascot (name) from the team's ESPN data.
func (t New) parseTeamMascot() string {
	var teamMascot string = t.ESPN.Team.Name
	return teamMascot
}

// parsePrimaryColor returns the team primary color from the team's ESPN data.
func (t New) parsePrimaryColor() string {
	var primaryColor string = t.ESPN.Team.PrimaryColor
	return primaryColor
}

// parseAlternateColor returns the team alternate color from the team's ESPN data.
func (t New) parseAlternateColor() string {
	var alternateColor string = t.ESPN.Team.AlternateColor
	return alternateColor
}

// parseRanking returns the team ranking from the team's ESPN data.
func (t New) parseRanking() uint {
	var ranking uint = t.ESPN.Team.Ranking
	return ranking
}

// Instantiate builds a TeamDetails value from the extracted team data.
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
