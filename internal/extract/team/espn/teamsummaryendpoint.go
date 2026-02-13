// Package espnteam defines the data structures for ESPN Team Summary API responses.
// These types represent the JSON structure returned by ESPN's team summary endpoint
// including team details, records, logos, and group information.
package espnteam

type TeamSummaryEndpoint struct {
	Team TeamProperty `json:"team"`
}

type TeamProperty struct {
	ID             string                `json:"slug"`
	Code           string                `json:"id"`
	Location       string                `json:"location"`
	Name           string                `json:"name"`
	PrimaryColor   string                `json:"color"`
	AlternateColor string                `json:"alternateColor"`
	Logos          []LogoProperty        `json:"logos"`
	OverallRecord  OverallRecordProperty `json:"record"`
	Groups         GroupsProperty        `json:"groups"`
	Ranking        uint                  `json:"rank"`
}

type GroupsProperty struct {
	ID string `json:"id"`
}

type LogoProperty struct {
	HREF string `json:"href"`
}

type OverallRecordProperty struct {
	RecordItems []RecordItemProperty `json:"items"`
}

type RecordItemProperty struct {
	Summary string `json:"summary"`
}
