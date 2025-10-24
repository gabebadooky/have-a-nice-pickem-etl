package transform

import (
	"fmt"
	"time"
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
	finished      string
}

func parseGameID(espnGameDetails map[string]any) string {
	var awayTeamID string = parseTeamID("away", espnGameDetails)
	var homeTeamID string = parseTeamID("home", espnGameDetails)
	var gameID string = fmt.Sprintf("%s-at-%s", awayTeamID, homeTeamID)
	return gameID
}

func parseWeek(espnGameDetails map[string]any) int8 {
	week, err := espnGameDetails["header"].(map[string]int8)["week"]
	if err {
		return -1
	}
	return week

}

func parseYear(espnGameDetails map[string]any) uint16 {
	year, err := espnGameDetails["header"].(map[string]any)["season"].(map[string]uint16)["year"]
	if err {
		return uint16(time.Now().Year())
	}
	return year

}

func parseEspnGameCode(espnGameDetails map[string]any) string {
	gameCode, err := espnGameDetails["header"].(map[string]string)["id"]
	if err {
		return "espn-game-code"
	}
	return gameCode

}

func parseTeamID(homeAway string, espnGameDetails map[string]any) string {
	competitorHomeAway, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[0].(map[string]string)["homeAway"]
	if err {
		return "team-id"
	}

	if competitorHomeAway == homeAway {
		teamID, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[0].(map[string]string)["slug"]
		if err {
			return "team-id"
		}
		return teamID

	} else {
		teamID, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["competitors"].([]any)[1].(map[string]string)["slug"]
		if err {
			return "team-id"
		}
		return teamID

	}

}

func parseGameZuluTimestamp(espnGameDetails map[string]any) string {
	gameDate, err := espnGameDetails["header"].(map[string]any)["competitions"].(map[string]string)["date"]
	if err {
		return ""
	}
	return gameDate
}

func parseBroadcast(espnGameDetails map[string]any) string {
	broadcast, err := espnGameDetails["header"].(map[string]any)["competitions"].(map[string]any)["broadcasts"].([]any)[0].(map[string]any)["media"].(map[string]string)["shortName"]
	if err {
		return ""
	}
	return broadcast
}

func parseGameStatus(espnGameDetails map[string]any) string {
	gameStatus, err := espnGameDetails["header"].(map[string]any)["competitions"].([]any)[0].(map[string]any)["status"].(map[string]any)["type"].(map[string]string)["completed"]
	if err {
		return "Not Started"
	}
	return gameStatus

}

func InstantiateGameDetails(league string, espnGameDetails map[string]any) Details {
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
