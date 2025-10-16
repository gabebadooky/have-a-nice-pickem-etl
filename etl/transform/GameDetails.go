package transform

type Details struct {
	gameID     string
	league     string
	week       uint8
	year       uint16
	espnCode   string
	cbsCode    string
	foxCode    string
	vegasCode  string
	awayTeamID string
	homeTeamID string
	date       string
	time       string
	tvCoverage string
	finished   string
}

type Team struct {

}


func parseAwayTeam(espnGameDetails map[string]any) map[string] any {
	var team map[string]any = espnGameDetails["header"]["competitions"][0]["competitors"][0]
	if (team["homeAway"] == "away") {
		return team
	} else {
		return espnGameDetails["header"]["competitions"][0]["competitors"][1]
	}
}


func parseHomeTeam(espnGameDetails map[string]any) map[string] any {
	var team map[string]any = espnGameDetails["header"]["competitions"][0]["competitors"][0]
	if (team["homeAway"] == "home") {
		return team
	} else {
		return espnGameDetails["header"]["competitions"][0]["competitors"][1]
	}
}


func InstantiateGameDetails(espnGameDetails map[string]any) Details {
	return Details{
		gameID: "",
		league: "",
		week: espnGameDetails["header"]["week"],
		year: espnGameDetails["header"]["season"]["year"],
		espnCode: espnGameDetails["header"]["id"],
		cbsCode: "",
		foxCode: "",
		vegasCode: "",
		awayTeamID: parseAwayTeam(espnGameDetails),
		homeTeamID: parseHomeTeam(espnGameDetails),
		date: "",
		time: "",
		tvCoverage: "",
		finished: espnGameDetails["header"]["competitions"][0]["status"]["type"]["completed"]
	}
}
