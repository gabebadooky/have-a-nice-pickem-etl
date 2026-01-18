package espnschedule

type ScoreboardEndpoint struct {
	Events []EventProperty `json:"events"`
}

type EventProperty struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Week         NumberProperty         `json:"week"`
	Competitions []CompetitionsProperty `json:"competitions"`
}

type NumberProperty struct {
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
