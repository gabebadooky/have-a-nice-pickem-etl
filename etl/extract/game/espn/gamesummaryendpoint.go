package espngame

type GameSummaryEndpoint struct {
	GameInfo   GameInfoProperty     `json:"gameInfo"`
	Header     HeaderProperty       `json:"header"`
	Pickcenter []PickcenterProperty `json:"pickcenter"`
	Predictor  PredictorProperty    `json:"predictor"`
}

type HeaderProperty struct {
	Week         int8                   `json:"week"`
	Season       SeasonProperty         `json:"season"`
	ESPNGameCode string                 `json:"id"`
	Competitions []CompetitionsProperty `json:"competitions"`
	League       LeagueProperty         `json:"league"`
}

type SeasonProperty struct {
	Year uint `json:"year"`
}

type CompetitionsProperty struct {
	Competitors []CompetitorsProperty `json:"competitors"`
	Date        string                `json:"date"`
	Broadcasts  []BroadcastsProperty  `json:"broadcasts"`
	Status      StatusProperty        `json:"status"`
}

type CompetitorsProperty struct {
	HomeAway   string                 `json:"homeAway"`
	Linescores []LinescoreProperty    `json:"linescores"`
	Score      string                 `json:"score"`
	Team       CompetitorTeamProperty `json:"team"`
}

type LinescoreProperty struct {
	DisplayValue string `json:"displayValue"`
}

type CompetitorTeamProperty struct {
	DisplayName string `json:"displayName"`
}

type BroadcastsProperty struct {
	Media MediaProperty `json:"media"`
}

type MediaProperty struct {
	ShortName string `json:"shortName"`
}

type StatusProperty struct {
	Type TypeProperty `json:"type"`
}

type TypeProperty struct {
	Completed bool `json:"completed"`
}

type LeagueProperty struct {
	Abbreviation string `json:"abbreviation"`
}

type GameInfoProperty struct {
	Venue VenueProperty `json:"venue"`
}

type VenueProperty struct {
	FullName string          `json:"fullName"`
	Address  AddressProperty `json:"address"`
}

type AddressProperty struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipCode"`
	Country string `json:"country"`
}

type PickcenterProperty struct {
	OverUnder    float32          `json:"overUnder"`
	Spread       float32          `json:"spread"`
	AwayTeamOdds TeamOddsProperty `json:"awayTeamOdds"`
	HomeTeamOdds TeamOddsProperty `json:"homeTeamOdds"`
}

type TeamOddsProperty struct {
	Moneyline int16 `json:"moneyline"`
}

type PredictorProperty struct {
	AwayTeam PredictorTeamProperty `json:"awayTeam"`
	HomeTeam PredictorTeamProperty `json:"homeTeam"`
}

type PredictorTeamProperty struct {
	GameProjection string `json:"gameProjection"`
}
