package espn

type ScoreboardEndpoint struct {
	Events []EventProperty `json:"events"`
}

type EventProperty struct {
	ID string `json:"id"`
}
