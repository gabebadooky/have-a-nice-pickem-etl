package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/bettingodds"
	"have-a-nice-pickem-etl/etl/transform/common"
	"strings"
)

type BettingOdds struct {
	GameID            string
	Source            string
	OverUnder         float32
	AwayMoneyline     uint16
	HomeMoneyline     uint16
	AwaySpread        float32
	HomeSpread        float32
	AwayWinPercentage string
	HomeWinPercentage string
}

func setOverUnder(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) float32 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseOverUnder(espnGameDetails)
	default:
		return 0.00
	}
}

func setAwayMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) uint16 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseAwayMoneyline(espnGameDetails)
	default:
		return 0
	}
}

func setHomeMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) uint16 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseHomeMoneyline(espnGameDetails)
	default:
		return 0
	}
}

func setAwaySpread(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) float32 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseOverUnder(espnGameDetails)
	default:
		return 0.00
	}
}

func setHomeSpread(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) float32 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseOverUnder(espnGameDetails)
	default:
		return 0.00
	}
}

func setAwayWinPercentage(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) string {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseAwayWinPercentage(espnGameDetails)
	default:
		return ""
	}
}

func setHomeWinPercentage(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) string {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseHomeWinPercentage(espnGameDetails)
	default:
		return ""
	}
}

// Instantiates Betting Odds record from various sources
func CreateBettingOddsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) BettingOdds {
	var newRecord BettingOdds

	newRecord.GameID = common.ParseGameID(espnGameDetails)
	newRecord.Source = source
	newRecord.OverUnder = setOverUnder(espnGameDetails, source)
	newRecord.AwayMoneyline = setAwayMoneyline(espnGameDetails, source)
	newRecord.HomeMoneyline = setHomeMoneyline(espnGameDetails, source)
	newRecord.AwaySpread = setAwaySpread(espnGameDetails, source)
	newRecord.HomeSpread = setHomeSpread(espnGameDetails, source)
	newRecord.AwayWinPercentage = setAwayWinPercentage(espnGameDetails, source)
	newRecord.HomeWinPercentage = setHomeWinPercentage(espnGameDetails, source)

	return newRecord
}
