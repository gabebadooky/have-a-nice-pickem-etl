package pickemstructs

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
	Finished      bool
}
