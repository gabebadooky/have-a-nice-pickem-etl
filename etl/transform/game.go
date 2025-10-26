package transform

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/types"
)

type Details struct {
	gameID        string
	league        string
	week          int8
	year          uint16
	espnCode      string
	cbsCode       string
	foxCode       string
	vegasCode     string
	awayTeamID    string
	homeTeamID    string
	zuluTimestamp string
	broadcast     string
	finished      bool
}

/*
type ESPNGameDetailsResponse struct {
	Header Header `json:"header"`
}

type Header struct {
	Week         int8           `json:"week"`
	Season       Season         `json:"season"`
	ESPNGameCode string         `json:"id"`
	Competitions []Competitions `[]json:"competitions"`
}

type Season struct {
	Year uint16 `json:"year"`
}

type Competitions struct {
	Competitors []Competitors `[]json:"competitors"`
	Date        string        `json:"date"`
	Broadcasts  []Media       `[]json:"broadcasts"`
	Status      Status        `json:"status"`
}

type Competitors struct {
	HomeAway string `json:"homeAway"`
	TeamID   string `json:"slug"`
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
	Completed string `json:"completed"`
}


func parseGameID(espnGameDetails map[string]any) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	return gameID
}

func parseWeek(espnGameDetails ESPNGameDetailsResponse) int8 {
	week := espnGameDetails["header"].(map[string]any)["week"].(int8)
	if err {
		print(err)
		return -1
	}
	return week

}

func parseYear(espnGameDetails map[string]any) uint16 {
	year, err := espnGameDetails["header"].(map[string]any)["season"].(map[string]any)["year"].(uint16)
	if err {
		print(err)
		return uint16(time.Now().Year())
	}
	return year

}

func parseEspnGameCode(espnGameDetails map[string]any) string {
	gameCode, err := espnGameDetails["header"].(map[string]any)["id"].(string)
	if err {
		print(err)
		return "espn-game-code"
	}
	return gameCode

}

func parseTeamID(homeAway string, espnGameDetails map[string]any) string {
	competitorHomeAway, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[0].(map[string]any)["homeAway"].(string)
	if err {
		print(err)
		return "team-id"
	}

	if competitorHomeAway == homeAway {
		teamID, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[0].(map[string]any)["slug"].(string)
		if err {
			print(err)
			return "team-id"
		}
		return teamID

	} else {
		teamID, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[1].(map[string]any)["slug"].(string)
		if err {
			print(err)
			return "team-id"
		}
		return teamID

	}

}

func parseGameZuluTimestamp(espnGameDetails map[string]any) string {
	gameDate, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["date"].(string)
	if err {
		print(err)
		return ""
	}
	return gameDate
}

func parseBroadcast(espnGameDetails map[string]any) string {
	broadcast, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["broadcasts"].([]any)[0].(map[string]any)["media"].(map[string]any)["shortName"].(string)
	if err {
		print(err)
		return ""
	}
	return broadcast
}

func parseGameStatus(espnGameDetails map[string]any) string {
	gameStatus, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["status"].(map[string]any)["type"].(map[string]any)["completed"].(string)
	if err {
		print(err)
		return "Not Started"
	}
	return gameStatus

}
*/

func parseGameID(espnGameDetails types.ESPNGameDetailsResponse) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	var formattedGameID string = FormatStringID(gameID)
	return formattedGameID
}

func parseWeek(espnGameDetails types.ESPNGameDetailsResponse) int8 {
	var week int8 = espnGameDetails.Header.Week
	return week

}

func parseYear(espnGameDetails types.ESPNGameDetailsResponse) uint16 {
	var year uint16 = espnGameDetails.Header.Season.Year
	return year

}

func parseEspnGameCode(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameCode string = espnGameDetails.Header.ESPNGameCode
	return gameCode

}

func parseTeamID(homeAway string, espnGameDetails types.ESPNGameDetailsResponse) string {
	var competitorHomeAway string = espnGameDetails.Header.Competitions[0].Competitors[0].HomeAway
	if homeAway == competitorHomeAway {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[0].Team.DisplayName
		var formattedTeamID string = FormatStringID(teamID)
		return formattedTeamID
	} else {
		var teamID string = espnGameDetails.Header.Competitions[0].Competitors[1].Team.DisplayName
		var formattedTeamID string = FormatStringID(teamID)
		return formattedTeamID
	}

}

func parseGameZuluTimestamp(espnGameDetails types.ESPNGameDetailsResponse) string {
	var gameDate string = espnGameDetails.Header.Competitions[0].Date
	return gameDate
}

func parseBroadcast(espnGameDetails types.ESPNGameDetailsResponse) string {
	var broadcast string = espnGameDetails.Header.Competitions[0].Broadcasts[0].Media.ShortName
	return broadcast
}

func parseGameStatus(espnGameDetails types.ESPNGameDetailsResponse) bool {
	var gameStatus bool = espnGameDetails.Header.Competitions[0].Status.Type.Completed
	return gameStatus

}

func Game(league string, espnGameDetails types.ESPNGameDetailsResponse) Details {
	var details Details

	details.gameID = parseGameID(espnGameDetails)
	details.league = league
	details.week = parseWeek(espnGameDetails)
	details.year = parseYear(espnGameDetails)
	details.espnCode = parseEspnGameCode(espnGameDetails)
	details.cbsCode = ""
	details.foxCode = ""
	details.vegasCode = ""
	details.awayTeamID = parseTeamID("away", espnGameDetails)
	details.homeTeamID = parseTeamID("home", espnGameDetails)
	details.zuluTimestamp = parseGameZuluTimestamp(espnGameDetails)
	details.broadcast = parseBroadcast(espnGameDetails)
	details.finished = parseGameStatus(espnGameDetails)

	return details
}
