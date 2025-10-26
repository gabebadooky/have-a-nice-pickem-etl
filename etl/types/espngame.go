package types

type ESPNGameDetailsResponse struct {
	Header Header `json:"header"`
}

type Header struct {
	Week         int8           `json:"week"`
	Season       Season         `json:"season"`
	ESPNGameCode string         `json:"id"`
	Competitions []Competitions `json:"competitions"`
}

type Season struct {
	Year uint16 `json:"year"`
}

type Competitions struct {
	Competitors []Competitors `json:"competitors"`
	Date        string        `json:"date"`
	Broadcasts  []Broadcasts  `json:"broadcasts"`
	Status      Status        `json:"status"`
}

type Competitors struct {
	HomeAway string `json:"homeAway"`
	Team     Team   `json:"team"`
}

type Team struct {
	DisplayName string `json:"displayName"`
}

type Broadcasts struct {
	Media Media `json:"media"`
}

type Media struct {
	ShortName string `json:"shortName"`
}

type Status struct {
	Type Type `json:"type"`
}

type Type struct {
	Completed bool `json:"completed"`
}
