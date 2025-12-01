package transform

import (
	"have-a-nice-pickem-etl/etl/pickemstructs"
	"have-a-nice-pickem-etl/etl/transform/bettingodds"
	"have-a-nice-pickem-etl/etl/transform/common"

	"github.com/PuerkitoBio/goquery"
)

type FoxOdds struct {
	pickemstructs.BettingOdds
}

// Instantiates ESPN Betting Odds record from various sources
func CreateESPNBettingOddsRecord(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) pickemstructs.BettingOdds {
	var espnRecord bettingodds.EspnOdds = bettingodds.EspnOdds{}

	var newRecord pickemstructs.BettingOdds = pickemstructs.BettingOdds{
		GameID:        common.ParseGameID(consolidatedGameProperties),
		Source:        "ESPN",
		OverUnder:     espnRecord.ParseOverUnder(consolidatedGameProperties),
		AwayMoneyline: espnRecord.ParseAwayMoneyline(consolidatedGameProperties),
		HomeMoneyline: espnRecord.ParseHomeMoneyline(consolidatedGameProperties),
		AwaySpread:    espnRecord.ParseAwaySpread(consolidatedGameProperties),
		HomeSpread:    espnRecord.ParseHomeSpread(consolidatedGameProperties),
	}

	return newRecord
}

// Instantiates CBS Betting Odds record from various sources
func CreateCBSBettingOddsRecord(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) pickemstructs.BettingOdds {
	var cbsRecord bettingodds.CbsOdds = bettingodds.CbsOdds{}
	var gameID string = common.ParseGameID(consolidatedGameProperties)
	var cbsGameCode = common.ExtractCbsGameCode(consolidatedGameProperties.CbsPage, gameID)
	var cbsGameTable *goquery.Selection = cbsRecord.ParseGameOddsTable(consolidatedGameProperties.CbsPage, cbsGameCode)

	var newRecord pickemstructs.BettingOdds = pickemstructs.BettingOdds{
		GameID:        gameID,
		Source:        "CBS",
		OverUnder:     cbsRecord.ParseOverUnder(cbsGameTable),
		AwayMoneyline: cbsRecord.ParseAwayMoneyline(cbsGameTable),
		HomeMoneyline: cbsRecord.ParseHomeMoneyline(cbsGameTable),
		AwaySpread:    cbsRecord.ParseAwaySpread(cbsGameTable),
		HomeSpread:    cbsRecord.ParseHomeSpread(cbsGameTable),
	}

	return newRecord
}
