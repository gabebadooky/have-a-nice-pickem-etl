package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/bettingodds"
	"have-a-nice-pickem-etl/etl/transform/common"
)

func CreateBettingOddsRecord(espnGameDetails pickemstructs.ESPNGameDetailsResponse, source string) pickemstructs.BettingOdds {
	var newRecord pickemstructs.BettingOdds

	newRecord.GameID = common.ParseGameID(espnGameDetails)
	newRecord.Source = source
	newRecord.OverUnder = bettingodds.ParseOverUnder(espnGameDetails)
	newRecord.AwayMoneyline = bettingodds.ParseAwayMoneyline(espnGameDetails)
	newRecord.HomeMoneyline = bettingodds.ParseHomeMoneyline(espnGameDetails)
	newRecord.AwaySpread = bettingodds.ParseAwaySpread(espnGameDetails)
	newRecord.HomeSpread = bettingodds.ParseHomeSpread(espnGameDetails)
	newRecord.AwayWinPercentage = bettingodds.ParseAwayWinPercentage(espnGameDetails)
	newRecord.HomeWinPercentage = bettingodds.ParseHomeWinPercentage(espnGameDetails)

	return newRecord
}
