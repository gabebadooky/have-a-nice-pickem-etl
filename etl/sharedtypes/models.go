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
	Q1Score       uint
	Q2Score       uint
	Q3Score       uint
	Q4Score       uint
	OvertimeScore uint
	TotalScore    uint
}

type GameDetails struct {
	GameID        string
	League        string
	Week          int8
	Year          uint
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
	Wins       uint
	Losses     uint
	Ties       uint
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
