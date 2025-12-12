package sharedtypes

type BettingOdds struct {
	GameID            string
	Source            string
	OverUnder         float32
	AwayMoneyline     int16
	HomeMoneyline     int16
	AwaySpread        float32
	HomeSpread        float32
	AwayWinPercentage string
	HomeWinPercentage string
}

type Boxscore struct {
	GameID        string
	TeamID        string
	Q1Score       uint8
	Q2Score       uint8
	Q3Score       uint8
	Q4Score       uint8
	OvertimeScore uint8
	TotalScore    uint8
}

type GameDetails struct {
	GameID        string
	League        string
	Week          int8
	Year          uint16
	ESPNCode      string
	CBSCode       string
	FoxCode       string
	VegasCode     string
	AwayTeamID    string
	HomeTeamID    string
	ZuluTimestamp string
	Broadcast     string
	Location      string
	Finished      bool
}

type Location struct {
	LocationID string
	Stadium    string
	City       string
	State      string
	Latitude   float64
	Longitude  float64
}

type Record struct {
	TeamID     string
	RecordType string
	Wins       uint8
	Losses     uint8
	Ties       uint8
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
	Ranking        uint8
}
