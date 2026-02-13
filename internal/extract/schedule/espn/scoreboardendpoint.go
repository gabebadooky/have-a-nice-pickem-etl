// Package espnschedule defines the data structures for ESPN Scoreboard API responses.
// These types represent the JSON structure returned by ESPN's scoreboard endpoint
// for schedule and event information.
package espnschedule

type ScoreboardEndpoint struct {
	Events []EventProperty `json:"events"`
}

type EventProperty struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Season       SeasonProperty         `json:"season"`
	Week         WeekProperty           `json:"week"`
	Competitions []CompetitionsProperty `json:"competitions"`
}

type SeasonProperty struct {
	Year uint `json:"year"`
	Type uint `json:"type"`
}

type WeekProperty struct {
	Number uint `json:"number"`
}

type CompetitionsProperty struct {
	Venue       VenueProperty         `json:"venue"`
	Competitors []CompetitorsProperty `json:"competitors"`
}

type CompetitorsProperty struct {
	ID string `json:"id"`
}

type VenueProperty struct {
	FullName string          `json:"fullName"`
	Address  AddressProperty `json:"address"`
}

type AddressProperty struct {
	City  string `json:"city"`
	State string `json:"state"`
}
