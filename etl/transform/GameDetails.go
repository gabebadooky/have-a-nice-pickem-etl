package transform

type Details struct {
	gameID     string
	league     string
	week       int8
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

func parseWeek(espnGameDetails map[string]interface{}) int8 {
	week := espnGameDetails["header"].(map[string]interface{})["week"]
	if value, ok := week.(int8); ok {
		return value
	} else {
		return -1
	}
}

func parseAwayTeam(espnGameDetails map[string]any) map[string]interface{} {
	var team map[string]any = espnGameDetails["header"]["competitions"][0]["competitors"][0]
	if team["homeAway"] == "away" {
		return team
	} else {
		return espnGameDetails["header"]["competitions"][0]["competitors"][1]
	}
}

func parseHomeTeam(espnGameDetails map[string]interface{}) map[string]interface{} {
	var team map[string]any = espnGameDetails["header"]["competitions"][0]["competitors"][0]
	if team["homeAway"] == "home" {
		return team
	} else {
		return espnGameDetails["header"]["competitions"][0]["competitors"][1]
	}
}

func InstantiateGameDetails(espnGameDetails map[string]interface{}) Details {
	var details Details

	details.gameID = ""
	details.league = ""
	details.week = int8(parseWeek(espnGameDetails))
	details.year = espnGameDetails["header"]["season"]["year"]
	details.espnCode = espnGameDetails["header"]["id"]
	details.cbsCode = ""
	details.foxCode = ""
	details.vegasCode = ""
	details.awayTeamID = parseAwayTeam(espnGameDetails)
	details.homeTeamID = parseHomeTeam(espnGameDetails)
	details.date = ""
	details.time = ""
	details.tvCoverage = ""
	details.finished = espnGameDetails["header"]["competitions"][0]["status"]["type"]["completed"]

	return details
}
