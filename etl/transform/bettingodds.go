package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/bettingodds"
	"have-a-nice-pickem-etl/etl/transform/common"
	"strings"
)

func setOverUnder(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) float32 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseOverUnder(espnGameDetails)
	default:
		return 0.00
	}
}

func setAwayMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) int16 {
	switch strings.ToUpper(source) {
	case "ESPN":
		return bettingodds.ParseAwayMoneyline(espnGameDetails)
	default:
		return 0
	}
}

func setHomeMoneyline(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) int16 {
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
func CreateBettingOddsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) pickemstructs.BettingOdds {
	var newRecord pickemstructs.BettingOdds = pickemstructs.BettingOdds{
		GameID:            common.ParseGameID(espnGameDetails),
		Source:            source,
		OverUnder:         setOverUnder(espnGameDetails, source),
		AwayMoneyline:     setAwayMoneyline(espnGameDetails, source),
		HomeMoneyline:     setHomeMoneyline(espnGameDetails, source),
		AwaySpread:        setAwaySpread(espnGameDetails, source),
		HomeSpread:        setHomeSpread(espnGameDetails, source),
		AwayWinPercentage: setAwayWinPercentage(espnGameDetails, source),
		HomeWinPercentage: setHomeWinPercentage(espnGameDetails, source),
	}

	return newRecord
}
